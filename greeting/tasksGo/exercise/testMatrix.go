package main

import "fmt"

func main() {
	m := [4][5]int{}

	N := len(m)
	M := 5

	var i, j, ibeg, iend, jbeg, jend int
	count := 1

	for count <= N*M {
		m[i][j] = count
		if i == ibeg && j < M-jend-1 {
			j++
		} else if j == M-jend-1 && i < N-iend-1 {
			i++
		} else if i == N-iend-1 && j > jbeg {
			j--
		} else {
			i--
		}
		if i == ibeg+1 && j == jbeg && jbeg != M-jend-1 {
			ibeg++
			iend++
			jbeg++
			jend++
		}
		count++
	}
	for _, row := range m {
		fmt.Printf("%3d\n", row)
	}
}
