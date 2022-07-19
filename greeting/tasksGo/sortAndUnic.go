package main

import "fmt"

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	if len(userIDs) < 2 {
		return userIDs
	}
	for i := 0; i < len(userIDs); i++ {
		for j := 0; j < len(userIDs); j++ {
			if userIDs[i] < userIDs[j] {
				m := userIDs[i]
				userIDs[i] = userIDs[j]
				userIDs[j] = m
			}
		}
	}
	count := 0
	for i := 0; i < len(userIDs); i++ {
		if userIDs[count] != userIDs[i] {
			count++
			userIDs[count] = userIDs[i]
		}
	}
	return userIDs[:count+1]
}

func main() {

	fmt.Println(UniqueSortedUserIDs([]int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}))
	fmt.Println(UniqueSortedUserIDs([]int64{}))
	fmt.Println(UniqueUserIDs([]int64{55, 55, 33, 22}))
}
