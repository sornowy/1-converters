package main

import "fmt"

func main() {
	//Конвертация из USD в EUR
	const convertEUR = 0.97
	const convertRUB = 104.55
	const convertEURtoRUB = convertRUB / convertEUR
	fmt.Print(convertEURtoRUB)
}
