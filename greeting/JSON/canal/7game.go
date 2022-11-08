package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var winner = make(chan string) // инициализация канала winner
var echo = make(chan struct{}) //инициализация канала echo
var winnerName string
var done = make(chan struct{}) // инициализация канала winner
func game_table(s string) {
	r := rand.Intn(5) + 1                              //  получение рандомного числа для временнова отрезка
	fmt.Println(s, r)                                  // имя и полученное число
	t := time.NewTimer(time.Duration(r) * time.Second) //временной отрезок
	defer func() {
		t.Stop()
		wg.Done()
	}()
	select { // выбор условии
	case <-t.C: // первый кейс временной отрезок
		//закрытие канала

		winner <- s // посылаем в канал имя победителя

	case _, ok := <-echo:
		if !ok {
			return
		} // второй кейс если канал echo  закрыт то
		// обьявляем для кого закрылся канал

	}
	return
}
func readWinner() {
	var i int
	// читаем все из канала winner.
	// чтение блокирующее, пока не закроется winner
	// нужно читать в цикле при ситуации когда время одинаково у всех иначе они все стопорятся на этапе отправки в канал.
	for val := range winner {
		// как только прочитали первого победителя, сигналим другие горутины чтоб закрылись
		// как правило это при ситуации когда у всех время разное
		if i == 0 {
			winnerName = val
			close(echo)
		}
		i++
	}
	done <- struct{}{}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg.Add(3)
	go readWinner()

	go game_table("Max")    //запуск горутины Макс
	go game_table("Vova")   //запуск горутины Вова
	go game_table("Valera") //запуск горутины Валера

	fmt.Println("Wait")
	wg.Wait()
	close(winner) //присваеваем имя победителя полученное из канала
	//обьевляем победителя

	//<-time.NewTicker(time.Duration(10) * time.Millisecond).C
	<-done
	fmt.Println("The winner is ", winnerName)
}
