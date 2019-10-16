package main

import "github.com/01-edu/z01"

func main() {
    var lower = 'a'
    for lower <= 'z' {
        z01.PrintRune(lower)
        lower++
	}
	z01.PrintRune('\n')

}
