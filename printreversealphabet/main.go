package main

import "github.com/01-edu/z01"

func main() {
	var lower = 'z'
	for lower >= 'a' {
		z01.PrintRune(lower)
		lower--
	}
	z01.PrintRune(10)
}
