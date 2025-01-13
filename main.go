package main

import "fmt"

func main() {
	//Конвертация из USD в EUR
	const convertUSDtoEUR = 0.97
	const convertUSDtoRUB = 104.55
	const convertEURtoRUB = convertUSDtoRUB / convertUSDtoEUR
	fmt.Println(convertEURtoRUB)
	fmt.Print(getUserInput())
}

func getUserInput() float64 {
	var value float64
	fmt.Print("Введите сумму: ")
	fmt.Scan(&value)
	return value
}

func calculateCurrency(convertUSDtoEUR, convertUSDtoRUB, value float64) float64 {
	const convertUSDtoEUR = 0.97
	const convertUSDtoRUB = 104.55
}
