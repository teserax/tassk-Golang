/*Дано целое число k (1<k<365). Определить, каким днем недели (поне-
дельник,вторник,воскресенье)являетсяk-день  невысокосного  года,если
1 января - понедельник.*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dayOfWeek(date int) {
	for date >= 8 {

		if date >= 8 {
			date = date - 7
		}
		if date <= 0 {
			fmt.Println("Error")
		}

	}
	switch date {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("Среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница")
	case 6:
		fmt.Println("Суббота")
	case 7:
		fmt.Println("Воскресение")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dayOfWeek(rand.Intn(365))

}
