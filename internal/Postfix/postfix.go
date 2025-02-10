package postfix

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func RegexToPostfix(expresion string) (string, []Symbol, error) {

	return "", nil, nil
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
