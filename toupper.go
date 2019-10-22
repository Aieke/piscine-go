package piscine

func ToUpper(s string) string {
	strRunes := []rune(s)
	for i := 0; i < len(s); i++ {
		if strRunes[i] >= 'a' && strRunes[i] <= 'z' {
			strRunes[i] = rune(s[i] - 32)
		}
	}
	return string(strRunes)
}
