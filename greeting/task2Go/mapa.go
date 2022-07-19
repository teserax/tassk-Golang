// Реализуйте функцию MostPopularWord(words []string) string,
// которая возвращает самое часто встречаемое слово в слайсе.
// Если таких слов несколько, то возвращается первое из них.
package main

import "fmt"

func MostPopularWord(words []string) string {
	m := map[string]int{}
	max := 0
	s := ""
	for _, val := range words {
		m[val]++
		if m[val] > max {
			max = m[val]
			s = val
		}
	}
	fmt.Println(m)
	return s
}
func main() {
	words := []string{"kk", "kk", "sd", "dd", "sd", "rd", "sd", "kk"}
	fmt.Println(MostPopularWord(words))
}

// func MostPopularWord(words []string) string {
// 	wordsCount := make(map[string]int, 0)
// 	mostPopWord := ""
// 	max := 0

// 	for _, word := range words {
// 		wordsCount[word]++
// 		if wordsCount[word] > max {
// 			max = wordsCount[word]
// 			mostPopWord = word
// 		}
// 	}

// 	return mostPopWord
// }
