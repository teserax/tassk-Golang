package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int { //функция читающая любое количество цифр
	out := make(chan int) // канал int
	go func() {           //горутина
		for _, n := range nums { //цикл for range переберает числа преходящие в функцию
			out <- n // и посылает в канал out
		}
		close(out) //закрытие канала
	}()
	return out // возвращение значения функции
}
func sg(in <-chan int) <-chan int { //функция читающая из канала и возвращант канал

	out := make(chan int) // канал int
	go func() {           //горутина
		for n := range in { //цикл for range читаем данные из канала
			out <- n * n //посылаем квадрат числа из канала
		}
		close(out) //закрытие канала
	}()
	return out // возвращение значения функции
}
func main() {
	in := gen(2, 3)                //функция и два значения()...возвращаемое значение виде канала присваеваем переменной  in
	c1 := sg(in)                   //функция принимает канал и возвращает значения в канал
	c2 := sg(in)                   //функция принимает канал и возвращает значения в канал
	for n := range merge(c1, c2) { //цикл for range для функции merge с передачей двух каналов
		fmt.Println(n) //вывод работы функции  merge
	}
}
func merge(cs ...<-chan int) <-chan int { // merge принимает любое количество каналов и возвращает канал

	var wg sync.WaitGroup          // пакет sync.WaitGroup
	out := make(chan int)          //// канал int
	output := func(c <-chan int) { //анонимная функция принимающая значение ввиде канала
		for n := range c { //цикл for range для функции  с передачей в канал out
			out <- n
		}
		wg.Done() //работа счетчиков каналов группы WaitGroup "минус"
	}

	wg.Add(len(cs))        //работа счетчиков каналов группы WaitGroup "плюс" c lenght"длина cs"
	for _, c := range cs { //цикл for range для значении  cs с передачей в канал
		go output(c) //передачей в канал
	}
	go func() { // анонимная функция для закрытия группы счетчиков WaitGroup
		wg.Wait()  //ожидание окончания всех запущеных горутин
		close(out) //закрытие канала out
	}()
	return out //возвращения значения out
}
