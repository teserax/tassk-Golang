//Напишите функцию, которая возьмет исходную строку,
//разобьет ее на слова и посчитает количество цифр в каждом слове.
// Причем подсчет для каждого слова запустите в отдельной горутине.
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func findInt(str string) {
	data := strings.Fields(str)
	for _, val := range str {
		if unicode.IsDigit(rune(val)) {
			fmt.Println("is integer", val)
		} else {
			//fmt.Println(string(val))
		}

	}
	fmt.Printf("%v\n", data[0])
}

func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func main() {
	str := "1 a b c 5 ,wwewe   werwrwr wrwrwerwer"
	findInt(str)
}
