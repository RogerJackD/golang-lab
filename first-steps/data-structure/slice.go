package main

import "fmt"

func promedio(numeros []int) float64 {
	suma := 0
	for _, n := range numeros {
		suma += n
	}
	return float64(suma) / float64(len(numeros))
}

func main() {
	valores := []int{10, 20, 30, 40, 50}
	fmt.Println("Promedio:", promedio(valores))
}
