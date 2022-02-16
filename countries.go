package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func GetCountries() Countries {
	data, err := os.Open("data.json")

	if err != nil {
		panic("Cannot open json")
	}

	defer data.Close()

	byteValue, _ := ioutil.ReadAll(data)
	var countries Countries

	json.Unmarshal(byteValue, &countries)

	return countries
}

func (count Countries) SortByCountry() []Country {
	arr := []string{}
	cnt := []Country{}

	for _, i := range count.Country {
		arr = append(arr, i.Country)
	}

	sort.Strings(arr)

	for _, i := range arr {
		for j, k := range count.Country {
			if i == k.Country {
				cnt = append(cnt, count.Country[j])
			}
		}
	}

	return cnt
}

func (count Countries) SortByHdi() []Country {
	arr := []float64{}
	cnt := []Country{}

	for _, i := range count.Country {
		arr = append(arr, i.Hdi)
	}

	sort.Float64s(arr)

	for _, i := range arr {
		for j, k := range count.Country {
			if i == k.Hdi {
				cnt = append(cnt, count.Country[j])
			}
		}
	}

	return cnt
}

func GetCountryInfo(lstByAlp []Country, lstByHdi []Country, country string) Country {
	arr := []string{}

	for _, i := range lstByAlp {
		arr = append(arr, i.Country)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("No country named %v was found\n\n", country)
		}
	}()

	i := sort.SearchStrings(arr, country)
	return lstByAlp[i]
}
