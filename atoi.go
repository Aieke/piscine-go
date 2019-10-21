package piscine

func Atoi(s string) int {
	var result int = 0
	var total int = 0
	ct := 1
	counter := 1

	negative := false

	for index, line := range s {
		if line == 45 {
			if index == 0 {
				negative = true
			} else {
				return 0
			}
		} else if line == 43 {
			if index != 0 {
				return 0
			}
		} else if line < 58 && line > 47 {
			ct *= counter
			counter = 10
		} else {
			return 0
		}
	}

	slice := [10000]rune{}

	for i, v := range s {
		if v < 58 && v > 47 {
			slice[i] = rune(v)
			result = ConvertRuneToInt(slice[i])
			total += result * ct
			ct /= 10
		}
	}
	if negative {
		return total * -1
	} else {
		return total
	}

}

func ConvertRuneToInt3(r rune) int {
	var value = 0
	for i := '0'; i <= '9'; i++ {
		if i == r {
			return value
		}
		value++
	}
	return value

}
