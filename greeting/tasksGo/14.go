package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func randomString(a int) string {
	bytes := make([]byte, a)
	for i := 0; i < a; i++ {
		bytes[i] = byte(randomInt(60, 90))
	}
	return string(bytes)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(randomString(10))
}
