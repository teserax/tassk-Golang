package main

import (
	"fmt"
	"strings"
)

func duplicate_count(s1 string) (c int) {
	m := map[string]int{}

	for _, val := range strings.ToUpper(s1) {

		m[(string(val))]++
		if m[(string(val))] == 2 {
			c++
		}

	}
	return
}
func main() {
	str := "abcdeaB11"
	fmt.Println(duplicate_count(str))

}
