package main

import "fmt"

var (
	firstOptions = []string{"Quit", "Show continents", "Show countries"}
	countryList  = []string{"Sort by name", "Sort by HDI", "Show HDI statistics", "Find a country"}
)

func main() {
programLoop:
	for {
		fmt.Println("Options:")

		for i, j := range firstOptions {
			fmt.Printf("%v.) %v\n", i, j)
		}

		choice := getInp("Choice: ")

		switch choice {
		case 0:
			break programLoop

		case 1:
			fmt.Println("Continents")

			continents := GetContinents()
			list := continents.Sort()

			fmt.Println("Global: ")
			fmt.Printf("Countries: %v\n", list[2].Countries)
			fmt.Printf("Average HDI: %v\n\n", list[2].HdiAverage)

			for i := len(list) - 1; i > -1; i-- {
				if i != 0 {
					fmt.Println(list[i].Continent)
					fmt.Printf("Countries: %v\n", list[i].Countries)
					fmt.Printf("Average HDI: %v\n\n", list[i].HdiAverage)
				}
			}

		case 2:
			fmt.Println("\nCountry options:")

			for i, j := range countryList {
				fmt.Printf("%v.) %v\n", i+1, j)
			}

			choice := getInp("Choice: ")
			var countries Countries = GetCountries()

			switch choice {
			case 1:
				list := countries.SortByCountry()
				fmt.Println("Sorting by country name")
				fmt.Println()

				for _, i := range list {
					fmt.Printf("Country: %v\n", i.Country)
					fmt.Printf("Continent: %v\n", i.Continent)
					fmt.Printf("HDI: %v\n\n", i.Hdi)
				}

			case 2:
				fmt.Println("Sorting by HDI")
				fmt.Println()

				list := countries.SortByHdi()

				for i := len(list) - 1; i > -1; i-- {
					fmt.Printf("Country: %v\n", list[i].Country)
					fmt.Printf("Continent: %v\n", list[i].Continent)
					fmt.Printf("HDI: %v\n\n", list[i].Hdi)
				}

			case 3:
				fmt.Println("Showing statistics")
				fmt.Println()

				list := countries.SortByHdi()
				quartiles := GetQuartiles(list)

				fmt.Printf("First quartile: %v\n", round(quartiles.first, 3))
				fmt.Printf("Second quartile: %v\n", round(quartiles.second, 3))
				fmt.Printf("Third quartile: %v\n\n", round(quartiles.third, 3))

				first, second, third := GetRealQuartiles(list)

				fmt.Println("Real Quartiles")
				fmt.Printf("First quartile: %v (%v)\n", list[first].Country, list[first].Hdi)
				fmt.Printf("Second quartile: %v (%v)\n", list[second].Country, list[second].Hdi)
				fmt.Printf("Third quartile: %v (%v)\n\n", list[third].Country, list[third].Hdi)

			case 4:
				var inp string
				list := countries.SortByCountry()
				list2 := countries.SortByHdi()

				fmt.Print("\nFind a country: ")
				fmt.Scanln(&inp)

				info := GetCountryInfo(list, list2, inp)

				fmt.Printf("Country: %v\n", info.Country)
				fmt.Printf("Continent: %v\n", info.Continent)
				fmt.Printf("HDI: %v\n", info.Hdi)

			default:
				fmt.Println("Invalid input")
			}

		default:
			fmt.Println("Invalid input")
		}
	}
}
