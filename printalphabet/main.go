package main

import "fmt" 

func main() {
	var lower = 'a'

	for lower <= 'z' {
	
		fmt.Print(string(lower))
		lower++
	}
}