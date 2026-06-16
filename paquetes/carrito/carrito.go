package carrito

import (
	"fmt"
	"hackers-ecommerce/paquetes/catalogo" // Importamos usando tu ruta exacta
)

type ItemCarrito struct {
	Producto catalogo.Producto
	Cantidad int
}

type Carrito struct {
	items []ItemCarrito
}

func NewCarrito() *Carrito {
	return &Carrito{items: []ItemCarrito{}}
}

func (c *Carrito) AgregarProducto(p *catalogo.Producto, cantidad int) error {
	if cantidad <= 0 {
		return fmt.Errorf("la cantidad de %s debe ser mayor a cero", p.Nombre())
	}
	if p.Stock() < cantidad {
		return fmt.Errorf("stock insuficiente de %s", p.Nombre())
	}
	c.items = append(c.items, ItemCarrito{Producto: *p, Cantidad: cantidad})
	return nil
}

func (c *Carrito) Items() []ItemCarrito { return c.items }

func (c *Carrito) CalcularTotal() float64 {
	var total float64
	for _, item := range c.items {
		total += item.Producto.Precio() * float64(item.Cantidad)
	}
	return total
}
