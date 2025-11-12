package main

import "fmt"

func invertirTexto(texto string) string {
	runas := []rune(texto)
	for i, j := 0, len(runas)-1; i < j; i, j = i+1, j-1 {
		runas[i], runas[j] = runas[j], runas[i]
	}
	return string(runas)
}

func main() {
	fmt.Println("Invertido:", invertirTexto("golang"))
}
