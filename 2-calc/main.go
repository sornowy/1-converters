package main

import (
	"fmt"
	"sort"
)

func main() {
	summary := []float64{}
	for {
		number := scanNumber()
		if number == 0 {
			break
		}
		summary = append(summary, number)
	}

	fmt.Printf("Введенные числа: %s\n", formatArray(summary))

	choose := outputFunc()
	if choose == "AVG" {
		fmt.Println(calcAvg(summary))
	} else if choose == "SUM" {
		fmt.Println(calcSum(summary))
	} else if choose == "MED" {
		fmt.Println(calcMed(summary))
	}

}

func scanNumber() float64 {
	var number float64
	fmt.Print("Введите число(n для выхода): ")
	fmt.Scan(&number)
	return number
}

func calcAvg(summary []float64) float64 {
	var average, sum float64
	for _, value := range summary {
		sum += value
	}
	average = sum / float64(len(summary))
	return average
}

func calcSum(summary []float64) float64 {
	var sum float64
	for _, value := range summary {
		sum += value
	}
	return sum
}

func calcMed(summary []float64) float64 {
	sort.Float64s(summary)
	med := len(summary)

	if med%2 == 1 {
		return summary[med/2]
	} else {
		return (summary[med/2-1] + summary[med/2]) / 2
	}
}

func outputFunc() string {
	var a string
	fmt.Print("Введите какой нужен калькулятор(AVG/SUM/MED): ")
	fmt.Scan(&a)
	return a
}

func formatArray(summary []float64) string {
	if len(summary) == 0 {
		return "[]"
	}

	result := "["
	for index, value := range summary {
		result += fmt.Sprintf("%.2f", value)
		if index != len(summary)-1 {
			result += ", "
		}
	}
	result += "]"
	return result
}
