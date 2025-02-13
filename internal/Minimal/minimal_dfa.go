package Minimal

import (
	"github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/dfa"
)

func ReplaceX_index(outerKey string, innerkey string, tabla Table) {
	for i := 0; i < len(tabla.Table_2D); i++ {
		for e := 0; e < len(tabla.Table_2D[i]); e++ {
			if tabla.Table_2D[i][e] == outerKey || tabla.Table_2D[i][e] == innerkey {
				tabla.Table_2D[i][e] = (innerkey + outerKey)
			}
		}
	}
	for i := 0; i < len(tabla.X_index); i++ {
		if tabla.X_index[i] == outerKey || tabla.X_index[i] == innerkey {
			tabla.X_index[i] = (innerkey + outerKey)
		}
	}
}

// Al momento de conseguir un pdf se pasa todos los valores a una tabla donde tenemos:
// Tabla nxn: A donde se dirige el estado
// Y_index: Las lista de transiciones
// X_index: Los estados que tienen
//Finals: verificar si es un estado estado final o no

func Initialize_Tabla_a_ADF(ADF *dfa.DFA) Table {

	lista := [][]string{}
	yY_index := make(map[string]bool)
	states := make(map[string]bool)
	xX_index_list := []string{}
	yY_index_list := []string{}

	var provisional_list []string
	lista = append(lista, []string{})
	for symbol, state := range ADF.StartState.Transitions {
		provisional_list = append(provisional_list, state.Id)
		yY_index[symbol] = true

	}
	lista[0] = append(lista[0], provisional_list...)

	if ADF.StartState.IsFinal {
		states[ADF.StartState.Id] = true
	} else {
		states[ADF.StartState.Id] = false
	}

	for _, state := range ADF.States {
		// OJO En el caso de que nuestro algoritmo
		if state.Id != ADF.StartState.Id {
			provisional_list = []string{}
			for symbol, state := range state.Transitions {
				provisional_list = append(provisional_list, state.Id)
				yY_index[symbol] = true
				if state.IsFinal {
					states[state.Id] = true
				} else {
					states[state.Id] = false
				}

			}
			lista = append(lista, provisional_list)
		}

	}

	for Index := range yY_index {
		yY_index_list = append(yY_index_list, Index)
	}

	for _, Index := range ADF.States {
		xX_index_list = append(xX_index_list, Index.Id)
	}

	return Table{
		Table_2D: lista,
		X_index:  xX_index_list,
		Y_index:  yY_index_list,
		Finals:   states,
		Initial:  ADF.StartState.Id,
	}

}

func Lista_a_marcar_antes_Finals(mappings map[string]map[string]bool) []Tuple {

	var tuples []Tuple

	for outerKey, innerMap := range mappings {
		for innerKey := range innerMap {
			tuple := Tuple{OuterKey: outerKey, InnerKey: innerKey}
			tuples = append(tuples, tuple)
		}
	}

	return tuples
}

func NewState(id string, isFinal bool) dfa.State {
	return dfa.State{
		Id:          id,
		IsFinal:     isFinal,
		Transitions: make(map[Symbol]dfa.State), // Initialize the map
	}
}

func Recorrer_x_tupla(tuplas []Tuple, tabla Table, mappings map[string]map[string]bool) []Tuple {
	var newTuplesAdded bool
	var initialLength = len(tuplas)

	for s := 0; s < len(tabla.Y_index); s++ {

		for d := 0; d < len(tabla.Table_2D)-1; d++ {
			for k := d + 1; k < len(tabla.Table_2D); k++ {
				tuple := Tuple{OuterKey: tabla.Table_2D[d][s], InnerKey: tabla.Table_2D[k][s]}
				if TupleExists(tuplas, tuple) {
					mappings[tabla.X_index[d]][tabla.X_index[k]] = true
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
	for i := range len(table.X_index) - 1 {

		//Genera por cada estado una transicion
		tabla[table.X_index[i]] = make(map[string]bool)
		//Agregar un if que si es mayor al contador

		for e := len(table.X_index) - 1; e >= 1; e-- {

			if 0 < e-contador {
				tabla[table.X_index[i]][table.X_index[e]] = false
				//Este es llenar si son estados finales con true
				if !(table.Finals[table.X_index[i]] && table.Finals[table.X_index[e]]) {
					if table.Finals[table.X_index[i]] {
						tabla[table.X_index[i]][table.X_index[e]] = true
					}
					if table.Finals[table.X_index[e]] {
						tabla[table.X_index[i]][table.X_index[e]] = true
					}
				}

			}
		}
		contador++
	}

	return tabla
}

//Hacer otra funcion para revisar si en la minimizacion se cambio el estado inicial

func Initialize_DFA_minimizado(tabla *Table) dfa.DFA {

	estados := []dfa.State{}
	transitions := map[string]int{}
	var inital int

	// Agrega todos los estados del x index
	for index, value := range tabla.X_index {
		estado := NewState(value, tabla.Finals[value])
		estados = append(estados, estado)
		transitions[value] = index

		if tabla.Initial == value {
			inital = index
		}
	}

	//Para cada uno de los estado
	for index, value := range estados {
		//Hay una transicion de cada y index
		for i := range len(tabla.Y_index) {
			value.Transitions[tabla.Y_index[i]] = estados[transitions[tabla.Table_2D[index][i]]]
		}
	}

	return dfa.DFA{
		StartState: estados[inital],
		States:     estados,
	}

}

func No_duplicates(tabla *Table) {
	seen := make(map[string]bool)
	unique := []string{}
	newMatrix := [][]string{}

	for index, i := range tabla.X_index {
		if !seen[i] {
			seen[i] = true
			unique = append(unique, i)
			newMatrix = append(newMatrix, tabla.Table_2D[index]) // Keep corresponding row
		}
	}

	// Update tabla with unique X_index and corresponding matrix rows
	tabla.X_index = unique
	tabla.Table_2D = newMatrix
}

func Showvalores_tabla(table Table) {
	println("Index in X")
	for f := 0; f < len(table.X_index); f++ {
		print(" ", table.X_index[f])
	}
	println()
	println("Index in Y")
	for f := 0; f < len(table.Y_index); f++ {
		print(" ", table.Y_index[f])
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
