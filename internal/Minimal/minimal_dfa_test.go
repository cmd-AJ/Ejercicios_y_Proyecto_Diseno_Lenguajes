// Aceptar cualquier caracter

package Minimal

import (
	"fmt"
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

func Test_checkgraph(t *testing.T) {

	// dfa.RenderDFA(adf, "adf_no_iniciado.png")
	table := Initialize_Tabla_Prueba1()

	var tablucha = Initilize_table(table)

	//Posibilidades
	var tuplas = Lista_a_marcar_antes_Finals(tablucha)
	fmt.Println(tablucha) //Ya esta
	fmt.Println(tuplas)   //Ya esta

	for _, values := range table.Table_2D {
		fmt.Println(values)
	}

	for _, values := range tuplas {
		fmt.Println(values)
	}

	tuplas = Recorrer_x_tupla(tuplas, table, tablucha)
	//Recorrer solo una vez para verificar si no hubo algo faltante
	tuplas = Recorrer_x_tupla(tuplas, table, tablucha)

	for _, tuple := range tuplas {
		tablucha[tuple.OuterKey][tuple.InnerKey] = true
	}

	for outerKey, innerMap := range tablucha {
		// Iterate over the inner map
		for innerKey := range innerMap {
			// Print the outer key, inner key, and value
			if innerMap[innerKey] == false {
				ReplaceX_index(outerKey, innerKey, table)
			}
		}
	}

	fmt.Println(tablucha)

	No_duplicates(&table)

	dfa_minimizado := Initialize_DFA_minimizado(&table)

	dfa.RenderDFA(&dfa_minimizado, "DFA_Minimizado.png")
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
		X_index: []string{
			"A", "B", "C", "D", "E", "F", "G",
		},
		Y_index: []string{
			"r", "b",
		},
		Finals: map[string]bool{
			"F": true, "G": true, "A": false, "B": false, "C": false, "D": false, "E": false,
		},
		Initial: "A",
	}

}
