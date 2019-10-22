package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	len := 0
	var nRune rune
	for i := range runes {
		len = i + 1
	}
	for i := 0; i <= len; i++ {
		if i == len {
			nRune = []rune(s)[i-1]
		}

	}
	return nRune

}
