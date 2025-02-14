// Aceptar cualquier caracter

package Minimal

import (
	"testing"

	dfa "github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/dfa"
)

func initialize_simpleDFA() dfa.DFA {

	q0 := dfa.State{
		Id:          "q0",
		IsFinal:     false,
		Transitions: make(map[Symbol]dfa.State),
	}

	q1 := dfa.State{
		Id:          "q1",
		IsFinal:     true,
		Transitions: make(map[Symbol]dfa.State),
	}

	// Define transitions
	q0.Transitions["a"] = q0
	q0.Transitions["b"] = q1
	q1.Transitions["b"] = q1
	q1.Transitions["a"] = q1

	// Create DFA
	dfa := dfa.DFA{
		StartState: q0,
		States:     []dfa.State{q0, q1},
	}

	return dfa

}

// Se inicializa la tabla y se revisa si el mapa tiene un estado y verificar si ese estado es final
func Test_check_DFA(t *testing.T) {

	k := initialize_simpleDFA()

	s := Initialize_Tabla_a_ADF(&k)

	if s.Finals["q1"] != true {
		t.Errorf("Expected %v, but got %v", s.Finals["q1"], "false")
	}

}
