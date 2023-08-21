package main

import (
	"fmt"
)

// Estructura "calzado"
type calzado struct {
	Modelo string
	Precio float64
	Talla  int
	Stock  int
}

// Función para vender zapatos
func venderZapatos(inventario *[]calzado, modelo string, talla int, cantidad int) {
	for i, zapato := range *inventario {
		if zapato.Modelo == modelo && zapato.Talla == talla { //Validación si hay talla en esa modelo (marca) de zapato.
			if cantidad <= zapato.Stock {
				(*inventario)[i].Stock -= cantidad //Validación para ver si hay suficiente stock de zapatos solicitado.
				fmt.Printf("Venta exitosa: Se vendieron %d pares de %s talla %d\n", cantidad, modelo, talla)
				return
			} else {
				fmt.Printf("No hay suficiente stock de %s talla %d\n", modelo, talla)
				return
			}
		}
	}
	fmt.Printf("No se encontró %s talla %d en el inventario\n", modelo, talla)
}

func main() {
	// Inicializar inventario
	inventario := []calzado{
		{"Nike", 50000, 42, 10},
		{"Adidas", 60000, 38, 5},
		{"Puma", 45000, 40, 8},
	}

	// Realizar ventas de zapatos
	venderZapatos(&inventario, "Nike", 42, 2)
	venderZapatos(&inventario, "Adidas", 38, 3)
	venderZapatos(&inventario, "Puma", 40, 10)
	venderZapatos(&inventario, "Nike", 42, 9) // Intenta vender más zapatos de los que hay en stock

	// Imprimir inventario después de las ventas
	fmt.Println("Inventario después de las ventas:")
	for _, zapato := range inventario {
		fmt.Printf("Modelo: %s, Talla: %d, Precio: %.2f colones, Stock: %d\n", zapato.Modelo, zapato.Talla, zapato.Precio, zapato.Stock)
	}
}
