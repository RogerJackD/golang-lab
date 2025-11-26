package main

import "fmt"

func max(numbers []int) int {
	maxValue := numbers[0]
	for _, value := range numbers {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func main() {
	list := []int{3, 10, 6, 21, 8}
	fmt.Println("Max value:", max(list))
}
