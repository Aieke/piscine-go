package piscine

func IsAlpha(str string) bool {
	runes := []rune(str)

	strlen := 0
	for i := range runes {
		strlen = i + 1
	}

	for i := 0; i < strlen; i++ {
		if !((runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= '0' && runes[i] <= '9')) {
			return false
		}

	}
	return true
}
