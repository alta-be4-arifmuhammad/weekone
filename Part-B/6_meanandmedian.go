package main

import "fmt"

func MeanMedian(arrayInput []float64) (float64, float64) {

	sum := 0.0
	median := 0.0

	for i := 0; i < len(arrayInput); i++ {
		sum = sum + float64(arrayInput[i])
	}
	mean := sum / float64(len(arrayInput))

	if len(arrayInput)%2 == 0 {
		indexMedian := (len(arrayInput) / 2)
		median = (float64(arrayInput[indexMedian]) + float64(arrayInput[indexMedian-1])) / 2
	} else {
		indexMedian := len(arrayInput) / 2
		median = float64(arrayInput[indexMedian])
	}
	return mean, median
}

func main() {
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4}))
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4, 5}))
	fmt.Println(MeanMedian([]float64{7, 8, 9, 13, 15}))
	fmt.Println(MeanMedian([]float64{10, 20, 30, 40, 50}))
	fmt.Println(MeanMedian([]float64{15, 20, 30, 60, 120}))
}
