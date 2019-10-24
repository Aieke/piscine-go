package piscine

func ConcatParams(args []string) string {
	len := 0
	var result string
	for i := range args {
		len = i + 1
	}
	for i := range args {
		if i != len-1 {
			result = result + args[i] + "\n"
		} else {
			result = result + args[i]
		}
	}
	return result
}
