// - первый этап: N1 количество горутин генерируют случайные числа с интервалом N2 секунд
//    в каком то диапазоне [a,b] и отправляет в канал 1
// - второй этап: M1 гороутин берут эти числа из канала 1
//    и делают какую то обработку их, например находят все делители этих чисел и
//    отправляют json структуру (число и его делители) в канал 2
// - третий этап: K1 гороутин получают json и выделяют
//    максимальное число переданной на 1вом этапе.
//--программа заканчивает работу после В секунд или если сгеренируется на первом этапе максимальное число с интервала --

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type NumberDivisors struct {
	Number  int
	Divisor []int
}

func main() {
	var inter, workTime time.Duration
	inter = 1
	workTime = 1
	ch := make(chan int)
	rand.Seed(time.Now().UnixNano())
	interval := time.NewTicker(time.Second * inter)
	timer := time.After(time.Second * workTime)
	select {

	case <-timer:
		close(ch)
		interval.Stop()
		fmt.Println("timeout")
	case <-interval.C:
		caller := make([]struct{}, 10)
		var wg = new(sync.WaitGroup)
		for range caller {
			wg.Add(1)
			go func() {
				gn := rand.Intn(20) + 1
				ch <- gn
				wg.Done()
			}()
		}
		go func() {
			defer close(ch)
			wg.Wait()

		}()
	}
	ch2 := make(chan []byte)
	var wg1 = new(sync.WaitGroup)
	for v := range ch {
		wg1.Add(1)
		go divisorSearch(v, ch2, wg1)

	}
	go func() {
		defer close(ch2)
		wg1.Wait()

	}()
	for v := range ch2 {
		fmt.Println(string(v))
	}

}
func divisorSearch(n int, ch chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	numbers := NumberDivisors{}
	for i := 1; i < n+1; i++ {

		if n%i == 0 {
			numbers.Number = n
			numbers.Divisor = append(numbers.Divisor, i)
		}
	}

	myJson, err := json.Marshal(numbers)
	if err != nil {
		log.Fatal(err)
	}
	ch <- myJson

}
