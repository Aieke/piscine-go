package piscine

func FindNextPrime(nb int) int {
	for {
		if IsPrime(nb) {
			return nb
		}
		nb++
	}
}
func IsPrime(nb int) bool {
	if nb == 0 || nb == 1 || nb < 0 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
