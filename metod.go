package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Persona struct {
	Id      string
	FIO     string
	Address string
	Product Product
}

type Product struct {
	Goods    []string
	Quantity int
	Price    string
}

func main() {
	file, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var persons []Persona
	for scanner.Scan() {
		//   fmt.Println(scanner.Text())
		var person Persona
		err = json.Unmarshal(scanner.Bytes(), &person)
		persons = append(persons, person)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(persons)
}
