/*13.3.	 Даны названия 26 городов и стран, в которых они нахо-
дятся. Среди них есть города, находящиеся в Италии. Напечатать
их названия.*/
package main

import "fmt"

func main() {

	countries := map[string][]string{
		"Albania": {
			"Elbasan",
			"Petran",
			"Pogradec",
			"Shkoder",
			"Tirana",
			"Ura Vajgurore",
		},
		"Antigua and Barbuda": {
			"All Saints",
			"Cassada Gardens",
			"Codrington",
			"Old Road",
			"Parham",
			"Woods",
		},
		"Armenia": {
			"Abovyan",
			"Agarak",
			"Apaga",
			"Aparan",
			"Arabkir",
			"Ashtarak",
			"Erebuni Fortress",
			"Hrazdan",
			"Ijevan",
			"Jermuk",
			"Kapan",
			"Tsaghkadzor",
			"Vanadzor",
			"Yerevan",
		},
		"Austria": {
			"Absam",
			"Absdorf",
			"Abtenau",
			"Abtsdorf",
			"Ach",
			"Achenkirch",
			"Achensee",
			"Admont",
			"Adnet",
			"Afritz",
			"Aggsbach",
		},
		"Italya": {
			"Rim",
			"Milan",
			"Neapolii",
			"Turin",
			"Palermo",
			"Boloni",
		},
	}
	for country, city := range countries {
		if country == "Italya" {
			fmt.Printf("%s", city)
		}
	}

}
