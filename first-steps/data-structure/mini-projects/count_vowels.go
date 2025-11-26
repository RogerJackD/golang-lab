package main

import (
	"fmt"
	"strings"
)

func countVowels(text string) int {
	vowels := "aeiouAEIOU"
	count := 0

	for _, char := range text {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}

func main() {
	word := "Hello Golang"
	fmt.Println("Text:", word)
	fmt.Println("Vowel count:", countVowels(word))
}
