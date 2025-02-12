package dfa

import (
	"fmt"

	postfix "github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/Postfix"
)

// Table for storing lastpost, first post and follow post for each node in the tree.
type positionTableRow struct {
	token     string
	nullable  bool
	isFinal   bool
	firstPos  []int
	lastPos   []int
	followPos []int
}

func BuildFromPostfix(expresion []postfix.Symbol) *DFA {

	return nil
}

func setFollowPos(root *Node, positionTable map[int]positionTableRow) {
	// calculate follow post of children
	if !root.IsOperator {
		return
	}

	if root.Value == "·" {
		c1 := positionTable[root.Children[0].Id]
		c2 := positionTable[root.Children[1].Id]
		fmt.Printf("NODE: %d\n", root.Id)
		for _, n := range c1.lastPos {
			node := positionTable[n]
			fmt.Printf("\tC1: %d C2: %v\n", n, c2.firstPos)
			node.followPos = getNumberSet(node.followPos, c2.firstPos)
			positionTable[n] = node
		}
	}

	if root.Value == "*" {
		c := positionTable[root.Id]
		for _, n := range c.lastPos {
			node := positionTable[n]
			node.followPos = getNumberSet(node.followPos, c.firstPos)
			positionTable[n] = node
		}
	}

	for _, child := range root.Children {
		setFollowPos(&child, positionTable)
	}
}

func getNodePosition(root *Node, positionTable map[int]positionTableRow) (bool, []int, []int) {
	// If Node is an operator with 2 operands
	if root.IsOperator && root.Operands == 2 {
		if root.Value == "·" {
			return positionConcatenationOperator(root, positionTable)
		} else if root.Value == "|" {
			return positionOrOperator(root, positionTable)
		}
	}
	// If Node is * operator
	if root.IsOperator && root.Operands == 1 && root.Value == "*" {
		return positionKleenOperator(root, positionTable)
	}

	// Else if node is empty string
	if root.Value == "ε" {
		isNullable := true
		firstPos := make([]int, 0)
		lastPos := make([]int, 0)

		positionTable[root.Id] = positionTableRow{
			token:    root.Value,
			nullable: isNullable,
			firstPos: firstPos,
			lastPos:  lastPos,
		}
		return isNullable, firstPos, lastPos
	}

	isNullable := false
	firstPos := []int{root.Id}
	lastPos := []int{root.Id}

	positionTable[root.Id] = positionTableRow{
		token:    root.Value,
		nullable: isNullable,
		firstPos: firstPos,
		lastPos:  lastPos,
	}
	// Then, this means is a leaf of a Final Symbol
	return false, []int{root.Id}, []int{root.Id}
}

func positionKleenOperator(n *Node, positionTable map[int]positionTableRow) (bool, []int, []int) {
	_, firstPos1, lastPos1 := getNodePosition(&n.Children[0], positionTable)
	isNullable := true
	firstPos := firstPos1
	lastPos := lastPos1
	positionTable[n.Id] = positionTableRow{
		token:    n.Value,
		nullable: isNullable,
		firstPos: firstPos,
		lastPos:  lastPos,
	}
	return isNullable, firstPos, lastPos
}

func positionOrOperator(n *Node, positionTable map[int]positionTableRow) (bool, []int, []int) {
	nullable1, firstPos1, lastPos1 := getNodePosition(&n.Children[0], positionTable)
	nullable2, firstPos2, lastPos2 := getNodePosition(&n.Children[1], positionTable)
	isNullable := nullable1 || nullable2
	firstPos := getNumberSet(firstPos1, firstPos2)
	lastPos := getNumberSet(lastPos1, lastPos2)
	positionTable[n.Id] = positionTableRow{
		token:    n.Value,
		nullable: isNullable,
		firstPos: firstPos,
		lastPos:  lastPos,
	}
	return isNullable, firstPos, lastPos
}

func positionConcatenationOperator(n *Node, positionTable map[int]positionTableRow) (bool, []int, []int) {
	nullable1, firstPos1, lastPos1 := getNodePosition(&n.Children[0], positionTable)
	nullable2, firstPos2, lastPos2 := getNodePosition(&n.Children[1], positionTable)
	var firstPos []int
	var lastPos []int

	isNullable := nullable1 && nullable2

	if nullable1 {
		firstPos = getNumberSet(firstPos1, firstPos2)
	} else {
		firstPos = firstPos1
	}

	if nullable2 {
		lastPos = getNumberSet(lastPos1, lastPos2)
	} else {
		lastPos = lastPos2
	}

	positionTable[n.Id] = positionTableRow{
		token:    n.Value,
		nullable: isNullable,
		firstPos: firstPos,
		lastPos:  lastPos,
	}
	return isNullable, firstPos, lastPos
}

// Return a list with all different final symbols (Not operators) from an expresion.
func findFinalSymbols(expresion []postfix.Symbol) []postfix.Symbol {
	symbolsSet := make(map[postfix.Symbol]bool)

	for _, symbol := range expresion {
		if !symbol.IsOperator {
			if _, exist := symbolsSet[symbol]; !exist {
				symbolsSet[symbol] = true
			}
		}
	}

	symbols := make([]postfix.Symbol, len(symbolsSet))

	for symbol := range symbolsSet {
		symbols = append(symbols, symbol)
	}

	return symbols
}

func getNumberSet(a, b []int) []int {
	unique := make(map[int]struct{})
	result := []int{}

	for _, str := range a {
		if _, exists := unique[str]; !exists {
			unique[str] = struct{}{}
			result = append(result, str)
		}
	}

	for _, str := range b {
		if _, exists := unique[str]; !exists {
			unique[str] = struct{}{}
			result = append(result, str)
		}
	}

	return result
}

func removeDuplicates(slice []int) []int {
	seen := make(map[int]struct{})
	result := []int{}

	for _, num := range slice {
		if _, exists := seen[num]; !exists {
			seen[num] = struct{}{}
			result = append(result, num)
		}
	}

	return result
}
