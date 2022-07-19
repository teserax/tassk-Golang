package main

import "fmt"

type Shop []struct {
	Address  Address
	Employee Employee
	Products
}
type Products struct {
	Ingredients map[string]int
}
type Address struct {
	country string
	city    string
	street  string
	zip     int
}
type Employee struct {
	name   string
	age    int
	salary []int
}

func main() {

	linella := Shop{{Address{
		country: "Moldova",
		city:    "Chisinau",
		street:  "Stefan",
		zip:     3110,
	}, Employee{
		name: "Vova", age: 42, salary: []int{2000, 2500, 3000}},
		Products{
			Ingredients: map[string]int{"мука": 18, "сухари": 33, "марковка": 8}},
	}, {Address{
		country: "Moldova",
		city:    "Chisinau",
		street:  "Costin",
		zip:     3300,
	}, Employee{
		name: "Alena", age: 18, salary: []int{2000, 3500, 4000}},
		Products{
			Ingredients: map[string]int{"мука": 15, "сахар": 33, "пшено": 7}},
	}, {Address{
		country: "Moldova",
		city:    "Chisinau",
		street:  "Pelivan",
		zip:     2150,
	}, Employee{
		name: "Valera", age: 40, salary: []int{1000, 2500, 5000}},
		Products{
			Ingredients: map[string]int{"мука": 19, "мороженое": 25, "гречка": 35}},
	}}

	expensiveFlour := 0
	flourShop := Shop{}
	sugarShop := Shop{}
	for _, shop := range linella {
		if _, exist := shop.Ingredients["сахар"]; exist {
			sugarShop = Shop{shop}
		}
		for product, price := range shop.Ingredients {

			if product == "мука" {

				if expensiveFlour < price {
					expensiveFlour = price
					flourShop = Shop{shop}
				}
			}

		}

	}
	fmt.Printf("Сахар есть в магазине по адресу %v \n", sugarShop[0].Address)
	fmt.Printf("Самая дорогая мука по %d lei в магазине адресу %v\n", expensiveFlour, flourShop[0].Address)

}
