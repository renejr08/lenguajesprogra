package main

import (
	"fmt" //import para poder imprimir
	"strings"
)

// Función para contar letras
func contarCaracteres(texto string) int {
	return len(texto) //usando función len
}

func contarPalabras(texto string) int {
	palabras := strings.Fields(texto) //Fields es una función para poder dividir palabras basadas en espacios.
	return len(palabras)              //Las palabras acá ya estan divididas, entonces con Len ya solo cuenta las palabras.
}

func contarLineas(texto string) int {
	lineas := strings.Split(texto, "\n") //Divide el texto basado en saltos de linea que se lee como \n
	return len(lineas)
}

func main() {
	//Texto para poder hacer la prueba
	texto := `Probando código
con varias líneas`

	numCaracteres := contarCaracteres(texto)
	numPalabras := contarPalabras(texto)
	numLineas := contarLineas(texto)

	//prints
	fmt.Printf("Número de caracteres: %d\n", numCaracteres)
	fmt.Printf("Número de palabras: %d\n", numPalabras)
	fmt.Printf("Número de líneas: %d\n", numLineas)
}
