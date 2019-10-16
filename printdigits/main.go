package main

import "github.com/01-edu/z01"

func main() {
	var number = 48
	for number <= 57 {
		z01.PrintRune(rune(number))
		number++
	}
	z01.PrintRune('\n')
}
