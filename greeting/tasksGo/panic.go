package main

import "fmt"

func main() {
	fmt.Println("main:defer") //3
	someFunc()                //1
}
func someFunc() {
	defer fmt.Println("someFunc:call") //
	panic("something get wrong")
}
