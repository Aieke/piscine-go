package piscine

func IsNumeric(str string) bool {
	runes := []rune(str)
	strLen := 0

	for i := range runes {
		strLen = i + 1
	}
	for i := 0; i < strLen; i++ {
		if !(runes[i] >= '0' && runes[i] <= '9') {
			return false
		}
	}
	return true
}
