package main

import "fmt"

func esPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	numero := 29
	if esPrimo(numero) {
		fmt.Println(numero, "es un número primo")
	} else {
		fmt.Println(numero, "NO es un número primo")
	}
}
