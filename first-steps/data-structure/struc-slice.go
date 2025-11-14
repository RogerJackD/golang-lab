package main

import "fmt"

type Producto struct {
	Nombre string
	Precio float64
}

func filtrarCaros(productos []Producto) []Producto {
	var resultado []Producto
	for _, p := range productos {
		if p.Precio > 50 {
			resultado = append(resultado, p)
		}
	}
	return resultado
}

func main() {
	productos := []Producto{
		{"Mouse", 30},
		{"Teclado", 60},
		{"Monitor", 200},
		{"Cable", 10},
	}

	fmt.Println("Productos caros:", filtrarCaros(productos))
}
