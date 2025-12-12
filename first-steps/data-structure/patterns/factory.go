package main

import "fmt"

// Interfaz común
type Transporte interface {
	Entregar()
}

// Implementación 1
type Camion struct{}

func (c *Camion) Entregar() {
	fmt.Println("Entrega realizada por CAMIÓN")
}

// Implementación 2
type Barco struct{}

func (b *Barco) Entregar() {
	fmt.Println("Entrega realizada por BARCO")
}

// Factory Method
func CrearTransporte(tipo string) Transporte {
	if tipo == "camion" {
		return &Camion{}
	}
	if tipo == "barco" {
		return &Barco{}
	}
	return nil
}

func main() {
	t1 := CrearTransporte("camion")
	t2 := CrearTransporte("barco")

	t1.Entregar()
	t2.Entregar()
}
