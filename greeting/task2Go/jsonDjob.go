package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Persona struct {
	ID       int      `json:"id,omitempty"` //если поле содержит нулевое значение не записываем в JSON
	Number   string   `json:"number"`       //прописываем теги для лучьшей десеарилизиции
	Year     string   `json:"year"`
	Students []string `json:"students"`
	Password string   `json:"-"` // если мы хотим что бы поле неотображалось
}

func main() {

	data := Persona{
		ID:       0,
		Number:   "55",
		Year:     "1988",
		Students: []string{"Vasia", "Petia"},
		Password: "superPassword",
	}
	file, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Marshal in JSON %v\n", string(file))

	infoData := Persona{}
	if err = json.Unmarshal(file, &infoData); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Unmarshal JSON %v\n", infoData)
	out, err := json.MarshalIndent(infoData, "", "  ") // MarshalIndent выводит в Json формате
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
