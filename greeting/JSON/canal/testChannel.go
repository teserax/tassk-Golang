// - первый этап: N1 количество горутин генерируют случайные числа с интервалом N2 секунд
//    в каком то диапазоне [a,b] и отправляет в канал 1
// - второй этап: M1 гороутин берут эти числа из канала 1
//    и делают какую то обработку их, например находят все делители этих чисел
//    и отправляют json структуру (число и его делители) в канал 2
// - третий этап: K1 гороутин получают json и выделяют
//    максимальное число переданной на 1вом этапе.
//--программа заканчивает работу после В секунд или если сгеренируется на первом этапе максимальное число с интервала --

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type exitEvent struct {
	exit  bool
	mutex sync.RWMutex
}

var (
	exitEventVar exitEvent
)

func randNum(min, max int) int {
	return rand.Intn(max-min) + min
}
func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	col := [10]struct{}{}
	ch := make(chan int)
	for range col {
		wg.Add(1)
		go func() {
			result := randNum(1, 9)
			wg.Done()
			if result == 7 {
				exitEventVar.exit = false
				fmt.Println("asia nara")
				return
			} else {
				ch <- result
				fmt.Println(result)
			}
			go func() {
				defer close(ch)
				defer wg.Wait()
			}()
		}()

	}
	for v := range ch {
		fmt.Println(v)
	}

}
