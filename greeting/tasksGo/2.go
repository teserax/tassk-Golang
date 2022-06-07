// Четыре команды (A, B, C и D) сыграли футбольный турнир.
//  Каждая команда сыграла с каждой по одному разу, так что всего прошло 6 матчей.
//  Рассчитайте, сколько очков получила каждая команда.
// Победа приносит 3 очка, ничья — 1, поражение — 0.
// Результаты игр ("ABW", "DCD", "DAW", "CBL", "BDL", "ACW")
// Каждая тройка букв означает одну игру.
//  Внутри тройки первая буква — название первой команды,
//   вторая буква — название второй,
//   третья — код результата (W — первая выиграла, L — первая проиграла, D — ничья).
//Итоговый ответ  "A6 B3 C1 D7"
package main

import (
	"fmt"
)

type match struct {
	first  string
	second string
	result string
}

func result(s []string) map[string]int {
	result := map[string]int{}
	match1 := match{}

	// for _, val := range s {
	// 	result[string(val[0])] = 0
	// 	result[string(val[1])] = 0
	// }
	for _, val := range s {

		match1.first = string(val[0])
		match1.second = string(val[1])
		match1.result = string(val[2])

		switch {
		case match1.result == "W":
			result[match1.first] += 3
		case match1.result == "D":
			result[match1.first] += 1
			result[match1.second] += 1
		case match1.result == "L":
			result[match1.second] += 3
		default:
			result[match1.second] += 0

		}
	}
	return result
}

func main() {
	str := []string{"ABW", "DCD", "DAW", "CBL", "BDL", "ACW"}

	fmt.Println(result(str))
}
