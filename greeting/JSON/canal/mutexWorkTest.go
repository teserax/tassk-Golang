package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var do bool
var counter int = 0 //  общий ресурс
func main() {
	var wg = new(sync.WaitGroup)
	ch := make(chan bool) // канал
	numCh := make(chan int)
	var mutex sync.Mutex // определяем мьютекс
	do = true
	for do {

		select {
		case <-ch:
			for i := 1; i < 10; i++ {
				wg.Add(1)
				go func() {
					num := rand.Intn(10) + 1
					work(num, ch, wg, &mutex)
					numCh <- num
					fmt.Println(num)

				}()
			}
		default:
			fmt.Println("we close")
			do = false
			close(numCh)
			//

		}

	}
	for i := range numCh {
		fmt.Println(i)
	}
	go func() {
		wg.Wait()
	}()
	fmt.Println("The End")
}

func work(number int, ch chan bool, wg *sync.WaitGroup, mutex *sync.Mutex) {
	wg.Done()
	mutex.Lock() // блокируем доступ к переменной counter
	if number == 3 {
		fmt.Println("NUMBER FOUND", number)

		ch <- false

	}
	mutex.Unlock() // деблокируем доступ
	ch <- true
}
