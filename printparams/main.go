package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	n := len(arguments)
	for i := 1; i < n; i++ {
		for _, letter := range arguments[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune(10)
	}
}
