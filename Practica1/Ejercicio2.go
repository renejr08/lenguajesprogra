package main

import (
	"fmt" //import para poder imprimir
	"math"
)

func main() {
	altura := 5 // Altura del diamante (ajusta este valor)

	if altura%2 == 0 {
		altura++ // Asegurar que la altura sea impar
	}

	mitad := altura / 2

	for i := 0; i < altura; i++ {
		for j := 0; j < int(math.Abs(float64(mitad-i))); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < altura-(2*int(math.Abs(float64(mitad-i)))); j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
