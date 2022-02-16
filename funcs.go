package main

import (
	"fmt"
	"math"
	"strconv"
)

func round(num float64, to int) float64 {
	return math.Round(num*math.Pow10(to)) / math.Pow10(to)
}

func getInp(msg string) uint64 {
	for {
		var inp string

		fmt.Print(msg)
		fmt.Scanln(&inp)

		num, err := strconv.ParseUint(inp, 10, 64)

		if err != nil {
			fmt.Println("Please only enter positive integers")
			continue
		}

		return num
	}
}
