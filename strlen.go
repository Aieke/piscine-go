package piscine

func StrLen(str string) int {
	strRune := []rune(str)
	nb := 0
	for _ := range strRune {
		nb++
	}
	return nb
}
