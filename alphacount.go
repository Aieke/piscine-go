package piscine

func AlphaCount(str string) int {
	runes := []rune(str)
	counter := 0
	for index := range runes {
		if runes[index] >= 'A' && runes[index] <= 'Z' || runes[index] >= 'a' && runes[index] <= 'z' {
			counter++
		}
	}
	return counter
}
