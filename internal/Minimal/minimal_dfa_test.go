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

func Test_minimizado(t *testing.T) {

	table := Initialize_Tabla_Prueba1()

	k := ConvertTableToDFA(table)
	dfa.RenderDFA(&k, "dfa.png")
	mapeo := Crear_Tabla_minimizar(table)
	for i := 0; i < len(table.Table_2D); i++ {
		mapeo = Tuplas_a_sacar(mapeo, table)
	}
	afd := Revisar_reemplazar(mapeo, k)

	dfa.RenderDFA(&afd, "ADF_MIN.png")

}

func Initialize_Tabla_Prueba1() Table {
	// State mapping: A → 0, B → 1, C → 2, D → 3, E → 4, F → 5, G → 6
	table2D := map[string]map[string]string{
		"0": {"r": "1", "b": "2"}, // A → B, C
		"1": {"r": "3", "b": "4"}, // B → D, E
		"2": {"r": "3", "b": "5"}, // C → D, F
		"3": {"r": "3", "b": "6"}, // D → D, G
		"4": {"r": "3", "b": "6"}, // E → D, G
		"5": {"r": "3", "b": "2"}, // F → D, C
		"6": {"r": "3", "b": "6"}, // G → D, G
	}

	return Table{
		Table_2D: table2D,
		Y_index:  2,
		X_index:  7,
		Finals: map[string]bool{
			"5": true, "6": true, // F and G are final
			"0": false, "1": false, "2": false, "3": false, "4": false,
		},
		Initial: "0", // A → 0
	}
}
func ConvertTableToDFA(tbl Table) dfa.DFA {
	states := make(map[string]dfa.State)

	// Create State objects
	for stateID := range tbl.Table_2D {
		states[stateID] = dfa.State{
			Id:          stateID,
			IsFinal:     tbl.Finals[stateID],
			Transitions: make(map[Symbol]dfa.State),
		}
	}

	// Populate transitions
	for stateID, transitions := range tbl.Table_2D {
		for symbol, targetID := range transitions {
			states[stateID].Transitions[symbol] = states[targetID]
		}
	}

	// Convert map to slice
	stateList := make([]dfa.State, 0, len(states))
	for _, s := range states {
		stateList = append(stateList, s)
	}

	// Build DFA
	return dfa.DFA{
		StartState: states[tbl.Initial],
		States:     stateList,
	}
}
