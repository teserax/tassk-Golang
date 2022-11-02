package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup //пакет позволяющий создовать группу ожидания для горутин(счетчик выполнения)
	wg.Add(2)             //начало (запуск счетчика).. выполнения "2"горутин..указываем в скобках Add

	go func() { //запуск горутины под средством аннонимной функции
		defer wg.Done()         // Done()отнимаем от счетчика -1 ...указаного выше в Add(2)-1, гарантирует, что горутина уменьшит счетчик перед выходом
		say(1, "go is awesome") //вызываем функцию say
	}()
	go func() { //запуск горутины под средством аннонимной функции
		defer wg.Done()        // отнимаем от счетчика -1 ...указаного выше в Add(1)-1
		say(2, "cat are cute") //вызываем функцию say
	}()
	wg.Wait() //ожидаем онуления на шего счетчикa

}
func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("worker #%d says: %s...\n", id, word)
	}
	dur := time.Duration(rand.Intn(100)) * time.Millisecond
	time.Sleep(dur)
}
