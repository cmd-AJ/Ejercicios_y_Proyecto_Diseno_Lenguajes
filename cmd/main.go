package main

import (
	"fmt"

	min "github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/Minimal"
	postfix "github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/Postfix"
	"github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/balancer"
	"github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/dfa"
)

func main() {
	fmt.Println("Bienvenido construccion directa de un AFD")

	var i string

	fmt.Println("Ingresar una cadena para crear un afd")
	fmt.Scan(&i)
	condition, valance := balancer.IsBalanced(i)

	fmt.Println("pasos a seguir: \n", valance)
	if condition {
		_, array_symbols, _ := postfix.RegexToPostfix(i)
		adf := dfa.BuildFromPostfix(array_symbols)
		dfa.RenderDFA(adf, "adf_no_iniciado.png")

		table := min.Initialize_Tabla_a_ADF(adf)

		var tablucha = min.Initilize_table(table)

		//Posibilidades
		var tuplas = min.Lista_a_marcar_antes_Finals(tablucha)
		fmt.Println(tablucha) //Ya esta
		fmt.Println(tuplas)   //Ya esta

		for _, values := range table.Table_2D {
			fmt.Println(values)
		}

		for _, values := range tuplas {
			fmt.Println(values)
		}

		tuplas = min.Recorrer_x_tupla(tuplas, table, tablucha)
		//Recorrer solo una vez para verificar si no hubo algo faltante
		tuplas = min.Recorrer_x_tupla(tuplas, table, tablucha)

		for _, tuple := range tuplas {
			tablucha[tuple.OuterKey][tuple.InnerKey] = true
		}

		for outerKey, innerMap := range tablucha {
			// Iterate over the inner map
			for innerKey := range innerMap {
				// Print the outer key, inner key, and value
				if innerMap[innerKey] == false {
					min.ReplaceX_index(outerKey, innerKey, table)
				}
			}
		}

		fmt.Println(tablucha)

		min.No_duplicates(&table)

		dfa_minimizado := min.Initialize_DFA_minimizado(&table)

		dfa.RenderDFA(&dfa_minimizado, "DFA_Minimizado.png")
	}

}
