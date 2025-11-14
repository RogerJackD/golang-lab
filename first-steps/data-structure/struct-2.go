package main

import "fmt"

type Rectangulo struct {
	Ancho float64
	Alto  float64
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func (r Rectangulo) Perimetro() float64 {
	return 2*r.Ancho + 2*r.Alto
}

func main() {
	r := Rectangulo{10, 5}
	fmt.Println("Área:", r.Area())
	fmt.Println("Perímetro:", r.Perimetro())
}
