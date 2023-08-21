package main

import (
	"fmt"
)

func rotateArray(arr *[]string, shift int, direction int) {
	length := len(*arr)
	shift %= length // Para manejar desplazamientos mayores a la longitud del arreglo

	if direction == 0 { // Rotación a la izquierda
		*arr = append((*arr)[shift:], (*arr)[:shift]...)
	} else if direction == 1 { // Rotación a la derecha
		*arr = append((*arr)[length-shift:], (*arr)[:length-shift]...)
	}
}

func main() {
	originalSequence := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println("Secuencia Original =", originalSequence)

	rotations := []struct {
		shift     int
		direction int
	}{
		{3, 0}, // Rotación a la izquierda
		{2, 1}, // Rotación a la derecha
		{5, 0}, // Otra rotación a la izquierda
	}

	for _, rotation := range rotations {
		sequenceCopy := make([]string, len(originalSequence))
		copy(sequenceCopy, originalSequence)

		rotateArray(&sequenceCopy, rotation.shift, rotation.direction)

		fmt.Printf("Cantidad de rotaciones = %d\n", rotation.shift)
		fmt.Printf("Dirección = %s\n", map[int]string{0: "izq", 1: "der"}[rotation.direction])
		fmt.Printf("Secuencia final rotada = %v\n\n", sequenceCopy)
	}
}
