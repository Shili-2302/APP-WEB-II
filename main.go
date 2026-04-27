package main

import ( 
	"fmt"
)


// ****** STRUCTS ******
type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	ID         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

func main() {

	clientes := []Cliente{
		{1, "Juan Pérez", "Ingeniería en Sistemas", 100.0},
		{2, "María Gómez", "Administración de Empresas", 150.0},
		{3, "Carlos Rodríguez", "Derecho", 200.0},
	}

	productos := []Producto{
		{1, "Café Americano", 2.50, 100, "Bebida"},
		{2, "Café con Leche", 3.00, 80, "Bebida"},
		{3, "Té Verde", 2.00, 50, "Bebida"},
		{4, "Sándwich", 5.00, 30, "Comida"},
	}

	var pedidos []Pedido

	// ===== CLIENTES =====
	fmt.Println("\n--- CLIENTES INICIALES ---")
	ListarClientes(clientes)

	nuevoCliente := Cliente{4, "Carlos Ruiz", "Derecho", 120.0}
	clientes = AgregarCliente(clientes, nuevoCliente)

	fmt.Println("\n--- DESPUÉS DE AGREGAR ---")
	ListarClientes(clientes)

	clientes = EliminarCliente(clientes, 1)

	fmt.Println("\n--- DESPUÉS DE ELIMINAR ---")
	ListarClientes(clientes)

	// ===== PRODUCTOS =====
	fmt.Println("\n--- PRODUCTOS INICIALES ---")
	ListarProductos(productos)

	nuevoProducto := Producto{5, "Jugo", 1.50, 40, "Bebida"}
	productos = AgregarProducto(productos, nuevoProducto)

	fmt.Println("\n--- DESPUÉS DE AGREGAR ---")
	ListarProductos(productos)

	productos = EliminarProducto(productos, 2)

	fmt.Println("\n--- DESPUÉS DE ELIMINAR ---")
	ListarProductos(productos)

	// ===== PEDIDOS =====
	fmt.Println("\n--- REGISTRAR PEDIDO ---")

	pedidos, err := RegistrarPedido(clientes, productos, pedidos, 2, 1, 2, "2026-04-17")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Pedido registrado correctamente")
	}

	fmt.Println("\n--- CLIENTES ACTUALIZADOS ---")
	ListarClientes(clientes)

	fmt.Println("\n--- PRODUCTOS ACTUALIZADOS ---")
	ListarProductos(productos)

	fmt.Println("\nTotal pedidos:", len(pedidos))
}

// ===== LISTAR =====

func ListarClientes(clientes []Cliente) {

	if len(clientes) == 0 {
		fmt.Println("(no hay clientes registrados)")
		return
	}

	fmt.Println("\n===== LISTA DE CLIENTES =====")
	fmt.Printf("%-5s %-20s %-25s %-10s\n", "ID", "Nombre", "Carrera", "Saldo")
	fmt.Println("----------------------------------------------------------")

	for _, c := range clientes {
		fmt.Printf("%-5d %-20s %-25s $%-10.2f\n",
			c.ID, c.Nombre, c.Carrera, c.Saldo)
	}
}

func ListarProductos(productos []Producto) {

	if len(productos) == 0 {
		fmt.Println("(no hay productos registrados)")
		return
	}

	fmt.Println("\n===== LISTA DE PRODUCTOS =====")
	fmt.Printf("%-5s %-20s %-15s %-10s %-10s\n", "ID", "Nombre", "Categoria", "Precio", "Stock")
	fmt.Println("----------------------------------------------------------")

	for _, p := range productos {
		fmt.Printf("%-5d %-20s %-15s $%-9.2f %-10d\n",
			p.ID, p.Nombre, p.Categoria, p.Precio, p.Stock)
	}
}

// ===== CLIENTES =====

func BuscarClientePorID(clientes []Cliente, id int) int {
	for i, c := range clientes {
		if c.ID == id {
			return i
		}
	}
	return -1
}

func AgregarCliente(clientes []Cliente, nuevo Cliente) []Cliente {
	return append(clientes, nuevo)
}

func EliminarCliente(clientes []Cliente, id int) []Cliente {
	index := BuscarClientePorID(clientes, id)

	if index == -1 {
		fmt.Println("Cliente no encontrado")
		return clientes
	}

	clientes = append(clientes[:index], clientes[index+1:]...)
	fmt.Println("Cliente eliminado")
	return clientes
}

//  PRODUCTOS 

func BuscarProductoPorID(productos []Producto, id int) int {
	for i, p := range productos {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func AgregarProducto(productos []Producto, nuevo Producto) []Producto {
	return append(productos, nuevo)
}

func EliminarProducto(productos []Producto, id int) []Producto {
	index := BuscarProductoPorID(productos, id)

	if index == -1 {
		fmt.Println("Producto no encontrado")
		return productos
	}

	productos = append(productos[:index], productos[index+1:]...)
	fmt.Println("Producto eliminado")
	return productos
}

//  PUNTEROS 

func DescontarSaldo(cliente *Cliente, monto float64) error {

	if monto <= 0 {
		return fmt.Errorf("monto inválido")
	}

	if cliente.Saldo < monto {
		return fmt.Errorf("saldo insuficiente")
	}

	cliente.Saldo -= monto
	return nil
}

func DescontarStock(producto *Producto, cantidad int) error {

	if cantidad <= 0 {
		return fmt.Errorf("cantidad inválida")
	}

	if producto.Stock < cantidad {
		return fmt.Errorf("stock insuficiente")
	}

	producto.Stock -= cantidad
	return nil
}

// ===== PEDIDOS =====

func RegistrarPedido(
	clientes []Cliente,
	productos []Producto,
	pedidos []Pedido,
	clienteID int,
	productoID int,
	cantidad int,
	fecha string,
) ([]Pedido, error) {

	idxC := BuscarClientePorID(clientes, clienteID)
	if idxC == -1 {
		return pedidos, fmt.Errorf("cliente no encontrado")
	}

	idxP := BuscarProductoPorID(productos, productoID)
	if idxP == -1 {
		return pedidos, fmt.Errorf("producto no encontrado")
	}

	total := productos[idxP].Precio * float64(cantidad)

	err := DescontarStock(&productos[idxP], cantidad)
	if err != nil {
		return pedidos, err
	}

	err = DescontarSaldo(&clientes[idxC], total)
	if err != nil {
		return pedidos, err
	}

	nuevo := Pedido{
		ID:         len(pedidos) + 1,
		ClienteID:  clienteID,
		ProductoID: productoID,
		Cantidad:   cantidad,
		Total:      total,
		Fecha:      fecha,
	}

	pedidos = append(pedidos, nuevo)

	return pedidos, nil
}
