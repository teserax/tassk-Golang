package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Test() {
	fmt.Println("hi")
}
func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
