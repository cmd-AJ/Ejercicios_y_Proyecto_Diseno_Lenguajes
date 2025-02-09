package postfix

import "strings"

// This file contains logic specifically to manipulating a list of symbols
// or strings

func convertToSymbols(expresion string) ([]Symbol, error) {
	tokens := strings.Split(expresion, "")
	symbols := make([]Symbol, 0)

	for i := 0; i < len(tokens); {
		t1, _ := getTokenInfo(tokens, i)
		t2, t2Exist := getTokenInfo(tokens, i+1)

		if t1 == ESCAPE_SYMBOL {
			if t2Exist {
				symbols = append(symbols, Symbol{
					value:      t2,
					precedence: 60,
					isOperator: false})

				i += 2
				continue
			}
		}
		if operator, isOperator := OPERATORS[t1]; isOperator {
			symbols = append(symbols, operator)
		} else {
			symbols = append(symbols, Symbol{
				value:      t1,
				precedence: 60,
				isOperator: false,
			})
		}
		i++
	}

	return symbols, nil
}

func addConcatenationSymbols(expresion []Symbol) ([]Symbol, error) {

	formattedTokens := make([]Symbol, 0)

	for i := 0; i < len(expresion); {
		s1, _ := getSymbolInfo(expresion, i)
		s2, s2Exist := getSymbolInfo(expresion, i+1)

		// SPECIAL CASE, if Class sctructure encontared skip([abc])
		if s1.value == "[" && s1.isOperator {
			newIndex := i
			// Search for the closing class bracket "]"
			for ; newIndex < len(expresion); newIndex++ {
				step, _ := getSymbolInfo(expresion, newIndex)

				if step.value == "]" && step.isOperator {
					break
				}

				formattedTokens = append(formattedTokens, step)
			}

			i = newIndex // To start with the next symbol after the class
			continue
		}

		formattedTokens = append(formattedTokens, s1)

		if s2Exist && shouldAddConcatenationSymbol(s1, s2) {
			formattedTokens = append(formattedTokens, OPERATORS[CONCAT_SYMBOL])
		}

		i++
	}

	return formattedTokens, nil
}

func shouldAddConcatenationSymbol(s1, s2 Symbol) bool {

	if s2.value == "" {
		return false
	}

	// If both are open or close parenthesis, false
	if (s1.isOperator && s2.isOperator) &&
		((s1.value == "(" && s2.value == "(") ||
			(s1.value == ")" && s2.value == ")")) {
		return false
	}

	// If the S1 is Operator and
	// 	just and need more that 1 operands,
	// 	is an open parenthesis
	// 	nees less than and operand but the next character is an operator
	if s1.isOperator {
		if s1.operands > 1 ||
			(s1.value == "(" && !s2.isOperator) ||
			(s1.operands < 1 && s2.isOperator) {
			return false
		}
	}
	// 	If S2 is an "(" operator
	if s2.isOperator &&
		((s2.value == "(") ||
			(s2.value == "[")) {
		return true
	}
	if s2.isOperator { // If c1 is not operand then
		return false
	}

	return true
}

func getTokenInfo(symbols []string, index int) (s string, exist bool) {
	if index >= len(symbols) {
		s = ""
		exist = false
		return
	}
	s = symbols[index]
	exist = true
	return
}

func getSymbolInfo(symbols []Symbol, index int) (s Symbol, exist bool) {
	if index >= len(symbols) {
		s = Symbol{}
		exist = false
		return
	}
	s = symbols[index]
	exist = true
	return
}
