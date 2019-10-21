package piscine

func IterativeFactorial(nb int) int {
	fact := 1
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		for i := 1; i <= nb; i++ {
			fact = fact * i
		}
	}
	return fact
}
