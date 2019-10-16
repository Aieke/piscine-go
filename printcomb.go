package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 48; i <= 57; i++ {
		for j := i + 1; j <= 57; j++ {
			for k := j + 1; k <= 57; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)

				if i != 55 || j != 56 || k != 57 {
					z01.PrintRune(44)
					z01.PrintRune(32)
				}
			}
		}
	}
}
