package main

import (
	"fmt"
)

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
func Capitalize(s string) string {

	strRuness := []rune(s)

	for i := 0; i < len(s); i++ {

		if strRuness[i] == 32 || strRuness[i] == '+' {
			i++
			if strRuness[i] >= 'a' && strRuness[i] <= 'z' {
				strRuness[i] = rune(s[i] - 32)
				continue
			}
		}

	}
	return string(strRuness)
}
