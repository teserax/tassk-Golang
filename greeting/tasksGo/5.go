// Напишите функцию, которая принимает строку из одного или нескольких слов
// и возвращает ту же строку, но с перевернутыми всеми словами из пяти или
// более букв .
// Передаваемые строки будут состоять только из букв и пробелов.
// Пробелы будут включены только в том случае, если присутствует более одного слова.
// Examples: spinWords( "Hey fellow warriors" ) => returns "Hey wollef sroirraw"

package main

import (
	"fmt"
	"strings"
)

func SpinWords(str string) string {

	str1 := ""
	s := strings.Fields(str)

	for _, words := range s {
		if len(words) < 5 {
			str1 += words

		} else if len(words) >= 5 {
			letter := ""
			for i := len(words) - 1; i > -1; i-- {
				letter += string(words[i])
			}
			str1 += " " + letter
		}

	}

	return str1

}

func main() {
	str := "Hey fellow warriors"
	result := SpinWords(str)
	fmt.Print(result)

}
