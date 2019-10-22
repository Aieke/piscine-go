package student

import "github.com/01-edu/z01"

func SortIntegerTable(table []int) []int {
	size := len(table)
	if size < 2 {
		return table
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if table[j] < table[j-1] {
				var temp = table[j]
				table[j] = table[j-1]
				table[j-1] = temp
			}
		}
	}
	return table
}
func toArray(n int) []int {
	var result []int
	for n > 0 {
		if n == 0 {
			result = append(result, n)
		} else {
			a := n % 10
			result = append(result, a)
		}
		n = n / 10
	}
	return result
}

func PrintNbrInOrder(n int) {
	if n < 0 {

	} else if n == 0 {
		z01.PrintRune('0')
	} else {
		mod := '0'
		for _, i := range SortIntegerTable(toArray(n)) {
			for k := 0; k < i; k++ {
				mod++
			}
			z01.PrintRune(mod)
			mod = '0'
		}
	}
}
