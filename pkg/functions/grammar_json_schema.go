package functions

// a golang port of https://github.com/ggerganov/llama.cpp/pull/1887

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/go-skynet/LocalAI/pkg/utils"
)

const (
	JSONBNF = `root   ::= object
value  ::= object | array | string | number | ("true" | "false" | "null") ws

object ::=
  "{" ws (
            string ":" ws value
    ("," ws string ":" ws value)*
  )? "}" ws

array  ::=
  "[" ws (
            value
    ("," ws value)*
  )? "]" ws

string ::=
  "\"" (
    [^"\\] |
    "\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F]) # escapes
  )* "\"" ws

number ::= ("-"? ([0-9] | [1-9] [0-9]*)) ("." [0-9]+)? ([eE] [-+]? [0-9]+)? ws

ws ::= ([ \t\n] ws)?`
)

var (
	SPACE_RULE = `" "?`

	PRIMITIVE_RULES = map[string]string{
		"boolean": `("true" | "false") space`,
		"number":  `("-"? ([0-9] | [1-9] [0-9]*)) ("." [0-9]+)? ([eE] [-+]? [0-9]+)? space`,
		"integer": `("-"? ([0-9] | [1-9] [0-9]*)) space`,
		"string": `"\"" (
			[^"\\] |
			"\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F])
		  )* "\"" space`,
		"freestring": `(
			[^"\\] |
			"\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F])
		  )* space`,
		"null": `"null" space`,
	}

	INVALID_RULE_CHARS_RE     = regexp.MustCompile(`[^a-zA-Z0-9-]+`)
	GRAMMAR_LITERAL_ESCAPE_RE = regexp.MustCompile(`[\r\n"]`)
	GRAMMAR_LITERAL_ESCAPES   = map[string]string{
		"\r": `\r`,
		"\n": `\n`,
		`"`:  `\"`,
	}
)

type JSONSchemaConverter struct {
	propOrder map[string]int
	rules     map[string]string
}

func NewJSONSchemaConverter(propOrder string) *JSONSchemaConverter {
	propOrderSlice := strings.Split(propOrder, ",")
	propOrderMap := make(map[string]int)
	for idx, name := range propOrderSlice {
		propOrderMap[name] = idx
	}

	rules := make(map[string]string)
	rules["space"] = SPACE_RULE

	return &JSONSchemaConverter{
		propOrder: propOrderMap,
		rules:     rules,
	}
}

func (sc *JSONSchemaConverter) formatLiteral(literal interface{}) string {
	escaped := GRAMMAR_LITERAL_ESCAPE_RE.ReplaceAllStringFunc(jsonString(literal), func(match string) string {
		return GRAMMAR_LITERAL_ESCAPES[match]
	})
	return fmt.Sprintf(`"%s"`, escaped)
}

func (sc *JSONSchemaConverter) addRule(name, rule string) string {
	escName := INVALID_RULE_CHARS_RE.ReplaceAllString(name, "-")
	key := escName
	if existingRule, ok := sc.rules[escName]; ok && existingRule != rule {
		i := 0
		for {
			key = fmt.Sprintf("%s%d", escName, i)
			if _, ok := sc.rules[key]; !ok {
				break
			}
			i++
		}
	}
	sc.rules[key] = rule
	return key
}

const array = `arr  ::=
  "[\n"  (
		realvalue
    (",\n"  realvalue)*
  )? "]"`

