//package basicatoi

package main

import (
	"fmt"
)

func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "000000"

	n := BasicAtoi(s)
	n2 := BasicAtoi(s2)
	n3 := BasicAtoi(s3)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
}


func BasicAtoi(s string) int {
	strRune:=[]rune(s)
	res := 0
	for index, val := range strRune {
		a := 0
		if strRune[i] >= '1' && strRune[i]<='9' || strRune[i]>='-9' {
			a++
		}
		res = res*10+ a
	}
	return res
}