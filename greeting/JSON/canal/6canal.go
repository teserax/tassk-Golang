//Напишите функцию, которая возьмет исходную строку,
//разобьет ее на слова и посчитает количество цифр в каждом слове.
// Причем подсчет для каждого слова запустите в отдельной горутине.
package main

import (
	"fmt"
	"strings"
	"unicode"
)

// var wg sync.WaitGroup

func findInt(str string) {
	words := strings.Fields(str)
	count := 0

	for _, word := range words {
		go func() {
			// wg.Add(1)
			for _, num := range word {
				if unicode.IsDigit(num) {

					fmt.Println(string(num))
					count++
				}

			}
			// wg.Done()
		}()

	}

	fmt.Printf("number of integer in text %d\n", count)

}

func main() {

	str := "1 a2a  4 b c 5 ,ww77ewe   werwrwr wrwrwerwer"

	findInt(str)
	// wg.Wait()

}
