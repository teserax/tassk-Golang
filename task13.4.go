//Случайные числа и символы
//https:gospodaretsva.com/random.html
/*случайное целое число,
случайное вещественное число,
случайный символ.
package main*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func randoNum() []int {
// 	rand.Seed(time.Now().UnixNano())
// 	var a, b int
// 	fmt.Print("Enter number max ")
// 	fmt.Scan(&a)

// 	fmt.Print("Enter number min ")
// 	fmt.Scan(&b)
// 	s := []int{}
// 	for i := 0; i < 10; i++ {
// 		number := b + rand.Intn((a-b)+1)
// 		s = append(s, number)
// 	}
// 	return s
// }
func main() {

	// fmt.Printf("Number random in range %d", randoNum())
	rand.Seed(time.Now().UnixNano())
	var a, b string
	a = "a"
	b = "x"
	var min, max int
	min = int(a[0])
	max = int(b[0])
	slice := []string{}
	for i := 0; i < 4; i++ {
		number := min + rand.Intn(int((max-min))+1)
		slice = append(slice, string(number))

	}
	fmt.Printf("%s\n", slice)
}
