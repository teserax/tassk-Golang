package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var CLOSEA = false
var DATA = make(map[int]bool)

func random(min, max int) int { //генерация случайного числа
	return rand.Intn(max-min) + min
}
func first(min, max int, out chan<- int) { //функция первого уровня
	for {
		if CLOSEA { //если правда(true)
			close(out) //закрыть канал
			return
		}
		out <- random(max, min) //если нет то посылаем полученное число в канал out
	}
}
func second(out chan<- int, in <-chan int) { //принимаем два канала out-запись вканал,,,,,in-чтение из канала
	for x := range in { //пребераем "читаем" канал in
		fmt.Print(x, " ") //выводим содержимое
		_, ok := DATA[x]  //проверяем если ключь х существует close= да
		if ok {
			CLOSEA = true
		} else { //если нет то записываем ключь и значение в нашу мапу
			DATA[x] = true
			out <- x //значение посылаем в канал для дольнецшей обработки
		}
	}
	fmt.Println()
	close(out) //закрываем канал
}
func third(in <-chan int) { // принимаем канал тольео чтение
	sum := 0
	for x2 := range in { //пребераем "читаем" канал in
		sum = sum + x2 //суммируем
	}
	fmt.Println("The sum of the random numbers is %d\n", sum) //выводим сумму всех чисел
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
		return
	}
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])
	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d\n", n1, n2)
		return
	}
	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)
	go first(n1, n2, A)
	go second(B, A)
	third(B)
}
