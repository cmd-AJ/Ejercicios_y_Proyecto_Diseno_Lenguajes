package dfa

// =====================
//	  DFA
// =====================

type Symbol = string

type DFA struct {
	StartState State
	States     []State
}

type State struct {
	Id          string
	IsFinal     bool
	Transitions map[Symbol]State
}

// =====================
// ABSTRACT SYNTAX TREE
// =====================

// Definition of a tree node
type Node struct {
	Id int
	// Character itself this node represents
	Value string
	// If this character is an operator or node.
	IsOperator bool
	// If is operator, how many operands needs
	Operands int
	// Insert Children
	Children []Node
	// Reserved for centinel character that marks the end of the parsing.
	// Just one node in the entire tree can have it.
	IsFinal bool
}
