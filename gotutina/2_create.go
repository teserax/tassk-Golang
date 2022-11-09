package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	n := flag.Int("n", 100, "Number of gorutines")

	flag.Parse()

	count := *n
	fmt.Printf("Going to create %d gorutines.\n", count)

	for i := 0; i < count; i++ {

		go func(x int) {

			fmt.Printf("%d", x)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}
