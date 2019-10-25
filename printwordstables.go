package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for i := range table {
		runes := []rune(table[i])
		for j := range runes {
			z01.PrintRune(runes[j])
		}
		z01.PrintRune(10)
	}
}
