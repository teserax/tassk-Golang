// Описание задачи
// Имитируется работа кассира. С клавиатуры вводятся
// стоимость покупки и сумма денег, полученных от
// покупателя. Кассир рассчитывает сумму сдачи, которую
// он должен вернуть покупателю. Сумма сдачи должна
// быть сформирована из минимального количества
// банкнот.
package main

import "fmt"

func costBancnot(rest int) {
	slice := []int{}
	var w, a, b, c, e, d int
	//naminal
	//w=500lei
	//a= 100 lei
	//b= 50 lei
	//c= 10 lei
	//e= 5 lei
	//d= 1 lei

	for rest > 0 {
		x := rest % 10
		rest = rest / 10
		slice = append(slice, x)
	}
	for i, val := range slice {
		if i == 0 {
			if val > 5 {
				d = val - 5
				e++
				fmt.Printf(" банкнота 5lei %dштук\n", e)
			} else {
				for val > 0 {
					d++
					val--
				}
			}
			fmt.Printf("  банкнота 1lei %dштук\n", d)
		}
		if i == 1 {
			if val > 5 {
				c = val - 5
				b++
				fmt.Printf("  банкнота50 lei %dштук\n", b)
			} else {
				for val > 0 {
					c++
					val--
				}
			}
			fmt.Printf("  банкнота 10 lei %dштук\n", c)
		}

		if i == 2 {
			if val > 500 {
				a = val - 5
				w++
				fmt.Printf("  банкнота 500 lei %dштук\n", w)
			} else {
				for val > 0 {
					a++
					val--
				}
			}
			fmt.Printf(" банкнота 100 lei %dштук\n", a)
		}

	}
	//fmt.Printf("  по 10lei %dштук\n по 5lei %dштук\n по 1lei %dштук\n", w, a, b, c, e, d)

}

func main() {
	var price, sum, rest int
	for {
		fmt.Println("Введите цену товара ")
		fmt.Scan(&price)
		if price == 0 || price < 0 {
			fmt.Println("Введенная сумма не некорректна")
		}
		if price > 0 {
			break
		}
	}
	for {
		fmt.Println("Введите полученую сумму ")
		fmt.Scan(&sum)
		if sum == 0 || sum < 0 {
			fmt.Println("Введенная сумма не некорректна")
		}
		if sum > 0 {
			break
		}

	}
	rest = sum - price
	fmt.Printf("Surrender %d\n", rest)
	costBancnot(rest)

}
