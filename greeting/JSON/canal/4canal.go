package main

import (
	"fmt"
	"math/rand"
	"time"
)

var CLOSEA = false
var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	}

}
func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, "  ")
		_, ok := DATA[x]
		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}
func third(in <-chan int) {
	var sum int
	for x := range in {
		sum += x
	}
	fmt.Printf("summa of random numbers is %d", sum)
}
func main() {

	n1 := 2
	n2 := 10

	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)
	go first(n1, n2, A)
	go second(B, A)
	third(B)

}
