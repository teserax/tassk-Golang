package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Countries struct {
	Countri []string
}
type City struct {
	City string
}

func main() {
	url := "https://raw.githubusercontent.com/russ666/all-countries-and-cities-json/master/countries.json"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	file := map[string][]string{}

	err = json.Unmarshal(data, &file)
	if err != nil {
		log.Fatal(err)
	}

	for countru, city := range file {
		fmt.Printf("Country %s\n city %s\n", countru, city)

	}
}
