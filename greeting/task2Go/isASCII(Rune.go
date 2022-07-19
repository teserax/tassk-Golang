package main

import "fmt"

func isASCII(s string) bool {

	for _, word := range s {

		if int32(word) > 256 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isASCII(" abc1"))
	fmt.Println(isASCII("хай"))
	fmt.Println(isASCII("hello \U0001F970"))
}
