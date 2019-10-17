package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	strRune := []rune(str)
	for i := range strRune {
		z01.PrintRune(strRune[i])
	}
}
