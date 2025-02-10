package dfa

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
