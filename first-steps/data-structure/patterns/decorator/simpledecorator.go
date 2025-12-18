package main

import "fmt"

// 1. Componente
type Cafe interface {
	Costo() int
	Descripcion() string
}

type CafeSimple struct{}

func (c *CafeSimple) Costo() int {
	return 5
}

func (c *CafeSimple) Descripcion() string {
	return "Café simple"
}

type CafeDecorador struct {
	cafe Cafe
}

type ConLeche struct {
	CafeDecorador
}

func (c *ConLeche) Costo() int {
	return c.cafe.Costo() + 2
}

func (c *ConLeche) Descripcion() string {
	return c.cafe.Descripcion() + " + leche"
}

type ConAzucar struct {
	CafeDecorador
}

func (c *ConAzucar) Costo() int {
	return c.cafe.Costo() + 1
}

func (c *ConAzucar) Descripcion() string {
	return c.cafe.Descripcion() + " + azúcar"
}

func main() {
	var cafe Cafe = &CafeSimple{}

	cafe = &ConLeche{CafeDecorador{cafe}}
	cafe = &ConAzucar{CafeDecorador{cafe}}

	fmt.Println(cafe.Descripcion())
	fmt.Println("Costo total:", cafe.Costo())
}
