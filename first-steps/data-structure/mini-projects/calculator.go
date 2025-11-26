package main

import "fmt"

func sumar(a, b float64) float64 {
	return a + b
}

func restar(a, b float64) float64 {
	return a - b
}

func multiplicar(a, b float64) float64 {
	return a * b
}

func dividir(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: división entre cero")
		return 0
	}
	return a / b
}

func main() {
	fmt.Println("Suma:", sumar(10, 5))
	fmt.Println("Resta:", restar(10, 5))
	fmt.Println("Multiplicación:", multiplicar(10, 5))
	fmt.Println("División:", dividir(10, 5))
}
