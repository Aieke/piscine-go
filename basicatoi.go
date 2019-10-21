package main

import "fmt"

func main() {
	e := "12345"
	runes := []rune(e)
	i := runes[0]
	fmt.Println(i) // 2
}
