package Minimal

import (
	"fmt"

	"github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/dfa"
)

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

// Al momento de conseguir un pdf se pasa todos los valores a una tabla donde tenemos:
// Tabla nxn: A donde se dirige el estado
// y_index: Las lista de transiciones
// x_index: Los estados que tienen
//Finals: verificar si es un estado estado final o no

func Initialize_Tabla_a_ADF(ADF dfa.DFA) Table {

	lista := [][]string{}
	yy_index := make(map[string]bool)
	states := make(map[string]bool)
	xx_index_list := []string{}
	yy_index_list := []string{}

	var provisional_list []string
	lista = append(lista, []string{})
	for symbol, state := range ADF.StartState.Transitions {
		provisional_list = append(provisional_list, state.Id)
		yy_index[symbol] = true

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
				yy_index[symbol] = true
				if state.IsFinal {
					states[state.Id] = true
				} else {
					states[state.Id] = false
				}

			}

			lista = append(lista, provisional_list)
		}

	}

	for Index := range yy_index {
		yy_index_list = append(yy_index_list, Index)
	}
	for Index := range states {
		xx_index_list = append(yy_index_list, Index)
	}

	fmt.Println(states)

	return Table{
		Table_2D: lista,
		x_index:  xx_index_list,
		y_index:  yy_index_list,
		finals:   states,
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

func NewState(id string, isFinal bool, estado dfa.State) dfa.State {
	return dfa.State{
		Id:          id,
		IsFinal:     isFinal,
		Transitions: make(map[Symbol]dfa.State), // Initialize the map
	}
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
					if table.finals[table.x_index[i]] {
						tabla[table.x_index[i]][table.x_index[e]] = true
					}
					if table.finals[table.x_index[e]] {
						tabla[table.x_index[i]][table.x_index[e]] = true
					}
				}

			}
		}
		contador++
	}

	return tabla
}

//Hacer otra funcion para revisar si en la minimizacion se cambio el estado inicial
