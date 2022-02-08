package main

import "fmt"

func costBancnot(rest int) {
	slice := []int{}
	var w, a, b, c, e, d int
	//naminal
	//w=500lei
	//a= 100 lei
	//b= 50 lei
	//c= 10 lei
	//e= 5 lei
	//d= 1 lei

	for rest > 0 {
		x := rest % 10
		rest = rest / 10
		slice = append(slice, x)
	}
	for i, val := range slice {
		if i == 0 {
			d = val - 5
			e++
		}
		if i == 1 {
			c = val - 5
			b++
		}
		if i == 2 {
			a = val - 5
			w++
		}
	}

	fmt.Printf(" в банкнотах \n по 500 lei %dштук\n по 100 lei %dштук\n по 50lei %dштук\n по 10lei %dштук\n по 5lei %dштук\n по 1lei %dштук\n", w, a, b, c, e, d)

}

func main() {
	sum := 1000

	rest := sum - 334
	fmt.Printf("Surrender %d\n", rest)
	costBancnot(rest)
}
