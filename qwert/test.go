package main

import "fmt"

// Person is a struct that keeps info about person's age
type Person struct {
	Age uint8
}
type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	agePopularity := make(map[uint8]int, 0)

	for _, ages := range pl {

		if _, ok := agePopularity[ages.Age]; ok {
			agePopularity[ages.Age]++
		} else {
			agePopularity[ages.Age] += 1
		}
	}

	return agePopularity
}
func main() {
	p1 := Person{25}
	p2 := Person{46}
	p3 := Person{25}
	personAges := PersonList{
		p1, p2, p3,
	}
	m := personAges.GetAgePopularity()
	fmt.Println(m)
}
