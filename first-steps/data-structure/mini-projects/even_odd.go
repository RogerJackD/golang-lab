package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	number := 8

	if isEven(number) {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}
}
