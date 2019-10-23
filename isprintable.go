package piscine

func IsPrintable(str string) bool {
	runes := []rune(str)

	strLen := 0
	for i := range runes {
		strLen = i + 1
	}
	for i := 0; i < strLen; i++ {
		if !(runes[i] >= 'A' && runes[i] <= 'Z' || runes[i] >= 'a' && runes[i] <= 'z') {
			return false
		}
	}
	return true
}
