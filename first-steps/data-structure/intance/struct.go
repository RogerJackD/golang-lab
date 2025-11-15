package main

import "fmt"

// Definimos un struct (similar a una clase sencilla)
type Persona struct {
	Nombre string
	Edad   int
}

func main() {

	// Instancia 1: usando valores directos
	p1 := Persona{"Ana", 22}

	// Instancia 2: usando nombres de campos
	p2 := Persona{
		Nombre: "Luis",
		Edad:   30,
	}

	// Instancia 3: usando new (devuelve *Persona)
	p3 := new(Persona)
	p3.Nombre = "Carlos"
	p3.Edad = 28

	// Instancia 4: con puntero y &struct literal
	p4 := &Persona{"María", 25}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
}
