package dfa

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	postfix "github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/Postfix"
	"github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/utils"
)

func intSliceToString(slice []int) string {
	strs := make([]string, len(slice))
	for i, v := range slice {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, ", ")
}

func printPositionTable(table map[int]positionTableRow) {
	fmt.Printf("%-5s %-10s %-8s %-8s %-15s %-15s %-15s\n",
		"Key", "Token", "Nullable", "IsFinal", "FirstPos", "LastPos", "FollowPos")
	fmt.Println(strings.Repeat("-", 80))

	for key, row := range table {
		fmt.Printf("%-5d %-10s %-8t %-8t %-20s %-15s %-15s\n",
			key, row.token, row.nullable, row.isFinal,
			intSliceToString(row.firstPos), intSliceToString(row.lastPos), intSliceToString(row.followPos))
	}
}

func printStateSetTable(states []*stateSet, transitionTokens []string) {
	// Print header
	fmt.Printf("%-5s | %-10s | %-7s", "ID", "Value", "isFinal")
	for _, token := range transitionTokens {
		fmt.Printf(" | %-10s", token)
	}
	fmt.Println("\n" + strings.Repeat("-", 23+12*len(transitionTokens)))

	// Print rows
	for _, state := range states {
		// Convert value slice to a comma-separated string
		valueStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(state.value)), ","), "[]")

		// Print ID, Value, and isFinal
		fmt.Printf("%-5d | %-10s | %-7t", state.id, valueStr, state.isFinal)

		// Print transitions
		for _, token := range transitionTokens {
			if nextState, exists := state.transitions[token]; exists {
				fmt.Printf(" | %-10d", nextState.id)
			} else {
				fmt.Printf(" | %-10s", "-")
			}
		}
		fmt.Println()
	}
}

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
