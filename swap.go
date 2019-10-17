package piscine

func Swap(a *int, b *int) {
	var z = *a
	*a = *b
	*b = z
}
