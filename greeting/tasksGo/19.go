/*Написать программу ,
 которая по дате рождения (день d месяц n) определяет знак Зодиака:
22 ноября по 21 декабря - Стрелец (12);
22 декабря по 21 января - Козерог (1);
22 января по 21 февраля - Водолей (2);
22 февраля по 21 марта - Рыбы (3).
22 марта по 21 апреля - Овен (4);
22 апреля по 21 мая - Телец (5);
22 мая по 21 июня - Близнецы (6);
22 июня по 21 июля -Рак (7);
22 июля по 21 августа -Лев (8);
22 августа по 21 сентября -Дева (9);
22 сентября по 21 октября - Весы (10);
22 октября по 21 ноября-	Скорпион (11);
*/

package main

import "fmt"

type Data struct {
	day   int
	month int
}

func main() {

	data := Data{23, 10}

	switch {
	case data.month == 1 && data.day >= 22 || data.month == 2 && data.day <= 21:
		fmt.Printf("Дата соответствует знаку Зодиака Водолей ")
	case data.month == 2 && data.day >= 22 || data.month == 3 && data.day <= 21:
		fmt.Printf("Дата соответствует знаку Зодиака Рыбы ")
	case data.month == 3 && data.day >= 22 || data.month == 4 && data.day <= 21:
		fmt.Printf("Дата соответствует знаку Зодиака Овен ")
	case data.month == 4 && data.day >= 22 || data.month == 5 && data.day <= 21:
		fmt.Printf("Дата соответствует знаку Зодиака Телец ")
	case data.month == 5 && data.day >= 22 || data.month == 6 && data.day <= 21:
		fmt.Printf("Дата соответствует знаку Зодиака Близнецы ")
	case data.month == 6 && data.day >= 22 || data.month == 7 && data.day <= 21:
		fmt.Printf("Дата соответствует знаку Зодиака Рак ")
	case data.month == 7 && data.day >= 22 || data.month == 8 && data.day <= 21:
		fmt.Printf("Дата соответствует знаку Зодиака Лев ")
	case data.month == 8 && data.day >= 22 || data.month == 9 && data.day <= 21:
		fmt.Printf("Дата соответствует знаку Зодиака Дева ")
	case data.month == 9 && data.day >= 22 || data.month == 10 && data.day <= 21:
		fmt.Printf("Дата соответствует знаку Зодиака Весы ")
	case data.month == 10 && data.day >= 22 || data.month == 11 && data.day <= 21:
		fmt.Printf("Дата соответствует знаку Зодиака Скорпион ")
	case data.month == 11 && data.day >= 22 || data.month == 12 && data.day <= 21:
		fmt.Printf("Дата соответствует знаку Зодиака Стрелец ")
	case data.month == 12 && data.day >= 22 || data.month == 1 && data.day <= 21:
		fmt.Printf("Дата соответствует знаку Зодиака Козерог ")
	}

}
