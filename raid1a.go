package student

import (
	"github.com/01-edu/z01"
)

func Raid1a(x, y int) {
	if x > 0 && y > 0 {
		if y > 0 && y < 3 {
			for j := 0; j < y; j++ {
				SecondExample(x, y)
				z01.PrintRune(10)
			}
		} else {
			for j := 1; j <= y; j++ {
				if j == 1 || j == y {
					SecondCheckLastOrFirstY(x, j)
					z01.PrintRune(10)
				} else {
					SecondCheckY(x)
					z01.PrintRune(10)
				}
			}
		}
	}
}

func SecondExample(x, y int) {
	if 0 < x && x < 3 {
		for i := 0; i < x; i++ {
			z01.PrintRune('o')
		}
	} else {
		for i := 1; i <= x; i++ {
			if i == 1 && i == y {
				z01.PrintRune('o')
			} else if i == x && i == 1 {
				z01.PrintRune('o')
			} else if i == 1 {
				z01.PrintRune('o')
			} else if i == x {
				z01.PrintRune('o')
			} else {
				z01.PrintRune(45)
			}
		}
	}
}

func SecondCheckY(x int) {
	for i := 1; i <= x; i++ {
		if i == 1 || i == x {
			z01.PrintRune('|')
		} else {
			z01.PrintRune(32)
		}
	}
}

func SecondCheckLastOrFirstY(x, y int) {
	for i := 1; i <= x; i++ {
		if i == 1 && y != 1 {
			z01.PrintRune('o')
		} else if i == x && y != 1 {
			z01.PrintRune('o')
		} else if i == 1 && y == 1 {
			z01.PrintRune('o')
		} else if i == x && y == 1 {
			z01.PrintRune('o')
		} else {
			z01.PrintRune(45)
		}
	}
}
