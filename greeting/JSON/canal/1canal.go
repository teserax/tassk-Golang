package main

//Calculating the sum of numbers in a slice
import "fmt"

func main() {
	s := []int{2, 8, 6, 9, 1, 7, 2}
	number := make(chan int)
	// result := make(chan int)

	go rangeNumbers(s, number)
	// go sumNumber(number, result)
	summ := <-number
	fmt.Println(summ)
}
func rangeNumbers(s []int, ch chan int) {
	result := 0
	for _, val := range s {
		result += val
		ch <- result
	}
	select {
	case ch <- result:
		ch <- result
	default:
		close(ch)
	}

}

// func sumNumber(m, result chan int) {
// 	r := 0
// 	for val := range m {
// 		r += val
// 	}
// 	result <- r
// }
