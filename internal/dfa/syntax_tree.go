package dfa

import (
	postfix "github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/Postfix"
	"github.com/rs/zerolog/log"
)

/*
BuildAST construye un AST a partir de una lista de símbolos en notación postfix.
Parámetros:
  - postfixSymbols: La expresión en notación postfix como una lista de símbolos (Symbol).

Retorno:
  - Un nodo (Node) que representa la raíz del AST construido a partir de la expresión postfix.

Panic:
 1. Si la expresión postfix es inválida, no está balanceada o en el stack hay menos símbolos de los que necesita un operador.
 2. Resultado del stack final no es un solo nodo (tal que la cantidad de operadores relacionados es incorrecta y faltan o sobran símbolos).
*/
func BuildAST(postfixSymbols []postfix.Symbol) Node {
	var stack []Node

	// Recorrer toda la lista de símbolos en notación postfix
	for i, symbol := range postfixSymbols {

		// fmt.Printf("Value: %s isOperator %v numOperators: %d \n", symbol.Value, symbol.IsOperator, symbol.Operands)
		// Verifica si el símbolo es un operador
		if symbol.IsOperator {

			// Obtener la cantidad de símbolos que necesita el operador
			operandCount := symbol.Operands
			if len(stack) < operandCount {
				log.Panic().Msg("Expresión postfix inválida: falta operando")
			}

			// Añadir los símbolos que necesita el operador a operands
			operands := make([]Node, operandCount)
			for i := range operands {
				operands[i] = stack[len(stack)-1] // Agregar el valor a operands
				stack = stack[:len(stack)-1]      // Eliminar ese operando del stack
			}

			// Invierte el orden de los operandos de operands para mantener el orden correcto
			for i, j := 0, len(operands)-1; i < j; i, j = i+1, j-1 {
				operands[i], operands[j] = operands[j], operands[i]
			}
			// Crear un nodo operador con los operandos
			node := Node{
				Id:         -i,
				Value:      symbol.Value,
				Operands:   symbol.Operands,
				Children:   operands,
				IsOperator: true}
			stack = append(stack, node)

		} else {
			// Si no es un operador, es un carácter (Symbol) y se añade al stack
			if symbol.Value == "ε" {
				node := Node{
					Id:         -1, // stands for leaf that must not be taken into account
					Value:      symbol.Value,
					IsOperator: false}
				stack = append(stack, node)
			} else {
				node := Node{
					Id:         i,
					Value:      symbol.Value,
					IsOperator: false}
				stack = append(stack, node)
			}
		}
	}

	if len(stack) != 1 {
		log.Panic().Msg("Expresión postfix inválida: el resultado final no es un solo nodo")
	}
	return stack[0]
}
