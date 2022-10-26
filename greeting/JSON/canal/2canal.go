package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
		time.Sleep(1 * time.Second)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	test := make(chan string, 2)
	test <- "hello"
	go testString(test)

}
func testString(ch chan string) {
	time.Sleep(1 * time.Second)
	str := <-ch
	fmt.Println(str)
}
func sleepyGopher(id int, c chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("...snore...", id)
	c <- id
}
