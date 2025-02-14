package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"

	min "github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/Minimal"
	postfix "github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/Postfix"
	"github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/balancer"
	"github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/dfa"
	"github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/simulate_regex"
)

func main() {
	// Define colored outputs
	bold := color.New(color.Bold).SprintFunc()
	success := color.New(color.FgGreen, color.Bold).SprintFunc()
	errorText := color.New(color.FgRed, color.Bold).SprintFunc()
	info := color.New(color.FgCyan).SprintFunc()

	fmt.Println(bold("🚀 Bienvenido a la construcción directa de un AFD"))

	scanner := bufio.NewScanner(os.Stdin)

	// Read regex pattern
	fmt.Println(info("📝 Ingresar una expresión regular para crear un AFD:"))
	scanner.Scan()
	regexPattern := scanner.Text()

	// Read test string
	fmt.Println(info("🔍 Ingresar una cadena de prueba:"))
	scanner.Scan()
	testString := scanner.Text()

	condition, _ := balancer.IsBalanced(regexPattern)

	if condition {
		fmt.Println(success("✅ La expresión está balanceada, continuando con la construcción del AFD..."))

		_, array_symbols, _ := postfix.RegexToPostfix(regexPattern)
		testDFA := dfa.BuildFromPostfix(array_symbols)
		dfa.RenderDFA(testDFA, "adf_no_iniciado.png")

		table := min.Initialize_Tabla_a_ADF(testDFA)
		mapeo := min.Crear_Tabla_minimizar(table)
		for i := 0; i < len(table.Table_2D); i++ {
			mapeo = min.Tuplas_a_sacar(mapeo, table)
		}
		afd := min.Revisar_reemplazar(mapeo, *testDFA)

		dfa.RenderDFA(&afd, "ADF_MIN.png")

		fmt.Println(success("🔎 Estado de la tabla después de la minimización:"))

		fmt.Println(success("🎉 AFD minimizado generado con éxito y guardado como 'DFA_Minimizado.png'"))
		fmt.Println(info("📥 Cadena ingresada para validación: "), testString)

		acceptedByInitialDFA := simulate_regex.SimulateDFA(testDFA, testString)

		//AREGLAR PARA MINIMIZAR
		acceptedByMinimizedDFA := simulate_regex.SimulateDFA(&afd, testString)

		if acceptedByInitialDFA {
			fmt.Println(success("🎉 CADENA ACEPTADA por automata ORIGINAL"))
		} else {
			fmt.Println(errorText("🎉 CADENA RECHAZADA por automata ORIGINAL"))
		}
		if acceptedByMinimizedDFA {
			fmt.Println(success("🎉 CADENA ACEPTADA por automata MINIMIZADO"))
		} else {
			fmt.Println(errorText("🎉 CADENA RECHAZADA por automata MINIMIZADO"))
		}

	} else {
		fmt.Println(errorText("❌ La expresión no está balanceada. Corrígela e inténtalo de nuevo."))
	}
}
