package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto //Lista para productos

var lProductos listaProductos //

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	//*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio}) Acá simplemente agrega un producto.
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
	prod, codError := l.buscarProducto(nombre) // Busca si el producto ya existe en la lista

	if codError == 0 { // Si no hay error, el producto existe.
		prod.cantidad += cantidad
		//Acá entra la modificación requerida.
		// Incrementar el precio si es diferente
		if prod.precio != precio {
			prod.precio = precio
		}
		// Actualizar el producto en la lista
		(*l)[codError] = prod
	} else {
		// Si el producto no existe, agregarlo a la lista
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
}

// Hacer una función adicional que reciba una cantidad potencialmente infinita de productos previamente creados y los agregue a la lista
/*
func (l *listaProductos) agregarProductosVarios(productos ...producto) {
    for _, p := range productos {
        l.agregarProducto(p.nombre, p.cantidad, p.precio)
    }
}
*/

// producto: producto encontrado; int: tipo de retorno para el valor de error.
func (l *listaProductos) buscarProducto(nombre string) (producto, int) { //el retorno es el índice del producto encontrado y -1 si no existe
	// modifique la función para que esta vez solo retorne el producto como tal y que retorne otro valor adicional "err" conteniendo un 0 si no hay error y números mayores si existe algún
	// tipo de error como por ejemplo que el producto no exista. Una vez implementada la nueva función, cambie los usos de la anterior función en funciones posteriores, realizando los cambios
	// que el uso de la nueva función ameriten
	var prod producto
	var i int
	var codError int = 0 //no hay error

	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			prod = (*l)[i] //esta instrucción asigna el producto encontrado en la lista de productos a la variable prod
			return prod, codError
		}
	}
	codError = 1 //Si hay error, no se encontró producto
	return prod, codError
}

func (l *listaProductos) venderProducto(nombre string) {
	// modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar dicha acción
	// modifique posteriormente para que se ingrese de parámetro la cantidad de productos a vender
	prod, errCode := l.buscarProducto(nombre)

	if errCode == 0 { //Si el producto fue encontrado en la lista.
		prodIndex := errCode
		prod.cantidad = prod.cantidad - 1 //cantidad -1 porque se vende.
		if prod.cantidad <= 0 {           // Eliminar el producto de la lista si no hay existencia. Si es el último producto.
			*l = append((*l)[:prodIndex], (*l)[prodIndex+1:]...) //Información buscada de internet - La parte (*l)[:prod] crea una nueva rebanada que contiene todos los elementos hasta justo antes del elemento a eliminar (prod).
			// La parte (*l)[prod+1:] crea una nueva rebanada que contiene todos los elementos después del elemento a eliminar.
			//Luego, el operador ... en la parte final une estas dos rebanadas en una nueva rebanada que se asigna nuevamente a *l, lo que resulta en la eliminación del elemento prod de la rebanada.
			fmt.Printf("Producto %s agotado y eliminado de la lista\n", nombre)
		}
	} else {
		fmt.Printf("Error: Producto %s no encontrado\n", nombre)
	}
}

// haga una función para a partir del nombre del producto, se pueda modificar el precio del mismo Pero utilizando la función buscarProducto MODIFICADA PREVIAMENTE
func (l *listaProductos) modificarPrecio(nombre string, nuevoPrecio int) {
	prod, errCode := l.buscarProducto(nombre)

	if errCode == 0 {
		prod.precio = nuevoPrecio // Producto encontrado, modificar su precio
		fmt.Printf("Precio de %s modificado a %d\n", nombre, nuevoPrecio)
	} else {
		fmt.Printf("Error: Producto %s no encontrado\n", nombre)
	}
}

func llenarDatos() { //Almancear productos
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)

}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	var productosMinimos listaProductos

	for _, prod := range *l {
		if prod.cantidad <= existenciaMinima {
			productosMinimos = append(productosMinimos, prod)
		}
	}

	return productosMinimos
}

//ACÁ ESTÁ LA PRÁCTICA 2 RESUELTA

// Práctica 2 - Punto 1.b
func (l *listaProductos) aumentarInventarioDeMinimos(listaMinimos listaProductos) {
	for _, prodMinimo := range listaMinimos {
		cantidadAComprar := existenciaMinima - prodMinimo.cantidad // Se calcula la cantidad de productos que deben ser comprados para alcanzar la existencia mínima.
		if cantidadAComprar > 0 {
			l.agregarProducto(prodMinimo.nombre, cantidadAComprar, prodMinimo.precio) //Se usa esta función para agregar la cantidad necesaria para el máximo.
			fmt.Printf("Se han comprado %d unidades de %s para alcanzar la existencia mínima.\n", cantidadAComprar, prodMinimo.nombre)
		}
	}
}

// Práctica 2 - Punto 2.b
// Usar función sort.Slice
// rebanadaProductos, es una rebanada de lo que se va a ordenar. ordenartipoOrdenamiento es si se ordena por nombre, cantidad o precio.
func ordenarProductos(rebanadaProductos listaProductos, tipoOrdenamiento string) {
	switch tipoOrdenamiento {
	case "nombre":
		//Índices i y j como argumentos y devuelve un booleano que indica si el elemento en el índice i debe estar antes que el elemento en el índice j en la ordenación.
		sort.SliceStable(rebanadaProductos, func(i, j int) bool { //La función SliceStable ordena la rebanada utilizando el criterio que proporcionas en el cierre.
			return rebanadaProductos[i].nombre < rebanadaProductos[j].nombre //retorna ordenado.
		})
	case "cantidad":
		sort.SliceStable(rebanadaProductos, func(i, j int) bool {
			return rebanadaProductos[i].cantidad < rebanadaProductos[j].cantidad
		})
	case "precio":
		sort.SliceStable(rebanadaProductos, func(i, j int) bool {
			return rebanadaProductos[i].precio < rebanadaProductos[j].precio
		})
	default:
		fmt.Println("Error, no válido el tipo de Ordenamiento")
		return
	}
}

func main() {
	llenarDatos()
	fmt.Println("Lista de productos antes de vender y del ordenamiento:")
	fmt.Println(lProductos)
	fmt.Println("----------------------------------------------------------------------")

	fmt.Println("Lista de productos ordenados por nombre:")
	ordenarProductos(lProductos, "nombre")
	fmt.Println(lProductos)
	fmt.Println(" ")

	fmt.Println("Lista de productos ordenados por precio:")
	ordenarProductos(lProductos, "precio")
	fmt.Println(lProductos)
	fmt.Println(" ")

	fmt.Println("Lista de productos ordenados por cantidad:")
	ordenarProductos(lProductos, "cantidad")
	fmt.Println(lProductos)

	// Agregar productos para vender
	lProductos.venderProducto("arroz")
	lProductos.venderProducto("frijoles")
	lProductos.venderProducto("leche")

	/*
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println("Lista de productos después de vender:")
		fmt.Println(lProductos)
	*/
}
