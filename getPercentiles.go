package main

import statistics "github.com/DanDovhun/Statistics"

func getPercentile(arr []float64, per float64) float64 {
	percentile := statistics.Percentile(arr, 25)

	return percentile
}
