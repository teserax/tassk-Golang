//заполнить матрицу построчно от начала и до конца сверху вниз-снизу вверх в цикле

package main

import "fmt"

func main() {
	const N = 10
	const M = 5
	matrix := [N][M]int{}

	var lift, step, count int
	count = 1
	for {
		//будем создавать два цикла один поднимает наш "count"другой опускает

		for lift = 0; lift < N; lift++ { //опускаем лифт вниз
			if step >= M {
				step = M - 1
			}
			matrix[lift][step] = count
			count++

		}
		step++ //шаг в соседниий подьезд
		if count >= N*M {
			break
		}
		for lift = N - 1; lift >= 0; lift-- { //поднимаем лифт вверх
			if step >= M {
				step = M - 1
			}
			matrix[lift][step] = count
			count++
		}
		step++

	}
	for _, row := range matrix {
		fmt.Printf("%2d\n", row)
	}
}
