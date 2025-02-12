// Aceptar cualquier caracter

package dfa

import (
	"testing"
)

func Test_Check_repeat(t *testing.T) {
	eldfa := initialize_simpleDFA()

	RenderDFA(&eldfa, "image.png")

}

func initialize_simpleDFA() DFA {

	q0 := State{
		Id:          "q0",
		IsFinal:     false,
		Transitions: make(map[Symbol]State),
	}

	q1 := State{
		Id:          "q1",
		IsFinal:     true,
		Transitions: make(map[Symbol]State),
	}

	// Define transitions
	q0.Transitions["a"] = q0
	q0.Transitions["a"] = q1
	q1.Transitions["b"] = q1
	q1.Transitions["a"] = q1

	// Create DFA
	dfa := DFA{
		StartState: q0,
		States:     []State{q0, q1},
	}

	return dfa

}
