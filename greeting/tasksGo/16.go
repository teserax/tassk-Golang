/*
Написать программу которая по коду города и
длительности переговоров вычисляет их стоимость и
результат выводит на экран
Одесса-код 048 стоимость 15грн
Киев-код   044 стоимость 18грн
Харьков-код 046 стоимость 13грн
Винница-код 045 стоимость 11грн
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randCity() int {
	s := []int{48, 44, 46, 45}
	num := rand.Intn(4)
	return s[num]
}
func randCallMinutes() int {
	return rand.Intn(19) + 1
}
func callСost(code func() int, minutes func() int) {
	min := minutes()
	cod := code()
	switch cod {
	case 44:
		fmt.Printf("Звонок совершен из Киева,длительность звонка  %d минут и стоимость состовляет %d гривень \n", min, min*18)

	case 45:
		fmt.Printf("Звонок совершен из Винница,длительность звонка  %d минут и стоимость состовляет %d гривень\n", min, min*11)

	case 48:
		fmt.Printf("Звонок совершен из Одесса,длительность звонка  %d минут и стоимость состовляет %d гривень\n", min, min*15)

	case 46:
		fmt.Printf("Звонок совершен из Харьков,длительность звонка  %d минут и стоимость состовляет %d гривень\n", min, min*13)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		callСost(randCity, randCallMinutes)
	}
}