func (sc *JSONSchemaConverter) finalizeGrammar(suffix string, maybeArray, maybeString bool) string {
	var lines []string

	swapRoot := maybeArray || maybeString || suffix != ""

	// write down the computed rules.
	// if maybeArray is true, we need to add the array rule and slightly tweak the root rule
	for name, rule := range sc.rules {
		if swapRoot && name == "root" {
			name = "realvalue"
		}
		lines = append(lines, fmt.Sprintf("%s ::= %s", name, rule))
	}

	if !swapRoot {
		return strings.Join(lines, "\n")
	}

	newRoot := "realvalue"
	if maybeArray {
		newRoot = "arr | realvalue"
	}

	if suffix != "" {
		// quote newlines in suffix
		suffix = utils.EscapeNewLines(suffix)

		if maybeArray && maybeString {
			newRoot = "(" + newRoot + ")"
		}

		if maybeString {
			//newRoot = "( (\"" + suffix + "\" " + newRoot + ") | freestring ) "
			newRoot = "( \"" + suffix + "\" " + newRoot + " | freestring ) "
		} else {
			newRoot = "\"" + suffix + "\" " + "" + newRoot + ""
		}
	} else if maybeString {
		if maybeArray {
			//	newRoot = "(" + newRoot + ")"
		}

		newRoot = "freestring | " + newRoot
	}

	lines = append(lines, fmt.Sprintf("%s ::= %s", "root", newRoot))
	lines = append(lines, array)

	return strings.Join(lines, "\n")
}

func (sc *JSONSchemaConverter) visit(schema map[string]interface{}, name string, rootSchema map[string]interface{}) string {
	st, existType := schema["type"]
	var schemaType string
	if existType {
		schemaType = st.(string)
	}
	ruleName := name
	if name == "" {
		ruleName = "root"
	}
	_, oneOfExists := schema["oneOf"]
	_, anyOfExists := schema["anyOf"]
	if oneOfExists || anyOfExists {
		var alternatives []string
		oneOfSchemas, oneOfExists := schema["oneOf"].([]interface{})
		anyOfSchemas, anyOfExists := schema["anyOf"].([]interface{})

		if oneOfExists {
			for i, altSchema := range oneOfSchemas {
				alternative := sc.visit(altSchema.(map[string]interface{}), fmt.Sprintf("%s-%d", ruleName, i), rootSchema)
				alternatives = append(alternatives, alternative)
			}
		} else if anyOfExists {
			for i, altSchema := range anyOfSchemas {
				alternative := sc.visit(altSchema.(map[string]interface{}), fmt.Sprintf("%s-%d", ruleName, i), rootSchema)
				alternatives = append(alternatives, alternative)
			}
		}

		rule := strings.Join(alternatives, " | ")
		return sc.addRule(ruleName, rule)
	} else if ref, exists := schema["$ref"].(string); exists {
		referencedSchema := sc.resolveReference(ref, rootSchema)
		return sc.visit(referencedSchema, name, rootSchema)
	} else if constVal, exists := schema["const"]; exists {
		return sc.addRule(ruleName, sc.formatLiteral(constVal))
	} else if enumVals, exists := schema["enum"].([]interface{}); exists {
		var enumRules []string
		for _, enumVal := range enumVals {
			enumRule := sc.formatLiteral(enumVal)
			enumRules = append(enumRules, enumRule)
		}
		rule := strings.Join(enumRules, " | ")
		return sc.addRule(ruleName, rule)
	} else if properties, exists := schema["properties"].(map[string]interface{}); schemaType == "object" && exists {
		propOrder := sc.propOrder
		var propPairs []struct {
			propName   string
			propSchema map[string]interface{}
		}

		for propName, propSchema := range properties {
			propPairs = append(propPairs, struct {
				propName   string
				propSchema map[string]interface{}
			}{propName: propName, propSchema: propSchema.(map[string]interface{})})
		}

		sort.Slice(propPairs, func(i, j int) bool {
			iOrder := propOrder[propPairs[i].propName]
			jOrder := propOrder[propPairs[j].propName]
			if iOrder != 0 && jOrder != 0 {
				return iOrder < jOrder
			}
			return propPairs[i].propName < propPairs[j].propName
		})

		var rule strings.Builder
		rule.WriteString(`"{" space`)

		for i, propPair := range propPairs {
			propName := propPair.propName
			propSchema := propPair.propSchema
			propRuleName := sc.visit(propSchema, fmt.Sprintf("%s-%s", ruleName, propName), rootSchema)

			if i > 0 {
				rule.WriteString(` "," space`)
			}

			rule.WriteString(fmt.Sprintf(` %s space ":" space %s`, sc.formatLiteral(propName), propRuleName))
		}

		rule.WriteString(` "}" space`)
		return sc.addRule(ruleName, rule.String())
	} else if items, exists := schema["items"].(map[string]interface{}); schemaType == "array" && exists {
		itemRuleName := sc.visit(items, fmt.Sprintf("%s-item", ruleName), rootSchema)
		rule := fmt.Sprintf(`"[" space (%s ("," space %s)*)? "]" space`, itemRuleName, itemRuleName)
		return sc.addRule(ruleName, rule)
	} else {
		primitiveRule, exists := PRIMITIVE_RULES[schemaType]
		if !exists {
			panic(fmt.Sprintf("Unrecognized schema: %v", schema))
		}
		if ruleName == "root" {
			schemaType = "root"
		}
		return sc.addRule(schemaType, primitiveRule)
	}
}
func (sc *JSONSchemaConverter) resolveReference(ref string, rootSchema map[string]interface{}) map[string]interface{} {
	if !strings.HasPrefix(ref, "#/$defs/") {
		panic(fmt.Sprintf("Invalid reference format: %s", ref))
	}

	defKey := strings.TrimPrefix(ref, "#/$defs/")
	definitions, exists := rootSchema["$defs"].(map[string]interface{})
	if !exists {
		fmt.Println(rootSchema)

		panic("No definitions found in the schema")
	}

	def, exists := definitions[defKey].(map[string]interface{})
	if !exists {
		fmt.Println(definitions)

		panic(fmt.Sprintf("Definition not found: %s", defKey))
	}

	return def
}
func (sc *JSONSchemaConverter) Grammar(suffix string, schema map[string]interface{}, maybeArray, maybeString bool) string {
	sc.addRule("freestring", PRIMITIVE_RULES["freestring"])
	sc.visit(schema, "", schema)
	return sc.finalizeGrammar(suffix, maybeArray, maybeString)
}

