package piscine

func BasicAtoi2(s string) int {
	var result int = 0
	var total int = 0
	ct := 1
	var peremennaya = 1
	counter := 1

	for index, line := range s {
		if line < 58 && line > 47 {
			peremennaya += index

			ct *= counter
			counter = 10
		} else {
			return 0
		}
	}

	slice := [1000]rune{}

	for i, v := range s {
		if v < 58 && v > 47 {
			slice[i] = rune(v)
			result = ConvertRuneToInt(slice[i])
			total += result * ct
			ct /= 10
		}
	}

	return total
}

func ConvertRuneToInt2(r rune) int {
	var value = 0
	for i := '0'; i <= '9'; i++ {
		if i == r {
			return value
		}
		value++
	}
	return value

}
