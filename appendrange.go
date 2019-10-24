package piscine

func AppendRange(min, max int) []int {
	var answer []int
	if min < max {
		for i := min; i < max; i++ {

			answer = append(answer, i)
		}
	} else {
		return answer
	}

	return answer
}
