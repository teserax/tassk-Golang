/*-Написать программу,
    которая в зависимости от характера ветра выдает сообщение о его скорости
от 1до 4 м/с - слабый (1);
от 5-10 м/c -умеренный (2);
от 9-18 м/c - сильный (3);
больше 19 м/c - ураганный (4).*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func windPower() int {
	s := []int{1, 2, 3, 4}
	num := s[rand.Intn(3)]
	return num
}
func weatherForecast(wind func() int) {

	switch wind() {
	case 1:
		fmt.Println("Ветер от 1до 4 м/с - слабый")
	case 2:
		fmt.Println("Ветер от 5-10 м/c -умеренный")
	case 3:
		fmt.Println("Ветер от 9-18 м/c - сильный ")
	case 4:
		fmt.Println("Ветер больше 19 м/c - ураганный ")

	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	weatherForecast(windPower)
}
