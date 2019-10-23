package piscine

func IsUpper(str string) bool {
	runes := []rune(str)

	strLen := 0
	for i := range runes {
		strLen = i + 1
	}
	for i := 0; i < strLen; i++ {
		if !(runes[i] >= 'A' && runes[i] <= 'Z') {
			return false
		}
	}
	return true
}
