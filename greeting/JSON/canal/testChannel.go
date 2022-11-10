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

var wg sync.WaitGroup

type divisorValue struct {
	value   int
	divisor []int
}

var ch = make(chan int)

func randomNumberGeneration(numberOfGourutin, interval int) {
	out := make(chan int)
	result := map[int][]int{}
	for i := 0; i < numberOfGourutin; i++ {
		go func() {
			time.Sleep(time.Duration(interval))
			num := rand.Intn(10) + 1
			out <- num

		}()

		select {
		case _, ok := <-out:
			if !ok {
				fmt.Println("канал закрыт ")
				return
			}
		case n := <-out:

			for i := 1; i < 6; i++ {
				result[n] = append(result[n])
				if n%i == 0 {

					if _, exist := result[n]; exist {
						result[n] = append(result[n], i)
					}

				}

			}

		}

	}
	fmt.Println(result)
	defer close(out)
	return
}
func main() {
	rand.Seed(time.Now().UnixNano())

	randomNumberGeneration(7, 1)

}
