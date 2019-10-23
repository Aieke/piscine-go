package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args

	for i := 1; i < argLen(arguments); i++ {
		for _, letter := range arguments[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune(10)
	}
}
func argLen(arguments []string) int {
	argLength := 0
	for i := range arguments {
		argLength = i + 1
	}
	return argLength
}