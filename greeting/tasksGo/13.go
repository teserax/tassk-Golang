package main

import (
	"C/Users/go/src/greeting/tasksGo/calendar"
	"fmt"
	"log"
)

type Date string

func (d Date) PrintDate() {
	fmt.Println(d)
}
func main() {
	persona := calendar.Persona{}
	err := persona.SetName("Valerius")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(persona.Name())
	persona1 := calendar.Persona{}
	err = persona1.SetName("Vova")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(persona1.Name())
	data := Date("25")
	data.PrintDate()
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

}
