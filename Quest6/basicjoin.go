package piscine

func BasicJoin(strs []string) string {
	var str string
	for i := range strs {
		str = str + strs[i]
	}
	return str
}
