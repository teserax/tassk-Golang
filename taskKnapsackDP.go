package main

import "fmt"

type item struct {
	weight int
	price  int
}

func main() {
	backpackСapacity := 6

	bow := item{3, 4}
	por := item{3, 3}
	mat := item{2, 2}
	var bpPrice, bbPrice, pmPrice int
	var first, second, third bool
	if bow.weight+por.weight <= backpackСapacity {
		first = true
		bpPrice = bow.price + por.price
		fmt.Printf("optimal 1 %v", first)
	}
	if bow.weight+mat.weight <= backpackСapacity {
		second = true
		bbPrice = bow.price + mat.price
		fmt.Printf("optimal 2")
	}
	if por.weight+mat.weight <= backpackСapacity {
		third = true
		pmPrice = por.price + mat.price
		fmt.Printf("optimal 3")
	}
	if first == true && second == true && third == true {
		if bpPrice > bbPrice && bpPrice > pmPrice {
			fmt.Printf("optimal bow, por ")
		}
		if bpPrice < bbPrice && bpPrice < pmPrice {
			fmt.Printf("optimal %d + %d", por, mat)
		}

	}
}
