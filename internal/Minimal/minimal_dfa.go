package Minimal

import (
	"github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/dfa"
)

func Minimize_dfa(table Table, dfa dfa.DFA) Table {

	for e := 0; e < len(table.x_index); e++ {

		if e > (len(table.x_index) - 1) {
			for t := 0; t < len(table.y_index); t++ {

			}
		}

	}

	return table

}

func NewState(id string, isFinal bool, estado dfa.State) dfa.State {
	return dfa.State{
		Id:          id,
		IsFinal:     isFinal,
		Transitions: make(map[Symbol]dfa.State), // Initialize the map
	}
}
