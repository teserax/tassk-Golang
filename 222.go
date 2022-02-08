package main

import "fmt"

func main() {
	b := 9 - 5
	i := 0
	if b > 0 {

		for {
			b--
			i++
			if b == 0 {
				break
			}

		}
	}
	fmt.Print(i)
}
