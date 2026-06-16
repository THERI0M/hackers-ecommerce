package main

import (
	"fmt"
	"hackers-ecommerce/paquetes/carrito"
	"hackers-ecommerce/paquetes/catalogo"
	"hackers-ecommerce/paquetes/checkout"
)

func main() {
	fmt.Println("======= EJECUTANDO SISTEMA DE GESTIÓN 'HACKERS' =======")

	// 1. Población del catálogo con validación de errores
	mouse, _ := catalogo.NewProducto(101, "Mouse Razer DeathAdder", 65.50, 10)
	teclado, _ := catalogo.NewProducto(102, "Teclado Corsair K70", 140.00, 2)

	// 2. Inicialización del Carrito
	miCarrito := carrito.NewCarrito()

	// Intentamos agregar productos
	_ = miCarrito.AgregarProducto(&mouse, 2)

	// Prueba de control de error: Forzamos exceso de stock adrede
	fmt.Printf("\n[Prueba Stock] Intentando agregar 3 teclados (Disponibles: %d)...\n", teclado.Stock())
	err := miCarrito.AgregarProducto(&teclado, 3)
	if err != nil {
		fmt.Printf("[ERROR CONTROLADO] -> %v\n", err)
	}

	// Agregamos la cantidad correcta
	_ = miCarrito.AgregarProducto(&teclado, 1)

	// Actualizamos stock físico
	for _, item := range miCarrito.Items() {
		prod := item.Producto
		_ = prod.DecrementarStock(item.Cantidad)
		fmt.Printf("Stock actualizado para %s. Nuevo stock: %d\n", prod.Nombre(), prod.Stock()-item.Cantidad)
	}

	// 3. Demostración de Polimorfismo mediante Interfaces
	pagoTarjeta := checkout.TarjetaCredito{NumeroTarjeta: "1234567812345678", Titular: "Arturo Guevara"}
	checkout.ProcesarOrden(miCarrito, pagoTarjeta)

	pagoTransferencia := checkout.TransferenciaBancaria{NumeroCuenta: "5544332211", Banco: "Banco Pichincha"}
	checkout.ProcesarOrden(miCarrito, pagoTransferencia)
}
