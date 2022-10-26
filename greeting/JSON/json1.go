package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type muStruct struct {
	Items []struct {
		Ratings string `json:"Ratings"`
	}
}

func main() {

	body, err := ioutil.ReadFile("json.txt")
	if err != nil {
		log.Fatal((err))
	}
	//fmt.Println(string(file))
	items := muStruct{}
	if err := json.Unmarshal(body, &items); err != nil {
		log.Fatal(err)
	}
	fmt.Println(items)
}
