package dfa

import (
	"testing"

	postfix "github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/Postfix"
	"github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/utils"
)

func TestPositionTable(t *testing.T) {
	utils.ConfigureLogger()
	regex := "(a|b)*abb#"
	_, postfix, _ := postfix.RegexToPostfix(regex)

	ast := BuildAST(postfix)

	table := make(map[int]positionTableRow, 0)

	getNodePosition(&ast, table)

	printPositionTable(table)
}

func TestPositionTableAndFollowPost(t *testing.T) {
	utils.ConfigureLogger()
	regex := "(a|b)*abb#"
	_, postfix, _ := postfix.RegexToPostfix(regex)

	ast := BuildAST(postfix)

	table := make(map[int]positionTableRow, 0)

	getNodePosition(&ast, table)
	setFollowPos(&ast, table)

	printPositionTable(table)
}

func TestFinalTable(t *testing.T) {
	utils.ConfigureLogger()
	regex := "(a|b)*abb#"
	_, postfix, _ := postfix.RegexToPostfix(regex)

	ast := BuildAST(postfix)

	table := make(map[int]positionTableRow, 0)

	_, firstPos, _ := getNodePosition(&ast, table)
	setFollowPos(&ast, table)

	tokens := []string{"a", "b"}
	//findFinalSymbols(postfix)
	states := simplifyStates(tokens, firstPos, table)

	printStateSetTable(states, tokens)
}

func TestDFA(t *testing.T) {
	utils.ConfigureLogger()
	regex := "(a|b)*abb"
	_, postfix, _ := postfix.RegexToPostfix(regex)

	BuildFromPostfix(postfix)

	// RenderDFA(dfa, "./dfa.jpg")
}
