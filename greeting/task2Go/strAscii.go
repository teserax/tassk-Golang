package main

import (
	"fmt"
)

func shiftASCII(s string, step int) string {
	str := []byte{}
	if step == 0 {
		return s
	}
	ss := []byte(s)
	fmt.Println(ss)
	//fmt.Println("gggg", (step))
	for i := 0; i < len(s); i++ {
		str = append(str, byte(int(s[i])+step))

	}

	return string(str)
}

// func shiftASCII(s string, step int) string {
// 	if step == 0 {
// 		return s
// 	}

// 	shifted := make([]byte, len(s))
// 	for i := 0; i < len(s); i++ {
// 		shifted[i] = byte(int(s[i]) + step)
// 	}

// 	return string(shifted)
// }

func main() {
	fmt.Println(shiftASCII("bcd2", -490))
	// fmt.Println(shiftASCII("hi", 10))
	// fmt.Println(shiftASCII("abc", 256))
}
