package main

import "fmt"

func main() {
	var a int
	s := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 6, 7, 7, 8}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[a] < s[j] {
				a++
				s[a] = s[j]

			}
		}

	}
	fmt.Println(s[:a+1])
}
