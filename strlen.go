package piscine

import "unicode/utf8"

func StrLen(str string) int {
	str := "Hello world!"
	count := 0
	for len(str) > 0 {
		_, size := utf8.DecodeLastRuneInString(str)
		count++

		str = str[:len(str)-size]
	}
	fmt.Println(count)
}
