package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}

func NRune(s string, n int) rune {
	charRune := []rune(s)
	len := 0
	var nrune rune
	for i := range charRune {
		len = i + 1
	}
	if n >= 0 && n <= len {
		nrune = []rune(s)[n-1]
	}
	return nrune
}
