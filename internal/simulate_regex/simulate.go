package simulate_regex

import (
	"strings"

	"github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/dfa"
)

/**
 * Simula el recorrido de un DFA (Autómata Finito Determinista) con una cadena de entrada.
 * La función utiliza la operación de Mover para determinar si la cadena es aceptada por el DFA.
 *
 * Parámetros:
 *  - dfa: Un puntero a la estructura DFA que representa el autómata finito determinista.
 *  - cadena: Un string que representa la cadena de entrada que se quiere evaluar.
 *
 * Retorno:
 *  - Un booleano que indica si la cadena es aceptada (true) o no (false) por el DFA.
 */
func SimulateDFA(dfa *dfa.DFA, cadena string) bool {
	// Convertir la cadena a un slice de caracteres
	simbolos := strings.Split(cadena, "")

	// Inicializar el estado actual como el estado inicial del DFA
	currentState := &dfa.StartState

	// Procesar cada símbolo en la cadena
	for _, simbolo := range simbolos {
		currentState = move(currentState, simbolo)
		if currentState == nil {
			// Si no hay transición definida para el símbolo, la cadena no es aceptada
			return false
		}
	}

	// Verificar si el estado final alcanzado es un estado aceptado
	return currentState.IsFinal
}

/**
 * Mover realiza la operación de transición en un DFA. Dado un estado y un símbolo,
 * retorna el estado alcanzable con ese símbolo.
 *
 * Parámetros:
 *  - state: Un puntero a la estructura DFAState que representa el estado actual.
 *  - symbol: Un string que representa el símbolo con el cual se realiza la transición.
 *
 * Retorno:
 *  - Un puntero a DFAState que contiene el estado alcanzable con el símbolo dado, o nil si no existe transición.
 */
func move(state *dfa.State, symbol string) *dfa.State {
	if nextState, exist := state.Transitions[symbol]; exist {
		return &nextState
	}
	return nil
}
