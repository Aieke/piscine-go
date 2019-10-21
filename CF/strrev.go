package piscine

func StrRev(s string) string {
	normalString := []rune(s)
	revString := []rune(s)

	var a int
	for i := range normalString {
		a = i
	}
	for i := range revString {
		revString[i] = normalString[a]
		a--
	}
	return string(revString)
}
