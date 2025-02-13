package dfa

import (
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

type stateSet struct {
	id          int
	value       []int
	transitions map[string]*stateSet
	isFinal     bool
}

func BuildFromPostfix(expresion []postfix.Symbol) *DFA {

	tree := BuildAST(expresion)
	centinelNode := Node{
		Id:         len(expresion),
		Value:      "#",
		Operands:   2,
		Children:   []Node{tree},
		IsOperator: false,
		IsFinal:    true}

	rootNode := Node{
		Id:         -(len(expresion) + 1),
		Value:      "·",
		Operands:   2,
		Children:   []Node{tree, centinelNode},
		IsOperator: true}

	positionTable := make(map[int]positionTableRow, 0)
	getNodePosition(&rootNode, positionTable)
	setFollowPos(&rootNode, positionTable)

	return nil
}

func simplifyStates(
	tokens []string,
	initState []int,
	positionTable map[int]positionTableRow) []*stateSet {

	inititialState := &stateSet{id: 0, value: initState, transitions: make(map[string]*stateSet)}
	states := []*stateSet{inititialState}
	queue := []*stateSet{inititialState}

	for len(queue) > 0 {
		currentState := queue[0] // Get a new element from queue
		queue = queue[1:]        // Pop the element

		// Get SET for each character
		for _, token := range tokens {
			newSet := getNewSetForToken(currentState.value, token, positionTable)
			// fmt.Printf("ID: %d SET: %v TOKEN: %s \n", currentState.id, newSet.value, token)

			setAlreadyExist, repeatedSet := setExists(&newSet, states)

			// If set does not exist append it
			if !setAlreadyExist {
				// fmt.Printf("\tENTER: %v\n", newSet.value)
				newSet.id = len(states)
				currentState.transitions[token] = &newSet
				queue = append(queue, &newSet)
				states = append(states, &newSet)
			} else {
				currentState.transitions[token] = repeatedSet
			}
			// for _, r := range states {
			// 	fmt.Printf("%v ", r.value)
			// }
			// fmt.Println("")
		}
		// fmt.Println("==================")
	}
	// fmt.Println("===================")
	// for _, a := range states {
	// 	fmt.Printf("ID: %d %v %v\n", a.id, a.value, a.transitions["b"])
	// }

	return states
}

func getNewSetForToken(items []int, token string, positionTable map[int]positionTableRow) stateSet {
	setItems := make([]int, 0)

	for _, i := range items {
		row := positionTable[i]
		if row.token == token {
			setItems = append(setItems, row.followPos...)
		}
	}

	finalItems := removeDuplicates((setItems))
	isFinal := false
	for _, item := range finalItems {
		if positionTable[item].isFinal {
			isFinal = true
			break
		}
	}

	return stateSet{
		value:       removeDuplicates(setItems), // fcalculate the UNION of followPos
		isFinal:     isFinal,
		transitions: make(map[string]*stateSet),
	}
}

// Function to check if a stateSet exists in a list based on value comparison
func setExists(newSet *stateSet, sets []*stateSet) (bool, *stateSet) {
	for _, existingSet := range sets {
		if slicesAreEqual(newSet.value, existingSet.value) {
			return true, existingSet
		}
	}
	return false, nil
}

// Helper function to check if two slices contain the same elements
func slicesAreEqual(a, b []int) bool {
	// fmt.Printf("\t %v %v \n", a, b)
	if len(a) != len(b) {
		return false
	}

	counts := make(map[int]int)

	// Count occurrences in the first slice
	for _, num := range a {
		counts[num]++
	}

	// Check if second slice has the same elements
	for _, num := range b {
		if counts[num] == 0 {
			return false
		}
		counts[num]--
	}

	return true
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

func setFollowPos(root *Node, positionTable map[int]positionTableRow) {
	// calculate follow post of children
	if !root.IsOperator {
		return
	}

	if root.Value == "·" {
		c1 := positionTable[root.Children[0].Id]
		c2 := positionTable[root.Children[1].Id]
		for _, n := range c1.lastPos {
			node := positionTable[n]
			//fmt.Printf("\tC1: %d C2: %v\n", n, c2.firstPos)
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
func findFinalSymbols(expresion []postfix.Symbol) []string {
	symbolsSet := make(map[postfix.Symbol]bool)

	for _, symbol := range expresion {
		if !symbol.IsOperator {
			if _, exist := symbolsSet[symbol]; !exist {
				symbolsSet[symbol] = true
			}
		}
	}

	symbols := make([]string, 0)

	for symbol := range symbolsSet {
		symbols = append(symbols, symbol.Value)
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
