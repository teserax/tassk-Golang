//Структуры данных. Файлы. Дидактический блок 14Основы программирования. Фундаментальный курс | 217
//2. Дана следующая структура:
//No карточки Пациент Заметки
//ФИО Возраст Группа крови Диагноз
//Разработай алгоритм, который:
////2.1. Создает файл с вышеуказанной структурой, информация..............."+"
////о пациенте считывается с клавиатуры.............."+"
//2.2. Выводит содержимое файла на экран................."+"
//2.3. Определяет количество пациентов:
//до 40 лет;
//от 41 до 60 лет;
//более 60 лет.
//2.4. Определяет количество пациентов с гипертонической
//болезнью и возрастом менее 50 лет.
//2.5. Определяет количество пациентов с артритом и группой
//крови A4.
//2.6. Определяет какая группа крови является наиболее рас-
//пространенной при анемии – A2 или A0?
//2.7. Определяет наиболее распространенную группу крови.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Patient struct {
	CardNumber int
	Age        int
	BloodType  string
	Diagnosis  string
	FullName   `json:"FullName"`
}
type FullName struct {
	FirstName  string
	LastName   string
	Patronymic string
}

func main() {
	file, err := os.Create("card-index.txt")//создаем фаил
	if err != nil {                         //проверяем на ошибки
		log.Fatal("Error make file")
		os.Exit(1)
	}
	defer file.Close()//закрываем фаил

	infoOfPacient := map[int]Patient{}//храниь инфу будем в мапе структур
	start := true                 
	inc := 0
	for start {                 //пременная ..старт.. если верно то продолжаем цикл

		var cardNumber, Age int         //переменные запроса на заполнение инфо пациента
		var blood, diagnosis string      //переменные запроса на заполнение инфо пациента
		var FirstName, LastName, Patronymic string    //переменные запроса на заполнение инфо пациента

		fmt.Println("Enter cardNumber of Pacient ")//  запрос на заполнение 
		fmt.Scan(&cardNumber)
		fmt.Println("Enter FirstName of Pacient ")
		fmt.Scan(&FirstName)
		fmt.Println("Enter LastName of Pacient ")
		fmt.Scan(&LastName)
		fmt.Println("Enter Patronymic of Pacient ")
		fmt.Scan(&Patronymic)
		fmt.Println("Enter Age of Pacient ")
		fmt.Scan(&Age)
		fmt.Println("Enter blood of Pacient ")
		fmt.Scan(&blood)
		fmt.Println("Enter diagnosis of Pacient ")
		fmt.Scan(&diagnosis)

		inc++                 //инкрементируем тем самым задаем колличество записанных анкет
		infoOfPacient[inc] = Patient{
			CardNumber: cardNumber,
			Age:        Age,
			BloodType:  blood,
			Diagnosis:  diagnosis,
			FullName: FullName{
				FirstName:  FirstName,
				LastName:   LastName,
				Patronymic: Patronymic,
			},
		}

		data := []byte{}
		data, _ = json.MarshalIndent(infoOfPacient[inc], " ", "")// орабатываем данные ввиде JSON фаила..функцией MarshalIndent для лучшей читабельности
		file.Write(data)//запись в фаил
		var answer string
		fmt.Println("Enter New Pacient ? y/n ")//запрос на добавление новой анкеты
		fmt.Scan(&answer)
		if answer != "y" {
			start = false
		}
	}
	info, err := os.ReadFile("card-index.txt")//читаем файл
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(info)) выводим полученную инфу 

}            
