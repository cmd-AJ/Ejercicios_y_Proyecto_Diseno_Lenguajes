// Aceptar cualquier caracter

package Minimal

import (
	"testing"

	dfa "github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/dfa"
)

func Test_Check_repeat(t *testing.T) {
	// Initialize the test table
	table := Initialize_Tabla_Prueba1()

	// Print the x_index (assuming you meant y_index)
	var tablucha = Initilize_table(table)

	//"A" es el row y "G" es la columna
	// print(tablucha["A"]["G"])

	// print(tablucha["G"]["G"])

	// showkeys(tablucha)
	var tuplas = Lista_a_marcar_antes_finals(tablucha)

	tuplas = Recorrer_x_tupla(tuplas, table, tablucha)
	//Recorrer solo una vez para verificar si no hubo algo faltante
	tuplas = Recorrer_x_tupla(tuplas, table, tablucha)

	for outerKey, innerMap := range tablucha {
		// Iterate over the inner map
		for innerKey := range innerMap {
			// Print the outer key, inner key, and value
			if innerMap[innerKey] == false {
				Replacex_index(outerKey, innerKey, table)
			}
		}
	}

	println("Index in X")
	for f := 0; f < len(table.x_index); f++ {
		print(" ", table.x_index[f])
	}
	println()
	println("Index in Y")
	for f := 0; f < len(table.y_index); f++ {
		print(" ", table.y_index[f])
	}
	println()

	//Ver si esta bien
	for i := 0; i < len(table.Table_2D); i++ {
		for e := 0; e < len(table.Table_2D[e]); e++ {
			print("	" + table.Table_2D[i][e] + " 	")
		}
		println()
	}
}

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
	q0.Transitions["a"] = q1
	q1.Transitions["b"] = q0

	// Create DFA
	dfa := dfa.DFA{
		StartState: q0,
		States:     []dfa.State{q0, q1},
	}

	return dfa

}

func Initialize_Tabla_Prueba1() Table {

	return Table{
		Table_2D: [][]string{
			{"B", "C"},
			{"D", "E"},
			{"D", "F"},
			{"D", "G"},
			{"D", "G"},
			{"D", "C"},
			{"D", "G"},
		},
		x_index: []string{
			"A", "B", "C", "D", "E", "F", "G",
		},
		y_index: []string{
			"r", "b",
		},
		finals: map[string]bool{
			"F": true, "G": true, "A": false, "B": false, "C": false, "D": false, "E": false,
		},
	}

}
