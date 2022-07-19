package main

import "fmt"

func UniqueUserIDs(userIDs []int64) []int64 {
	m := map[int64]int64{}
	s := []int64{}
	for _, val := range userIDs {
		if _, ok := m[val]; ok {
			continue
		}
		m[val] = 1
		s = append(s, val)
	}
	return s

}

func main() {

	fmt.Println(UniqueUserIDs([]int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}))
}

// a.Equal([]int64{},
// a.Equal([]int64{10},
// a.Equal([]int64{55},
// a.Equal([]int64{55, 33, 22},
// a.Equal([]int64{55, 2, 88, 33, 103},
