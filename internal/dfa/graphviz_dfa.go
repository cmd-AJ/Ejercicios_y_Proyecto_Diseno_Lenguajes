package dfa

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GenerateDOT generates a DOT representation of a DFA as a string.
func GenerateDOT_DFA(dfa *DFA) string {
	var sb strings.Builder

	// Write the Graphviz dot header
	sb.WriteString("digraph DFA {\n")
	sb.WriteString("    rankdir=LR;\n") // Left to right orientation

	// Check if the DFA has any states
	if len(dfa.States) == 0 {
		panic("DFA has no states defined.")
	}

	// Define the nodes (states)
	for _, state := range dfa.States {
		shape := "circle"
		if state.IsFinal {
			shape = "doublecircle"
		}
		sb.WriteString(fmt.Sprintf("    \"%s\" [shape=%s];\n", state.Id, shape))

		// Define the transitions
		for _, state := range dfa.States {
			for symbol, toState := range state.Transitions {
				sb.WriteString(fmt.Sprintf("    \"%s\" -> \"%s\" [label=\"%s\"];\n",
					state.Id, toState.Id, symbol))
			}
		}
	}

	// Define the start state
	sb.WriteString(fmt.Sprintf("    \"\" [shape=plaintext,label=\"\"];\n"))
	sb.WriteString(fmt.Sprintf("    \"\" -> \"%s\";\n", dfa.StartState.Id))

	sb.WriteString("}\n")

	return sb.String()
}

// getShape returns the shape for the state node based on whether it's a final state.
func getShape(isFinal bool) string {
	if isFinal {
		return "doublecircle"
	}
	return "circle"
}

// GenerateImage generates an image from the DOT representation using Graphviz
func GenerateImageFromDOT(dot string, outputPath string) error {
	cmd := exec.Command("dot", "-Tpng", "-o", outputPath)
	cmd.Stdin = strings.NewReader(dot)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RenderDFA(dfa *DFA, filename string) error {
	DOT := GenerateDOT_DFA(dfa)
	err := GenerateImageFromDOT(DOT, filename)
	return err
}
