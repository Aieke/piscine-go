package piscine

// PrintCombN 123
func PrintCombN(n int) {
	res := 1

	for i := n; i > 1; i-- {
		res *= i
		//	z01.PrintRune(res)
	}

}
