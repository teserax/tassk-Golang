package main

import (
	"fmt"
	"math/rand"
	"time"
)

var winner = make(chan string) // инициализация канала winner
var echo = make(chan struct{}) //инициализация канала echo

func game_table(s string) {
	r := rand.Intn(5) + 1                              //  получение рандомного числа для временнова отрезка
	fmt.Println(s, r)                                  // имя и полученное число
	t := time.NewTimer(time.Duration(r) * time.Second) //временной отрезок
	for {
		select { // выбор условии
		case <-t.C: // первый кейс временной отрезок
			fmt.Println(s, "press botton") // имя перврго кто нажал
			close(echo)                    //закрытие канала
			defer close(winner)            // зарытие канала дефером
			winner <- s                    // посылаем в канал имя победителя
			defer t.Stop()                 // останавливаем дефером  таймер для имени(вошедшего первым)
			return

		case <-echo: // второй кейс если канал echo  закрыт то
			fmt.Println("Закрыто для ", s) // обьявляем для кого закрылся канал
			defer t.Stop()                 //останавливаем дефером  таймер для имени(вошедшего )
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	go game_table("Max")    //запуск горутины Макс
	go game_table("Vova")   //запуск горутины Вова
	go game_table("Valera") //запуск горутины Валера

	fmt.Println("Wait")
	s := <-winner                    //присваеваем имя победителя полученное из канала
	fmt.Println("The winner is ", s) //обьевляем победителя

	<-time.NewTicker(time.Duration(10) * time.Millisecond).C
}
