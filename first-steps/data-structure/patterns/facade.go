package main

import "fmt"

// Subsistema 1
type CPU struct{}

func (c *CPU) Encender() {
	fmt.Println("CPU encendida")
}

// Subsistema 2
type Memoria struct{}

func (m *Memoria) Cargar() {
	fmt.Println("Memoria cargada")
}

// Subsistema 3
type Disco struct{}

func (d *Disco) Leer() {
	fmt.Println("Disco leyendo datos")
}

// Facade
type ComputadoraFacade struct {
	cpu     *CPU
	memoria *Memoria
	disco   *Disco
}

func NewComputadoraFacade() *ComputadoraFacade {
	return &ComputadoraFacade{
		cpu:     &CPU{},
		memoria: &Memoria{},
		disco:   &Disco{},
	}
}

func (f *ComputadoraFacade) Iniciar() {
	f.cpu.Encender()
	f.memoria.Cargar()
	f.disco.Leer()
}

// Cliente
func main() {
	computadora := NewComputadoraFacade()
	computadora.Iniciar()
}
