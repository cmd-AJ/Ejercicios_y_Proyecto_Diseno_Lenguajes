package postfix

import (
	"fmt"
	"strings"

	"github.com/golang-collections/collections/stack"
)

// Converts a regex string to a slice of symbols in postfix
//
// NOTE: Asummes the expresion is balanced in open-close symbols like "()", "[]"
func RegexToPostfix(tokens string) (string, []Symbol, error) {

	// Convert tokens to Symbols, taking into account escaped symbols.
	symbols, err := convertToSymbols(tokens)
	if err != nil {
		return "", nil, err
	}
	// Interchange Especial operators (?, []) to its equivalents
	primitiveExpresion := convertToPrimitiveOperators(symbols)

	// Add Concatenation Symbols
	expresionPrepared, err := addConcatenationSymbols(primitiveExpresion)
	if err != nil {
		return "", nil, err
	}

	// Reorder expresion in postfix notation
	postfixSymbols := shuntingyard(expresionPrepared, false)
	var sb strings.Builder
	for _, token := range postfixSymbols {
		sb.WriteString(token.value)
	}

	return sb.String(), postfixSymbols, nil
}

func shuntingyard(tokens []Symbol, showLogs bool) []Symbol {
	postfix := make([]Symbol, 0)
	stack := stack.New()

	for i, token := range tokens {
		if token.value == "(" && token.isOperator {
			stack.Push(token)
		} else if token.value == ")" && token.isOperator {
			for {
				tokenValue, _ := stack.Peek().(Symbol)

				if tokenValue.value == "(" && token.isOperator {
					break
				}

				postfix = append(postfix, stack.Pop().(Symbol))
			}
			stack.Pop()
		} else {
			for stack.Len() > 0 {
				peekedChar := stack.Peek().(Symbol)

				if peekedChar.precedence >= token.precedence {
					postfix = append(postfix, stack.Pop().(Symbol))
				} else {
					break
				}
			}
			stack.Push(token)
		}

		if showLogs {
			fmt.Printf("\t=== STEP %d ====\n", i)
			fmt.Printf("\t\tStack: %+v \n", (tokens)[0:i+1])
			fmt.Printf("\t\tResponse: %+v \n", postfix)
		}
	}

	for stack.Len() > 0 {
		postfix = append(postfix, stack.Pop().(Symbol))
	}

	return postfix
}
