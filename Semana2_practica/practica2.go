package main

import "fmt"

type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

func agregarProducto(productos []Producto, producto Producto) []Producto {
	return append(productos, producto)
}

func calcularTotal(productos []Producto) float64 {
	total := 0.0
	for _, producto := range productos {
		total += producto.Precio * float64(producto.Stock)
	}
	return total
}
func buscarProducto(productos []Producto, nombre string) (Producto, bool) {
	for _, producto := range productos {
		if producto.Nombre == nombre {
			return producto, true
		}
	}
	return Producto{}, false
}

func main() {
	productos := []Producto{
		{Nombre: "Producto 1", Precio: 10.00, Stock: 20},
		{Nombre: "Producto 2", Precio: 20.00, Stock: 10},
		{Nombre: "Producto 3", Precio: 30.00, Stock: 30},
	}
	procucto := Producto{Nombre: "Producto 4", Precio: 40.00, Stock: 40}
	productos = agregarProducto(productos, procucto)
	total := calcularTotal(productos)
	fmt.Printf("Total del inventario: %.2f\n", total)
	procucto, ok := buscarProducto(productos, "Producto 4")
	if ok {
		fmt.Println("Producto encontrado: ", procucto)
	} else {
		fmt.Println("Producto no encontrado")
	}
}
