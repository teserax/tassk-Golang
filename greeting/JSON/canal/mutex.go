package main

import (
	"fmt"
	"sync"
)

type exitEvent struct {
	// то значение меняется если достигнуто максимальное число или через какое то время тикер отправил сигнал в канал тикера.
	// все горутины имеют проверку этого поля и если оно поменялось на true, все выходят.
	exit  bool // это на запись закрывается мютексом
	mutex sync.RWMutex
}

var (
	exitEventVar exitEvent
)

// first step
func gen(goroutineNumbers int, startNum int, EndNum int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < goroutineNumbers; i++ {
		wg.Add(1)
		go func() {
			// generate number
			defer wg.Done()
			genNumber := 5

			// check number with max value
			if genNumber == EndNum {
				out <- genNumber
				fmt.Println("tut")
				// set exitEvent
				exitEventVar.mutex.Lock() // this not block reading
				exitEventVar.exit = true  // write event to exit
				exitEventVar.mutex.Unlock()
				return
			}
			if exitEventVar.exit {
				fmt.Println("Exit")
				return
			}

			// return number
			out <- genNumber
		}()
	}
	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out
}

func main() {
	result := gen(10, 1, 5)
	for v := range result {
		fmt.Println(v)
	}
}
