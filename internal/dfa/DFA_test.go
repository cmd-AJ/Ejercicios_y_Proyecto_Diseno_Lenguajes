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

func printPosishTable(table map[int]posishTableRow) {
	fmt.Printf("%-5s %-10s %-8s %-8s %-15s %-15s %-15s\n",
		"Key", "Token", "Nullable", "IsFinal", "FirstPos", "LastPos", "FollowPos")
	fmt.Println(strings.Repeat("-", 80))

	for key, row := range table {
		fmt.Printf("%-5d %-10s %-8t %-8t %-20s %-15s %-15s\n",
			key, row.token, row.nullable, row.isFinal,
			intSliceToString(row.firstPos), intSliceToString(row.lastPos), intSliceToString(row.followPos))
	}
}

func TestPositionTable(t *testing.T) {
	utils.ConfigureLogger()
	regex := "(a|b)*abb#"
	_, postfix, _ := postfix.RegexToPostfix(regex)

	ast := BuildAST(postfix)

	table := make(map[int]posishTableRow, 0)

	getNodePosition(&ast, table)

	printPosishTable(table)
}
