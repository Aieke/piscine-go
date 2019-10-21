package piscine

func Sqrt(nb int) int {
	var root int
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			root = i
		}
	}
	return root

}
