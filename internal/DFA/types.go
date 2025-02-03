package dfa

type Symbol = string

type DFA struct {
	startState State
	States     []State
}

type State struct {
	id          string
	IsFinal     bool
	Transitions map[Symbol]State
}
