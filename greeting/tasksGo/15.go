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
	} else if i == 0 {
		s := nums[1:]
		return s
	} else {
		return nums
	}
}

func main() {
	//nums := [5]int{}
	fmt.Println(Remove([]int{1, 2}, -5))
	fmt.Println(Remove([]int{1, 2, 3}, 0))

	fmt.Println(Remove([]int{1, 2, 3}, 2))
	fmt.Println(Remove([]int{1, 2, 3}, 3))
	fmt.Println(Remove([]int{1, 2, 3}, 5))
}

/*

         0             3-1
	nums[i] = nums[len(nums)-1]

	return nums[:len(nums)-1]


*/
