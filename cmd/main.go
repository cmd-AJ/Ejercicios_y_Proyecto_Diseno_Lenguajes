package main

import (
	"fmt"

	io "github.com/cmd-AJ/Ejercicios_y_Proyecto_Diseno_Lenguajes/internal/IO"
)

func main() {
	fmt.Println("Bienvenido construccion directa de un AFD")

	var i string

	fmt.Scan(&i)

	io.ReadFile(i)

}
