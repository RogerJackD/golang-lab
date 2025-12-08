package main

import "fmt"

// Función de Búsqueda Binaria
func busquedaBinaria(arr []int, objetivo int) int {
	izq := 0
	der := len(arr) - 1

	for izq <= der {
		medio := (izq + der) / 2

		if arr[medio] == objetivo {
			return medio
		} else if arr[medio] < objetivo {
			izq = medio + 1
		} else {
			der = medio - 1
		}
	}
	return -1
}

func main() {
	numeros := []int{1, 3, 5, 7, 9, 11, 13, 15}
	objetivo := 7

	indice := busquedaBinaria(numeros, objetivo)

	if indice != -1 {
		fmt.Println("Número encontrado en la posición:", indice)
	} else {
		fmt.Println("Número no encontrado")
	}
}
