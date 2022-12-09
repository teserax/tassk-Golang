//Структуры данных. Файлы. Дидактический блок 14Основы программирования. Фундаментальный курс | 217
//2. Дана следующая структура:
//No карточки Пациент Заметки
//ФИО Возраст Группа крови Диагноз
//Разработай алгоритм, который:
////2.1. Создает файл с вышеуказанной структурой, информация..............."+"
////о пациенте считывается с клавиатуры.............."+"
//2.2. Выводит содержимое файла на экран................."+"
//2.3. Определяет количество пациентов:,,,,,,,,,,,,,,,"+"
//до 40 лет;........................."+"
//от 41 до 60 лет;..............."+"
//более 60 лет.............."+"
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
	FullName
}
type FullName struct {
	FirstName  string
	LastName   string
	Patronymic string
}
type List []Patient
type Options struct {
	MinAge  int
	MaxAge  int
	Group   string
	Diagnos string
}

func (list List) Filter(option Options) List {
	var result List
	for _, pacient := range list {
		if option.MinAge > pacient.Age {
			result = append(result, pacient)
		}
		if option.Group == pacient.BloodType {
			result = append(result, pacient)
		}
	}
	return result
}

func main() {
	file, err := os.Create("card-index.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	ListOfPacient := List{}
	start := true
	data := []byte{}

	for start {

		hospitalQuestionnaire := Patient{}
		fmt.Println("Enter cardNumber of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.CardNumber)
		fmt.Println("Enter FirstName of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.FirstName)
		fmt.Println("Enter LastName of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.LastName)
		fmt.Println("Enter Patronymic of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.Patronymic)
		fmt.Println("Enter Age of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.Age)
		fmt.Println("Enter blood of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.BloodType)
		fmt.Println("Enter diagnosis of Pacient ")
		fmt.Scan(&hospitalQuestionnaire.Diagnosis)

		ListOfPacient = append(ListOfPacient, hospitalQuestionnaire)
		var answer string

		data, err = json.MarshalIndent(ListOfPacient, " ", "")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Enter New Pacient ? y/n ")
		fmt.Scan(&answer)
		if answer != "y" {
			start = false
		}
	}
	file.Write(data)
	info, err := os.ReadFile("card-index.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	fmt.Println(string(info))
	infoResult := ListOfPacient.Filter(Options{MinAge: 40}).Filter(Options{Group: "AO"})

	fmt.Println(infoResult)
}
