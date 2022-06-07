package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(filter(slice))

}

func filter(s []int) []int {
	sl := []int{}
	for _, val := range s {

		if predicat(val) == true {
			sl = append(sl, val)
		}
	}
	return sl
}

func predicat(val int) bool {
	if val%2 == 0 {
		return true
	}
	return false
}
