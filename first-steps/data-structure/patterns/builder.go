package main

import "fmt"

// Producto
type Computadora struct {
	CPU string
	RAM int
	SSD int
}

// Builder
type ComputadoraBuilder struct {
	computadora *Computadora
}

// Constructor del builder
func NewComputadoraBuilder() *ComputadoraBuilder {
	return &ComputadoraBuilder{
		computadora: &Computadora{},
	}
}

func (b *ComputadoraBuilder) SetCPU(cpu string) *ComputadoraBuilder {
	b.computadora.CPU = cpu
	return b
}

func (b *ComputadoraBuilder) SetRAM(ram int) *ComputadoraBuilder {
	b.computadora.RAM = ram
	return b
}

func (b *ComputadoraBuilder) SetSSD(ssd int) *ComputadoraBuilder {
	b.computadora.SSD = ssd
	return b
}

func (b *ComputadoraBuilder) Build() *Computadora {
	return b.computadora
}

// Cliente
func main() {
	pc := NewComputadoraBuilder().
		SetCPU("Intel i7").
		SetRAM(16).
		SetSSD(512).
		Build()

	fmt.Println("Computadora creada:", pc)
}
