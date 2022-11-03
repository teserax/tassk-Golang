//Напишите функцию, которая возьмет исходную строку,
//разобьет ее на слова и посчитает количество цифр в каждом слове.
// Причем подсчет для каждого слова запустите в отдельной горутине.
package main

import (
	"fmt"
	"sync"
	"unicode"
)

var wg sync.WaitGroup

func findInt(str string) {
	count := 0
	for _, val := range str {
		if unicode.IsDigit(rune(val)) {
			count++
		}

	}
	fmt.Printf("number of integer in text is %d\n", count)
	wg.Done()
}

func main() {

	str := "1 a2  4 b c 5 ,wwewe   werwrwr wrwrwerwer"
	str2 := "stop 2 and look 3 jn 4 moon"
	str3 := "афые 2 ар6кlook 3 ууке 4 moo7n"
	wg.Add(1)
	go findInt(str)
	wg.Add(1)
	go findInt(str2)
	wg.Add(1)
	go findInt(str3)
	wg.Wait()

}
