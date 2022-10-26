package main

import "fmt"

func main() {

	matrix := [5][5]int{}
	value := 1
	step := 0
	for k := 0; k < len(matrix); k++ {

		// первый цикл заполнения слева на право
		for r := step; r < len(matrix)-step; r++ {
			matrix[step][r] = value
			value++
			fmt.Println(value)
		}
		step++
		for b := step; b < len(matrix)-step; b++ {
			matrix[b][len(matrix)-step] = value
			value++
		}

		for l := len(matrix) - step; l >= step; l-- {

			matrix[len(matrix)-step][l] = value
			value++
		}
		for t := len(matrix) - step; t >= step; t-- {
			matrix[t][0] = value
			value++

		}

	}

	for _, row := range matrix {

		fmt.Printf("%2d\n", row)
	}
}
