package main

import (
	"fmt"
	"unicode"
)

func revert(str string) string {
	var result string
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

func main() {
	str := "Valera, deres "
	s := ""
	s2 := ""
	for _, val := range str {
		if unicode.IsLetter(val) {

			s += string(val)

		} else {
			s2 += string(val)

		}
		revert(s)
		fmt.Print(s)

	}

	fmt.Print(s2)
}
