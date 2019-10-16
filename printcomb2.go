package main

import "github.com/01-edu/z01"

func main() {

	for a1:= '0'; a1 <= '9'; a1++ {
		for b1 :='0'; b1 <= '9'; b1++ {
			if a1=='9' && b1=='9' {
				z01.PrintRune(10)
				return
			}

				
	for a2:='0'; a2<='9'; a2++ {
		for b2:='0'; b2<='9'; b2++{
			if a2<a1 || (a2==a1 && b2<=b1) {
				continue
			}
				z01.PrintRune(a1)
				z01.PrintRune(b1)
				z01.PrintRune(32)
				z01.PrintRune(a2)
				z01.PrintRune(b2)

				if a1!=57 || b1!=56 || a2!=57 || b2!=57 {
					z01.PrintRune(44)
					z01.PrintRune(32)
				}
		}
	}
}
}

}
