package main

import "fmt"

func esPar(num int) bool {
	return num%2 == 0
}

func main() {
	fmt.Println("¿4 es par?", esPar(4))
	fmt.Println("¿7 es par?", esPar(7))
}
