package piscine

func RecursivePower(nb int, power int) int {
	result := 1
	if power < 0 || nb == 0 && power > 0 {
		return 0
	} else if nb == 0 && power == 0 {
		return 1
	} else if power > 0 {
		return result * nb * RecursivePower(nb, power-1)
	} else {
		return 1
	}
}
