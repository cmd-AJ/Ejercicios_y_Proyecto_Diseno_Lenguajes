// Aceptar cualquier caracter

package Minimal

import (
	"fmt"
	"testing"
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
	var tuplas = lista_a_marcar_antes_finals(tablucha)

	tuplas = recorrer_x_tupla(tuplas, table, tablucha)
	//Recorrer solo una vez para verificar si no hubo algo faltante
	tuplas = recorrer_x_tupla(tuplas, table, tablucha)

	for outerKey, innerMap := range tablucha {
		// Iterate over the inner map
		for innerKey := range innerMap {
			// Print the outer key, inner key, and value
			fmt.Printf("OuterKey: %s, InnerKey: %s, Value: %v\n", outerKey, innerKey, innerMap[innerKey])
		}
	}

	//recorrer false

}

func showkeys(mappings map[string]map[string]bool) {

	for outerKey, innerMap := range mappings {
		fmt.Printf("Outer Key: %s\n", outerKey)

		// Loop through the inner map
		for innerKey, value := range innerMap {
			fmt.Printf("  Inner Key: %s, Value: %v\n", innerKey, value)
		}
	}

}

func lista_a_marcar_antes_finals(mappings map[string]map[string]bool) []Tuple {

	var tuples []Tuple

	for outerKey, innerMap := range mappings {
		for innerKey := range innerMap {
			tuple := Tuple{OuterKey: outerKey, InnerKey: innerKey}
			tuples = append(tuples, tuple)
		}
	}

	return tuples
}

func recorrer_x_tupla(tuplas []Tuple, tabla Table, mappings map[string]map[string]bool) []Tuple {
	var newTuplesAdded bool
	var initialLength = len(tuplas)

	for s := 0; s < len(tabla.y_index); s++ {

		for d := 0; d < len(tabla.Table_2D)-1; d++ {
			for k := d + 1; k < len(tabla.Table_2D); k++ {
				tuple := Tuple{OuterKey: tabla.Table_2D[d][s], InnerKey: tabla.Table_2D[k][s]}
				if tupleExists(tuplas, tuple) {
					newTuple := Tuple{OuterKey: tabla.x_index[d], InnerKey: tabla.x_index[k]}
					if !tupleExists(tuplas, newTuple) {
						tuplas = append(tuplas, newTuple)
						mappings[newTuple.OuterKey][newTuple.InnerKey] = true
						newTuplesAdded = true

					}
				}

			}
		}
	}
	if newTuplesAdded && len(tuplas) > initialLength {
		// Only update mappings if the tuple was actually added
		recorrer_x_tupla(tuplas, tabla, mappings)

		return tuplas
	}

	return tuplas

}

func tupleExists(tuples []Tuple, target Tuple) bool {
	for _, t := range tuples {
		if t == target { // Direct comparison since all fields are comparable
			return true
		}
	}
	return false
}

// Debe llenar LLenar la tabla de espacios vacios
func Initilize_table(table Table) map[string]map[string]bool {

	tabla := make(map[string]map[string]bool)
	for i := range len(table.x_index) - 1 {

		tabla[table.x_index[i]] = make(map[string]bool)
		for e := range len(table.x_index) {
			println(table.x_index[e])
			if e > 0 {
				//Este es llenar si son estados finales con true
				if !(table.finals[table.x_index[i]] && table.finals[table.x_index[e]]) {
					if table.finals[table.x_index[i]] == true {
						tabla[table.x_index[i]][table.x_index[e]] = true
					}
					if table.finals[table.x_index[e]] == true {
						tabla[table.x_index[i]][table.x_index[e]] = true
					}
				}

			}
		}
	}

	return tabla

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
