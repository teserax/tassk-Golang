package main

import (
	"C/Users/go/src/greeting/tasksGo/calendar"
	"fmt"
	"log"
)

func main() {
	data := calendar.Date{}
	err := data.SetYear(1980)
	if err != nil {
		log.Fatal(err)
	}
	err = data.SetMonth(7)
	if err != nil {
		log.Fatal(err)
	}
	err = data.SetDay(23)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data.Year(), data.Month(), data.Day())
	event := calendar.Event{}
	err = event.SetTitle("Mama Happy")
	if err != nil {
		log.Fatal(err)
	}

	event.SetYear(1950)
	event.SetMonth(5)
	event.SetDay(4)

	fmt.Println(event.Day())
	event.SetName("Alena")
	fmt.Println(event)
}
