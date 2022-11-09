//заполнить матрицу построчно от начала и до конца зейкой

package main

import "fmt"

func main() {

	matrix := [6][6]int{}
	count := 0
	for i := 0; i < len(matrix); i++ {
		step := len(matrix)

		for j := 0; j < len(matrix); j++ {

			//matrix[i][j] = i*len(matrix) + j + 1

			count++
			if count%5 == 0 {
				matrix[i][step-1] = i*len(matrix) + j + 1

			}

		}
	}

	for _, val := range matrix {
		fmt.Printf("%2d\n", val)
	}

}
