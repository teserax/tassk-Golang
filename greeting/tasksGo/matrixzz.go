//заполнить матрицу построчно от начала и до конца

package main

import "fmt"

func main() {
	const N = 7
	const R = 7
	m := [N][R]int{}
	var i, j, step int
	count := 1

	for count <= N*R {
		m[i][j] = count
		if j < len(m)-1-step && i == step { //если верно двигаемся  с лева на право
			j++
		} else if j == len(m)-1-step && i < len(m)-1-step { //если верно двигаемся сверху вниз
			i++
		} else if i == len(m)-1-step && j > step { //если верно с права на влево
			j--
		} else if j == step && i > step { // с низу в вверх
			i--
		}
		if i == step+1 && j == step { //
			step++
		}
		count++
	}
	for _, val := range m {
		fmt.Printf("%2d\n", val)
	}

}
