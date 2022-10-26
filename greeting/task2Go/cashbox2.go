// Описание задачи
// Имитируется работа кассира. С клавиатуры вводятся
// стоимость покупки и сумма денег, полученных от
// покупателя. Кассир рассчитывает сумму сдачи, которую
// он должен вернуть покупателю. Сумма сдачи должна
// быть сформирована из минимального количества
// банкнот.

package main

import "fmt"

// данные кассы одного номинала
type banknota struct {
	Value int // номинал
	Count int // количество в кассе
}

func main() {
	// ввод цены не буду писать
	// просто возьму тестовые данные
	// сумма к оплате
	totalSumma := 1000
	// сумма от покупателя
	fromBuyerSumma := 2000

	// формируем кассу

	cassa := []banknota{

		{
			Value: 500, Count: 1,
		},
		{
			Value: 200,
			Count: 1,
		}, {
			Value: 100,
			Count: 0,
		}, {
			Value: 50,
			Count: 1,
		}, {
			Value: 20,
			Count: 1,
		}, {
			Value: 10,
			Count: 20,
		},
	}

	forRest := fromBuyerSumma - totalSumma
	fmt.Println("Надо дать сдачи ", forRest)

	if forRest < 0 {
		fmt.Println("Не все деньги дал покупатель, еще должен дать ", totalSumma-fromBuyerSumma, ". Программа закрывается.")
		return
	}

	if forRest == 0 {
		fmt.Println("Покупатель дал точную сумму. Программа закрывается.")
	}

	var (
		rest      []banknota
		restSumma int
	)

	tempTotalRest := forRest

	for i := 0; i < len(cassa) && tempTotalRest > 0; i++ {

		n := tempTotalRest / cassa[i].Value

		fmt.Println("банкнот ", cassa[i].Value, "надо дать ", n)
		fmt.Println("n:=", n)
		if cassa[i].Count < n {
			n = cassa[i].Count
			fmt.Println("В кассе банкнот ", cassa[i].Value, "не хватает, поэтому даем ", n)

		}

		rest = append(rest, banknota{
			Value: cassa[i].Value,
			Count: n,
		})
		tempTotalRest -= n * cassa[i].Value
		restSumma += n * cassa[i].Value

		fmt.Println("Осталось из сдачь собрать еще", tempTotalRest)
	}

	if restSumma < forRest {
		fmt.Println("Денег в кассе не хватило. Eсть только сумма: ", restSumma)
	}

	fmt.Printf("Сдачи собраны %d. Банкноты для выдачи: %#v\n", restSumma, rest)
}
