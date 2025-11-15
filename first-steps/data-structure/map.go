package main

import "fmt"

func contadorPalabras(palabras []string) map[string]int {
	contador := make(map[string]int)

	for _, palabra := range palabras {
		contador[palabra]++
	}

	return contador
}

func main() {
	lista := []string{"go", "java", "go", "python", "go", "java"}
	fmt.Println(contadorPalabras(lista))
}
	