package postfix

import (
	"slices"
	"strings"
	"testing"
)

func areSlicesEqual(t *testing.T, response []Symbol, expect []string) {
	value := ""
	for _, v := range response {
		value += v.Value
	}

	if len(response) < len(expect) {
		t.Fatalf("Response has less characters than expected. Has %d, %d given. %s", len(response), len(expect), value)
	} else if len(response) > len(expect) {
		t.Fatalf("Response has more characters than expected. Has %d, %d given. %s", len(response), len(expect), value)
	}
	for i, expected := range expect {
		if response[i].Value != expected {
			t.Fatalf("Characters not match, Given string %s", value)
		}
	}
}

func TestEscapedSymbols(t *testing.T) {
	symbol, _ := convertToSymbols("\\+")
	expect := Symbol{Value: "+", IsOperator: false, Precedence: 60}

	if symbol[0] != expect {
		t.Fatalf("Symbols %s was not ESCAPED.", (&symbol[0]).Value)
	}
}
func TestFindSubExpresionIndex(t *testing.T) {
	symbols, _ := convertToSymbols("(a+b)c")
	start, end := getSubExpresionIndex(symbols, 4, "(", ")")
	t.Logf("start: %d end: %d \n", start, end)
}

func TestClassInterchangeClasses(t *testing.T) {
	symbols, _ := convertToSymbols("a-cz1-3z")
	answer := interchangeClasses(symbols)
	var sb strings.Builder
	for _, a := range answer {
		sb.WriteString(a.Value)
	}
	t.Logf(sb.String())
}

func TestClassInterchangeClassesEscaped(t *testing.T) {
	symbols, _ := convertToSymbols("a-c\\]d-e")
	answer := interchangeClasses(symbols)
	var sb strings.Builder
	for _, a := range answer {
		sb.WriteString(a.Value)
	}
	t.Logf(sb.String())
}

func TestClassInterchangeOptional(t *testing.T) {
	symbols, _ := convertToSymbols("(ab)")
	answer := interchangeOptional(symbols)
	slices.Reverse(answer)
	expect := strings.Split("((ab)|ε)", "")
	areSlicesEqual(t, answer, expect)
}

func TestClassPositiveLock(t *testing.T) {
	symbols, _ := convertToSymbols("(ab)")
	answer := interchangePositiveLock(symbols)
	slices.Reverse(answer)
	expect := strings.Split("((ab)(ab)*)", "")
	areSlicesEqual(t, answer, expect)
}
func TestTranslateToPrimitivesNested(t *testing.T) {
	symbols, _ := convertToSymbols("(a(j)+)")
	answer := convertToPrimitiveOperators(symbols)
	expect := strings.Split("(a((j)(j)*))", "")
	areSlicesEqual(t, answer, expect)
}

func TestTranslateToPrimitives(t *testing.T) {
	symbols, _ := convertToSymbols("(ab)?[1-3M-O](a(j)+)")
	answer := convertToPrimitiveOperators(symbols)
	expect := strings.Split("((ab)|ε)(1|2|3|M|N|O)(a((j)(j)*))", "")
	areSlicesEqual(t, answer, expect)
}

func TestAddConcatenation(t *testing.T) {
	symbols, _ := convertToSymbols("(ab)?a|m\\+")
	answer, _ := addConcatenationSymbols(symbols)
	expect := strings.Split("(a·b)?·a|m·+", "")
	areSlicesEqual(t, answer, expect)
}

func TestShuntinYard(t *testing.T) {
	symbols, _ := convertToSymbols("(a·b)*·a|m")
	answer := shuntingyard(symbols)
	expect := strings.Split("ab·*a·m|", "")
	areSlicesEqual(t, answer, expect)
}

func TestFormatLongRegex(t *testing.T) {
	symbols, _ := convertToSymbols("(a|b?c+|d*e|fgh|i|j)")
	response, _ := addConcatenationSymbols(symbols)
	expect := strings.Split("(a|b?·c+|d*·e|f·g·h|i|j)", "")
	areSlicesEqual(t, response, expect)
}

func TestPostfix(t *testing.T) {
	_, response, _ := RegexToPostfix("([a-cA-C])+@([a-c])+.(com|org|net)?")
	expect := strings.Split("AB|C|a|b|c|AB|C|a|b|c|*·@·ab|c|ab|c|*··.·co·m·or·g·|ne·t·|ε|·", "")
	areSlicesEqual(t, response, expect)
}

func TestPostfix2(t *testing.T) {
	_, response, _ := RegexToPostfix("(b+a)+")
	expect := strings.Split("bb*·a·bb*·a·*·", "")
	areSlicesEqual(t, response, expect)
}
