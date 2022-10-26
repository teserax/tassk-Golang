//             Описание задачи
// Имитируется работа кассира. С клавиатуры вводятся
// стоимость покупки и сумма денег, полученных от
// покупателя. Кассир рассчитывает сумму сдачи, которую
// он должен вернуть покупателю. Сумма сдачи должна
// быть сформирована из минимального количества
// банкнот.

package main

import "fmt"

// func change(sum int, s struct) { //функции расчета сдачи
// 	var totalCash int
// 	ch := map[int]int{}
// 	for nom, val := range m {
// 		totalCash += (nom * val)
// 	}
// 	if totalCash < sum {
// 		fmt.Println(" К сожелению в кассе недостатачно денег приходите позже")
// 	} else {
// 		for nom, val := range m {

// 		}

// 	}
// 	return ch
// }

type Cash struct {
	nominal int
	coint   int
}

func main() {

	var costPrice, amountOfMoney int
	var step bool
	cashBox := []Cash{
		{
			nominal: 50,
			coint:   2},
		{nominal: 10,
			coint: 3,
		}, {
			nominal: 5,
			coint:   5,
		}, {
			nominal: 2,
			coint:   5,
		}, {
			nominal: 1,
			coint:   5,
		}}
	for !step {
		fmt.Println("Введите стоимость покупки")
		fmt.Scanln(&costPrice)
		if costPrice <= 0 {
			fmt.Println("Вы ошиблись")
		} else {
			step = true
		}
	}

	for amountOfMoney < costPrice {
		fmt.Println("Введите сумму")
		fmt.Scanln(&amountOfMoney)
		if amountOfMoney < costPrice || amountOfMoney < 0 {
			fmt.Println("К сожалению недостаточно средств")
		}
	}
	result := amountOfMoney - costPrice
	fmt.Printf(" Здача к оплате: %d\n ", result)
	rest := make([]Cash, len(cashBox))
	for i := 0; i < len(cashBox); i++ {
		nomi := result / cashBox[i].nominal
		rest[i].nominal = cashBox[i].nominal
		if cashBox[i].coint < nomi {
			nomi = cashBox[i].coint
		}
		result -= cashBox[i].nominal * nomi
		cashBox[i].coint -= nomi

		rest[i].coint = nomi

		fmt.Printf(" Выдать из кассы банкноту нoминалом %d в количестве %d\n", rest[i].nominal, rest[i].coint)
	}
	if result != 0 {
		fmt.Println("К сожелению в кассе нехватает денег")
		fmt.Println("остаток ", result)
	} else {
		fmt.Println("здача выдана в полном обьеме")
	}
}
