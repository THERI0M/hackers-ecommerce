package checkout

import (
	"fmt"
	"hackers-ecommerce/paquetes/carrito"
)

// MetodoPago define el contrato abstracto (Interfaz)
type MetodoPago interface {
	ProcesarPago(monto float64) (string, error)
}

// Estructuras concretas que implementan la interfaz implícitamente
type TarjetaCredito struct {
	NumeroTarjeta string
	Titular       string
}

func (t TarjetaCredito) ProcesarPago(monto float64) (string, error) {
	if len(t.NumeroTarjeta) != 16 {
		return "", fmt.Errorf("tarjeta inválida (debe tener 16 dígitos)")
	}
	return "TX-CRED-99882", nil
}

type TransferenciaBancaria struct {
	NumeroCuenta string
	Banco        string
}

func (tb TransferenciaBancaria) ProcesarPago(monto float64) (string, error) {
	if tb.Banco == "" {
		return "", fmt.Errorf("se requiere especificar el banco de origen")
	}
	return "TX-BANK-11223", nil
}

// ProcesarOrden aprovecha el Polimorfismo al recibir la Interfaz
func ProcesarOrden(c *carrito.Carrito, metodo MetodoPago) {
	total := c.CalcularTotal()
	const IVA = 0.15 // IVA de Ecuador
	totalConIVA := total * (1 + IVA)

	fmt.Printf("\n--- PROCESANDO ORDEN DE COMPRA (HACKERS E-COMMERCE) ---\n")
	fmt.Printf("Subtotal: $%.2f | Total con IVA (15%%): $%.2f\n", total, totalConIVA)

	txID, err := metodo.ProcesarPago(totalConIVA)
	if err != nil {
		fmt.Printf("[ERROR CRÍTICO] Pago rechazado: %v\n", err)
		return
	}
	fmt.Printf("[ÉXITO] Transacción aprobada. ID: %s\n", txID)
}
