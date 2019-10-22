package piscine

func ToLower(s string) string {
	strRunes := []rune(s)
	for i := 0; i < len(s); i++ {
		if strRunes[i] >= 'A' && strRunes[i] <= 'Z' {
			strRunes[i] = rune(s[i] + 32)
		}
	}
	return string(strRunes)
}
