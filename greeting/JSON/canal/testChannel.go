package main

import "fmt"

func main() {
	ch := make(chan int)

	go Sum(5, 4, ch)

	result := <-ch
	fmt.Printf("our result from channel %d\n", result)
	go Sum(1, 4, ch)
	ourResult := <-ch
	fmt.Printf("our result from channel %d", ourResult)
}
func Sum(a, b int, ch chan int) {
	ch <- a + b
}
