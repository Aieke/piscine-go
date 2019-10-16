package main

import "github.com/01-edu/z01"

func main() {

	isNegative(1)
	isNegative(0)
	isNegative(-1)

	z01.PrintRune('\n')
}

func isNegative(nb int) {
	if nb >= 0 {
		z01.PrintRune('F')
	} else {
		z01.PrintRune('T')
	}
}
