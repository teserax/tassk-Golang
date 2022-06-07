package main

import (
	"fmt"
)

func Remove(nums []int, i int) []int {
	if len(nums) > i && i > 0 {
		s := nums[:i]
		d := nums[i+1:]
		s = append(s, d...)
		return s
	} else {
		return nums
	}
}

// BEGIN (write your solution here)
func main() {
	//nums := [5]int{}
	fmt.Println(Remove([]int{1, 2}, -5))
	fmt.Println(Remove([]int{1, 2, 3}, 0))

	fmt.Println(Remove([]int{1, 2, 3}, 2))
	fmt.Println(Remove([]int{1, 2, 3}, 3))
	fmt.Println(Remove([]int{1, 2, 3}, 5))
}
