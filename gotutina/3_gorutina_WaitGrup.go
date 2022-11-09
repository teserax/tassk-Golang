package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 100, "Number of gorutines")

	flag.Parse()

	count := *n
	fmt.Printf("Going to create %d gorutines.\n", count)
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Printf("%d", x)
		}(i)
	}
	fmt.Printf("%#v\n", wg)
	wg.Wait()
	fmt.Println("\nExiting...")
}
