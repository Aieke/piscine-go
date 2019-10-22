package piscine

// NRune function returns the nth rune of a string.
func NRune(s string, n int) rune {
	runes := []rune(s)
	len := 0

	var nrune rune

	for i := range runes {
		len = i + 1
	}

	if n > 0 && n <= len {
		nrune = []rune(s)[n-1]
	}
	return nrune
}