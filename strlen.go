package piscine

func StrLen(str string) int {
	str := "Hello world!"
	count := 0
	for len(str) > 0 {
		count++

		str = str[:len(str)-size]
	}
	fmt.Println(count)
}
