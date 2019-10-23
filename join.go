package piscine

func Join(strs []string, sep string) string {
	var str string
	strLen := 0

	for i := range strs {
		strLen = i + 1
	}

	for i := range strs {
		if i != strLen-1 {
			str = str + strs[i] + sep
		} else {
			str = str + strs[i]
		}
	}
	return str
}
