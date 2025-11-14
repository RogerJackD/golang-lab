package main

import "fmt"

type Estudiante struct {
	Nombre   string
	Edad     int
	Promedio float64
}

func main() {
	e1 := Estudiante{"Ana", 20, 17.5}
	e2 := Estudiante{"Luis", 22, 15.3}
	e3 := Estudiante{"María", 19, 18.0}

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
}
