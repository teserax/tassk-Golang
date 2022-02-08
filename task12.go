// Описание задачи
// Имитируется работа кассира. С клавиатуры вводятся
// стоимость покупки и сумма денег, полученных от
// покупателя. Кассир рассчитывает сумму сдачи, которую
// он должен вернуть покупателю. Сумма сдачи должна
// быть сформирована из минимального количества
// банкнот

package main

import "fmt"

func main() {
	var price, sum int
	fmt.Println("Enter the price of the item ")
	fmt.Scan(&price)
	fmt.Println("Enter the sum ")
	fmt.Scan(&sum)

	fmt.Println(price, sum)

}
