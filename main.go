package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Калькулятор конвертаций")
	for {
		var currency string
		var err error

		for {
			currency, err = getUserCurrency()
			if err != nil {
				fmt.Println("Вы ввели неверно валюту. Введите валюту заново")
			} else {
				break
			}
		}

		fmt.Println(" Выбранная валюта:", currency)

		var value float64
		for {
			value, err = getUserValue()
			if err != nil {
				fmt.Println("Вы ввели неверно суммую. Введите сумму заново")
			} else {
				break
			}
		}

		fmt.Println(" Выбранная сумма:", value)

		var targetCurrency string
		for {
			targetCurrency, err = getUserTargetCurrency(currency)
			if err != nil {
				fmt.Println("Вы ввели неверно целевую валюту. Введите валюту заново")
			} else {
				break
			}
		}
		fmt.Println(" Выбранная целевая валюта: ", targetCurrency)

		conversion := calculateCurrency(currency, value, targetCurrency)
		result := fmt.Sprintf("Сумма: %.02f", conversion)
		fmt.Println(result)
		isRepeatConv := repeatConversion()
		if !isRepeatConv {
			break
		}
	}

}

func getUserValue() (float64, error) {
	var value float64
	fmt.Print("Введите сумму: ")
	fmt.Scan(&value)
	if value <= 0 {
		return 0, errors.New("Invalid_Currency")
	}
	return value, nil
}

func getUserCurrency() (string, error) {
	var originCurrency string
	fmt.Print("Введите исходную валюту (USD/EUR/RUB): ")
	fmt.Scan(&originCurrency)
	if originCurrency != "USD" && originCurrency != "EUR" && originCurrency != "RUB" {
		return "", errors.New("Invalid_Currency")
	}
	return originCurrency, nil
}

func getUserTargetCurrency(currency string) (string, error) {
	var targetCurrency string
	fmt.Print("Введите целевую валюту: ")
	fmt.Scan(&targetCurrency)
	if currency == targetCurrency {
		return "", errors.New("Invalid_Target_Currency")
	}
	if targetCurrency != "USD" && targetCurrency != "EUR" && targetCurrency != "RUB" {
		return "", errors.New("Invalid_Target_Currency")
	}
	return targetCurrency, nil
}

func calculateCurrency(currency string, value float64, targetCurrency string) float64 {
	const convertUSDtoEUR = 0.97
	const convertEURtoUSD = 1.03
	const convertUSDtoRUB = 104.55
	const convertEURtoRUB = convertUSDtoRUB / convertUSDtoEUR
	var currencyNow float64

	if currency == "USD" && targetCurrency == "EUR" {
		currencyNow = convertUSDtoEUR
	} else if currency == "EUR" && targetCurrency == "USD" {
		currencyNow = convertEURtoUSD
	} else if currency == "RUB" && targetCurrency == "USD" {
		currencyNow = 1 / convertUSDtoRUB
	} else if currency == "RUB" && targetCurrency == "EUR" {
		currencyNow = convertEURtoRUB / convertUSDtoRUB
	} else if currency == "EUR" && targetCurrency == "RUB" {
		currencyNow = convertEURtoRUB
	} else if currency == "USD" && targetCurrency == "RUB" {
		currencyNow = convertUSDtoRUB
	} else {
		fmt.Println("Неверная валюта или цель конверсии")
	}

	conversion := value * currencyNow
	return conversion
}

func repeatConversion() bool {
	var reapeat string
	fmt.Print("Если хотите продолжить введите (Y/N): ")
	fmt.Scan(&reapeat)
	if reapeat == "Y" {
		return true
	}
	return false
}
