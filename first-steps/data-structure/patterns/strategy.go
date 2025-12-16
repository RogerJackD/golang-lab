package main

import "fmt"

// 1. Interfaz Strategy
type MetodoPago interface {
	Pagar(monto float64)
}

// 2. Estrategias concretas
type PagoTarjeta struct{}

func (p *PagoTarjeta) Pagar(monto float64) {
	fmt.Printf("Pagando %.2f con TARJETA\n", monto)
}

type PagoEfectivo struct{}

func (p *PagoEfectivo) Pagar(monto float64) {
	fmt.Printf("Pagando %.2f en EFECTIVO\n", monto)
}

// 3. Contexto
type Carrito struct {
	metodo MetodoPago
}

func (c *Carrito) SetMetodoPago(m MetodoPago) {
	c.metodo = m
}

func (c *Carrito) Pagar(monto float64) {
	c.metodo.Pagar(monto)
}

// 4. Cliente
func main() {
	carrito := &Carrito{}

	carrito.SetMetodoPago(&PagoTarjeta{})
	carrito.Pagar(150)

	carrito.SetMetodoPago(&PagoEfectivo{})
	carrito.Pagar(80)
}
