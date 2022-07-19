package main

import (
	"fmt"
	"strings"
)

func latinLetters(s string) string {

	str := &strings.Builder{}

	for _, word := range s {

		if (word >= 'a' && word <= 'z') || (word >= 'A' && word <= 'Z') {
			str.WriteRune(word)
		}
	}

	return str.String()
}

// func latinLetters(s string) string {
// 	sb := &strings.Builder{}

// 	for _, r := range s {
// 		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
// 			sb.WriteRune(r)
// 		}
// 	}

// 	return sb.String()
// }
func main() {
	fmt.Println(latinLetters("привет world!"))

	fmt.Println(latinLetters(" abc1"))
	fmt.Println(latinLetters(" привет"))
	fmt.Println(latinLetters("11 ! t e    s t %)"))
	fmt.Println(latinLetters("John Уолтер"))
	fmt.Println(latinLetters("wo[r]d"))
}
