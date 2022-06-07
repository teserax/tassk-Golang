package main

import "fmt"

func main() {
	str := "{}[]()"
	for _, val := range str {
		fmt.Println(val)
	}

}