func (sc *JSONSchemaConverter) GrammarFromBytes(suffix string, b []byte, maybeArray, maybeString bool) string {
	var schema map[string]interface{}
	_ = json.Unmarshal(b, &schema)
	return sc.Grammar(suffix, schema, maybeArray, maybeString)
}

func jsonString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

type FunctionName struct {
	Const string `json:"const"`
}

type FunctionProperties struct {
	Function  FunctionName `json:"function"`
	Arguments Argument     `json:"arguments"`
}

type NameProperties struct {
	Function  FunctionName `json:"name"`
	Arguments Argument     `json:"arguments"`
}

type Argument struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
}

type ItemName struct {
	Type       string         `json:"type"`
	Properties NameProperties `json:"properties"`
}

type ItemFunction struct {
	Type       string             `json:"type"`
	Properties FunctionProperties `json:"properties"`
}

type JSONFunctionStructureName struct {
	OneOf []ItemName             `json:"oneOf,omitempty"`
	AnyOf []ItemName             `json:"anyOf,omitempty"`
	Defs  map[string]interface{} `json:"$defs,omitempty"`
}

func (j JSONFunctionStructureName) Grammar(suffix string, propOrder string, maybeArray, maybeString bool) string {
	dat, _ := json.Marshal(j)
	return NewJSONSchemaConverter(propOrder).GrammarFromBytes(suffix, dat, maybeArray, maybeString)
}

type JSONFunctionStructureFunction struct {
	OneOf []ItemFunction         `json:"oneOf,omitempty"`
	AnyOf []ItemFunction         `json:"anyOf,omitempty"`
	Defs  map[string]interface{} `json:"$defs,omitempty"`
}

func (j JSONFunctionStructureFunction) Grammar(suffix string, propOrder string, maybeArray, maybeString bool) string {
	dat, _ := json.Marshal(j)
	return NewJSONSchemaConverter(propOrder).GrammarFromBytes(suffix, dat, maybeArray, maybeString)
}
