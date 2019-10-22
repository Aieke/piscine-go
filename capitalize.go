package piscine

func Capitalize(s string) string {

	strRuness := []rune(s)

	lenStr := 0
	for i := range strRuness {
		lenStr = i + 1
	}

	for i := 0; i < lenStr; i++ {

		if strRuness[i] == 32 || strRuness[i] == '+' {
			i++
			if strRuness[i] >= 'a' && strRuness[i] <= 'z' || strRuness[i] >= 'A' && strRuness[i] <= 'Z' || strRuness[i] >= '0' && strRuness[i] <= '9' {
				strRuness[i] = rune(s[i] - 32)
				continue
			}
		}

	}
	return string(strRuness)
}
