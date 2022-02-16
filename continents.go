package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
)

func GetContinents() Continents {
	data, err := os.Open("data.json")

	if err != nil {
		panic("Cannot open json")
	}

	defer data.Close()

	byteValue, _ := ioutil.ReadAll(data)
	var continents Continents

	json.Unmarshal(byteValue, &continents)

	return continents
}

func (cont Continents) Sort() []Continent {
	arr := []float64{}
	cnt := []Continent{}

	for _, i := range cont.Continent {
		arr = append(arr, i.HdiAverage)
	}

	sort.Float64s(arr)

	for _, i := range arr {
		for j, k := range cont.Continent {
			if i == k.HdiAverage {
				cnt = append(cnt, cont.Continent[j])
			}
		}
	}

	return cnt
}
