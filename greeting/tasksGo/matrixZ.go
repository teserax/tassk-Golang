//заполнить матрицу построчно от начала и до конца

package main

import "fmt"

func main() {

	matrix := [7][7]int{}

	right := 0
	// bottom := 0
	left := 0
	top := 0

	value := 0
	// по сути мы по кругу обходим со всех сторон, поэтому нам надо 4 цикла. И каджый из них работает до половины матрицы
	for k := 0; k < len(matrix)/2+1; k++ {
		// первый цикл заполнения слева на право
		for r := left; r < len(matrix)-right; r++ {
			matrix[top][r] = value
			// temp = append(temp, value)
			value++
			fmt.Println(value)
		}
		// // fmt.Println(temp)
		// //fmt.Printf("right: %d, left: %d, top: %d, bottom: %d, values: %+v\n", right, left, top, bottom, temp)
		// top++

		// temp = []int{}
		// for b := top; b < len(matrix)-bottom; b++ {
		// 	matrix[b][len(matrix)-right-1] = value //m[1]
		// 	// temp = append(temp, value)
		// 	value++
		// }
		// // fmt.Println(temp)
		// right++
		// //fmt.Printf("right: %d, left: %d, top: %d, bottom: %d, values: %+v\n", right, left, top, bottom, temp)
		// temp = []int{}
		// for l := len(matrix) - right - 1; l >= left; l-- {
		// 	matrix[len(matrix)-right][l] = value
		// 	// temp = append(temp, value)
		// 	value++
		// }
		// // fmt.Println(temp)
		// bottom++

		// //fmt.Printf("right: %d, left: %d, top: %d, bottom: %d, values: %+v\n", right, left, top, bottom, temp)
		// temp = []int{}
		// for t := len(matrix) - bottom - 1; t >= top; t-- {
		// 	matrix[t][left] = value
		// 	temp = append(temp, value)
		// 	value++
		// }
		// //fmt.Println(temp)
		// left++
		// //fmt.Printf("right: %d, left: %d, top: %d, bottom: %d, values: %+v\n", right, left, top, bottom, temp)
		// temp = []int{}
	}
	for _, val := range matrix {
		fmt.Printf("%2d\n", val)
	}

}
