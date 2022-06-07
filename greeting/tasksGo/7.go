/*
Напишите функцию, которая принимает строку из одного или нескольких слов
и возвращает ту же строку, но с перевернутыми всеми словами из пяти или
более букв .
Передаваемые строки будут состоять только из букв и пробелов.
Пробелы будут включены только в том случае, если присутствует более одного слова.
Examples: spinWords( "Hey fellow warriors" ) => returns "Hey wollef sroirraw"
*/

/*
суть решения сводится к поиску начала слова и его конца. замене отрезка между началом и концом на обратное слово.
*/

package main

import (
	"fmt"
	"unicode"
)

const wordLen = 5

func main() {

	str := " HEI wollef sroirraw  "

	foundStart, foundEnd := false, false

	start, end := 0, 0

	var result string
	for i := range str {
		if unicode.IsLetter(rune(str[i])) {

			if start == 0 && !foundStart {
				foundStart = true
				foundEnd = false
				start = i

			} else {
				end = i
			}
			continue
		}

		if !foundEnd {
			foundEnd = true
			end = i
		}
		foundStart = false

		if end-start >= wordLen {
			fmt.Println(str[start:end])
			// выделяем слово
			word := str[start:end]
			// переворачиваем
			wordRevert := revert(word)
			// в результирующий слайс добавляем отрезок от конца препыдущего добавления до начала слова
			result += str[len(result):start]
			// добавляем перевернутое слово
			result += wordRevert
		}
		// обнуляем старт и конец для следующего поиска слова
		start = 0
		end = 0
	}

	// когда заканчивается цикл нужно сделать еще одну проверку расстояния, и если надо перевернуть слово и добавить в результат.
	// len(str) > end  если заканчитватеся словом
	if end-start >= wordLen && len(str) > end {
		fmt.Println(str[start : end+1])
		word := str[start : end+1]
		wordRevert := revert(word)
		result += str[len(result):start]
		result += wordRevert
	} else {
		// иначяе добавляем окончание строки
		result += str[len(result):]
	}

	fmt.Printf("%q\n", str)
	fmt.Printf("%q\n", result)
}

func revert(str string) string {
	var result string
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}
