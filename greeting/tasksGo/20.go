/*Почтовый автомат предлагает поздравительные открытки на три темы
1- Новогодние,
2 - С Днем Рождения,
3 - С Днем Защитника Отечества
в трех вариантах (a, b, c) по цене 2 гривны.
 Ввести с клавиатуры номер темы ,
вариант, и купюру оплаты (5,10, 20 гривен).
Выдать нужную открытку (сообщение, например: «Новогодние, вариант с»,
 а также сдачу (купюрами 1,2,5,10) с виде сообщения, например, «2гр+1гр».
  Предусмотреть обработку неправильного номера или варианта.*/
package main

import (
	"fmt"
	"os"
)

func main() {

	postcards := []string{"Новогодние", "С Днем Рождения", "С Днем Защитника Отечества"}
	var theme, banknote int
	var themeOption string
	fmt.Println("Enter the number of theme")
	fmt.Fscan(os.Stdin, &theme)
	fmt.Println("Enter a theme option")
	fmt.Fscan(os.Stdin, &themeOption)
	fmt.Println("Enter the denomination of the banknote")
	fmt.Fscan(os.Stdin, &banknote)
	fmt.Printf("You have chosen a postcard %d with an option %s and paid %d", theme, themeOption, banknote)
	fmt.Println(postcards)

}
