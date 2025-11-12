package main

import "fmt"

func invertirTexto(texto string) string {
	runas := []rune(texto)
	for i, j := 0, len(runas)-1; i < j; i, j = i+1, j-1 {
		runas[i], runas[j] = runas[j], runas[i]
	}
	return string(runas)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("Invertido:", invertirTexto("golang"))
	fmt.Println("Factorial de 5:", factorial(5))
}
