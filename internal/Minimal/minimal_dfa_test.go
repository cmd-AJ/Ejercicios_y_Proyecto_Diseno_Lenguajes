// Aceptar cualquier caracter

package Minimal

import (
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

func Replacex_index(outerKey string, innerkey string, tabla Table) {
	for i := 0; i < len(tabla.Table_2D); i++ {
		for e := 0; e < len(tabla.Table_2D[i]); e++ {
			if tabla.Table_2D[i][e] == outerKey || tabla.Table_2D[i][e] == innerkey {
				tabla.Table_2D[i][e] = (innerkey + outerKey)
			}
		}
	}
	for i := 0; i < len(tabla.x_index); i++ {
		if tabla.x_index[i] == outerKey || tabla.x_index[i] == innerkey {
			tabla.x_index[i] = (innerkey + outerKey)
		}
	}
}

func Lista_a_marcar_antes_finals(mappings map[string]map[string]bool) []Tuple {

	var tuples []Tuple

	for outerKey, innerMap := range mappings {
		for innerKey := range innerMap {
			tuple := Tuple{OuterKey: outerKey, InnerKey: innerKey}
			tuples = append(tuples, tuple)
		}
	}

	return tuples
}

func Recorrer_x_tupla(tuplas []Tuple, tabla Table, mappings map[string]map[string]bool) []Tuple {
	var newTuplesAdded bool
	var initialLength = len(tuplas)

	for s := 0; s < len(tabla.y_index); s++ {

		for d := 0; d < len(tabla.Table_2D)-1; d++ {
			for k := d + 1; k < len(tabla.Table_2D); k++ {
				tuple := Tuple{OuterKey: tabla.Table_2D[d][s], InnerKey: tabla.Table_2D[k][s]}
				if TupleExists(tuplas, tuple) {
					mappings[tabla.x_index[d]][tabla.x_index[k]] = true
				}

			}
		}
	}
	if newTuplesAdded && len(tuplas) > initialLength {
		Recorrer_x_tupla(tuplas, tabla, mappings)

		return tuplas
	}

	return tuplas

}

func TupleExists(tuples []Tuple, target Tuple) bool {
	for _, t := range tuples {
		if t == target {
			return true
		}
	}
	return false
}

// Debe llenar LLenar la tabla de espacios vacios
func Initilize_table(table Table) map[string]map[string]bool {
	var contador = 0
	tabla := make(map[string]map[string]bool)
	for i := range len(table.x_index) - 1 {

		tabla[table.x_index[i]] = make(map[string]bool)
		//Agregar un if que si es mayor al contador

		for e := len(table.x_index) - 1; e >= 1; e-- {

			if 0 < e-contador {
				tabla[table.x_index[i]][table.x_index[e]] = false
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
		contador++
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
