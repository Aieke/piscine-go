package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	argumentRune := []rune(os.Args[0])
	for i := range argumentRune {
		z01.PrintRune(argumentRune[i])
	}
	z01.PrintRune(10)
}
