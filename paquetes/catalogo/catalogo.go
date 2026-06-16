package catalogo

import (
	"errors"
	"fmt"
)

// Producto define un accesorio con encapsulación estricta (atributos privados).
type Producto struct {
	id     int
	nombre string
	precio float64
	stock  int
}

// NewProducto es el constructor que valida el estado inicial de los datos.
func NewProducto(id int, nombre string, precio float64, stock int) (Producto, error) {
	if precio <= 0 {
		return Producto{}, errors.New("error: el precio debe ser mayor a cero")
	}
	if stock < 0 {
		return Producto{}, fmt.Errorf("error: %s no puede iniciar con stock negativo", nombre)
	}
	return Producto{id: id, nombre: nombre, precio: precio, stock: stock}, nil
}

// Getters para lectura segura desde fuera del paquete
func (p Producto) Nombre() string { return p.nombre }

func (p Producto) Precio() float64 { return p.precio }

func (p Producto) Stock() int { return p.stock }

// DecrementarStock es un Setter controlado con manejo de errores
func (p *Producto) DecrementarStock(cantidad int) error {
	if cantidad <= 0 {
		return fmt.Errorf("operación inválida: cantidad debe ser positiva")
	}
	if p.stock < cantidad {
		return fmt.Errorf("quiebre de stock para '%s': solicitado %d, disponible %d", p.nombre, cantidad, p.stock)
	}
	p.stock -= cantidad
	return nil
}
