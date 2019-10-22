package piscine

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
