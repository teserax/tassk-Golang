package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := []int{}
	slice2 := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		number := -25 + rand.Intn((50)+1)
		slice = append(slice, number)
	}
	fmt.Println(slice)

	for _, val := range slice {
		if val <= 0 {
			continue
		}

		slice2 = append(slice2, val)
		slice = slice2
	}
	fmt.Println(slice)
}
