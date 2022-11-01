package main

//Calculating the sum of numbers in a slice
import (
	"fmt"
	"time"
)

func main() {
	s1 := []int{-2, 3, -10, 7, 9, -3, 8, -6, 5, 4, -1}
	s2 := []int{-8, 6, -4, -9, 2, 10, -7, 3, 1, -5, 9} //------------------------
	numberCh := make(chan int)
	// resultCh := make(chan int)

	go rangeNumbers(s1, numberCh)
	go rangeNumbers(s2, numberCh) //-----------------------
	fmt.Println(<-numberCh)
	fmt.Println("ura")
	// summ := <-resultCh
	// summ2 := <-resultCh
	// fmt.Println("sum->", summ, summ2)
}

func rangeNumbers(s []int, numberCh chan int) {
	for _, val := range s {
		numberCh <- val
		fmt.Println(val)
		time.Sleep(1 * time.Millisecond)
	}

}

func sumNumber(numberCh chan int) {
	result := 0
	for val := range numberCh {
		result += val

	}

	fmt.Println(result)

}
