package testing

func SumOfNumbers(s ...int) int {
	result := 0
	for v := range s {
		result += v
	}
	return result
}

func multiply(s ...int) []int {
	result := []int{}
	for v := range s {
		result = append(result, (v * 2))
	}
	return result
}
