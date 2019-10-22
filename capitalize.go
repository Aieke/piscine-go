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
			if i != 0 && (!((strRuness[i-1] >= 'a' && strRuness[i-1] <= 'z') || (strRuness[i-1] >= 'A' && strRuness[i-1] <= 'Z') || (strRuness[i-1] >= '0' && strRuness[i-1] <= '9'))) {
				if strRuness[i] >= 'a' && strRuness[i] <= 'z' {
					strRuness[i] = rune(s[i] - 32)
					continue
				}
			}
		}

	}
	return string(strRuness)
}
