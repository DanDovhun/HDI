package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var global = Continent{
	Continent:  "Global",
	Countries:  0,
	HdiAverage: 0,
}

var africa = Continent{
	Continent:  "Africa",
	Countries:  0,
	HdiAverage: 0,
}

var asia = Continent{
	Continent:  "Asia",
	Countries:  0,
	HdiAverage: 0,
}

var centralAmerica = Continent{
	Continent:  "Central America",
	Countries:  0,
	HdiAverage: 0,
}

var europe = Continent{
	Continent:  "Europe",
	Countries:  0,
	HdiAverage: 0,
}

var northAmerica = Continent{
	Continent:  "North America",
	Countries:  0,
	HdiAverage: 0,
}

var oceania = Continent{
	Continent:  "Oceania",
	Countries:  0,
	HdiAverage: 0,
}

var southAmerica = Continent{
	Continent:  "South America",
	Countries:  0,
	HdiAverage: 0,
}

func CreateData() {
	file, err := ioutil.ReadFile("hdi.csv")

	var countryList []Country
	var list JSON

	if err != nil {
		panic("Cannot open the csv file")
	}

	content := strings.Split(string(file), "\n")

	for _, i := range content {
		j := strings.Split(i, ",")

		hdi, _ := strconv.ParseFloat(strings.ReplaceAll(j[2], "\r", ""), 64)

		countryList = append(countryList, Country{
			Country:   j[0],
			Continent: j[1],
			Hdi:       hdi,
		})

		global.Countries++
		global.HdiAverage += hdi

		switch j[1] {
		case "Africa":
			africa.Countries++
			africa.HdiAverage += hdi

		case "Asia":
			asia.Countries++
			asia.HdiAverage += hdi

		case "Central America":
			centralAmerica.Countries++
			centralAmerica.HdiAverage += hdi

		case "Europe":
			europe.Countries++
			europe.HdiAverage += hdi

		case "North America":
			northAmerica.Countries++
			northAmerica.HdiAverage += hdi

		case "Oceania":
			oceania.Countries++
			oceania.HdiAverage += hdi

		case "South America":
			southAmerica.Countries++
			southAmerica.HdiAverage += hdi
		}
	}

	global.HdiAverage /= float64(global.Countries)
	global.HdiAverage = round(global.HdiAverage, 3)

	africa.HdiAverage /= float64(africa.Countries)
	africa.HdiAverage = round(africa.HdiAverage, 3)

	asia.HdiAverage /= float64(asia.Countries)
	asia.HdiAverage = round(asia.HdiAverage, 3)

	centralAmerica.HdiAverage /= float64(centralAmerica.Countries)
	centralAmerica.HdiAverage = round(centralAmerica.HdiAverage, 3)

	europe.HdiAverage /= float64(europe.Countries)
	europe.HdiAverage = round(europe.HdiAverage, 3)

	northAmerica.HdiAverage /= float64(northAmerica.Countries)
	northAmerica.HdiAverage = round(northAmerica.HdiAverage, 3)

	oceania.HdiAverage /= float64(oceania.Countries)
	oceania.HdiAverage = round(oceania.HdiAverage, 3)

	southAmerica.HdiAverage /= float64(southAmerica.Countries)
	southAmerica.HdiAverage = round(southAmerica.HdiAverage, 3)

	var continents = []Continent{global, africa, asia, centralAmerica, europe, northAmerica, oceania, southAmerica}
	list = JSON{
		Continent: continents,
		Country:   countryList,
	}

	fmt.Println(list.Country)

	AddToJSON(list)
}

func AddToJSON(countries JSON) {
	data, err := os.Open("data.json")
	var list JSON

	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, &list)

	list.Country = append(list.Country, countries.Country...)
	list.Continent = append(list.Continent, countries.Continent...)

	j, _ := json.MarshalIndent(list, "", "    ")
	_ = ioutil.WriteFile("data.json", j, 0644)
}
