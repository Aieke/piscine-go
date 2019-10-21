package piscine

func SortIntegerTable(table []int) {
	size := len(table)
	if size < 2 {
		return
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
}
