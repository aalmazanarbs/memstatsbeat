package beater

import "math"

func prettyBytesNumber(val uint64) float64 {
	return toFixed(bytesToGB(val), 2)
}

func prettyKilobytesNumber(val int32) float64 {
	return toFixed(kBToGB(val), 6)
}

func kBToGB(val int32) float64 {
	return float64(val) / 1024 / 1024
}

func bytesToGB(val uint64) float64 {
	return float64(val) / 1024 / 1024 / 1024
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}