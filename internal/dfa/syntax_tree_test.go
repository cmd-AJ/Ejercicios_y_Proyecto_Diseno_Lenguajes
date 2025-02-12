package dfa

import (
	"fmt"
	"testing"

	postfix "github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/Postfix"
	"github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/utils"
)

// Función auxiliar para comparar valores de nodos
func compareNodes(t *testing.T, got, want Node) {
	utils.ConfigureLogger()
	if got.Value != want.Value {
		t.Errorf("got %v, want %v", got.Value, want.Value)
	}
}

// Test para una expresión regex simple evaluando una concatenación
func TestBuildASTSimpleConcatenation(t *testing.T) {
	utils.ConfigureLogger()
	regex := "abc"
	_, postfix, _ := postfix.RegexToPostfix(regex)

	want := Node{
		Value: "·",
		Children: []Node{
			{
				Value: "·",
				Children: []Node{
					{Value: "a"},
					{Value: "b"},
				},
			},
			{Value: "c"},
		},
	}

	got := BuildAST(postfix)
	compareNodes(t, got, want)
}

func TestRenderAST(t *testing.T) {
	utils.ConfigureLogger()
	regex := "abc"
	_, postfix, _ := postfix.RegexToPostfix(regex)

	got := BuildAST(postfix)

	fmt.Println(got.String())
	// GenerateImageFromRoot(got, "./hello.png")
}
