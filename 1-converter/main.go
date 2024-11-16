package main

import (
	"fmt"
	"strings"
)

const usdToEur = 0.93
const usdToRub = 98.29
const eurToRub = usdToEur * usdToRub

func main() {
	fmt.Println("\nДобро пожаловать в конвертер валют!")
	originalCurrency := getInput("\nВведите начальную валюту", "")
	value := getInputValue()
	resultCurrency := getInput("\nВведите итоговую валюту", originalCurrency)
	value = conversion(originalCurrency, value, resultCurrency)
	fmt.Printf("Результат конвертации - %.2f", value)
}

func getInput(output string, usingCurrency string) string {
	var userInput string
	var originalCurrency = "(USD, EUR, RUB)"
	switch usingCurrency {
	case "USD":
		originalCurrency = "(EUR, RUB)"
	case "EUR":
		originalCurrency = "(USD, RUB)"
	case "RUB":
		originalCurrency = "(USD, EUR)"
	}
	for {
		fmt.Printf(output + " " + originalCurrency +": ")
		fmt.Scan(&userInput)
		userInput = strings.ToUpper(userInput)
		if userInput == usingCurrency && userInput != "" {
			fmt.Println("Ошибка, эта валюта уже используется, введите другую")
			continue
		}
		if userInput != "USD" && userInput != "EUR" && userInput != "RUB" {
			fmt.Println("Ошибка при вводе, попробуйте снова")
			continue
		}
		return userInput
	}
}

func getInputValue () float64 {
	var value float64
	for {
		fmt.Printf("\nВведите количество валюты: ")
		fmt.Scan(&value)
		if value <= 0 {
			fmt.Println("Ошибка ввода, попробуйте снова")
			continue
		}
		return value
	}
}

func conversion (originalCurrency string, value float64, resultCurrency string) float64 {
	switch originalCurrency {
	case "USD" :
		return conversionUSD(value, resultCurrency)
	case "EUR" :
		return conversionEUR(value, resultCurrency)
	case "RUB":
		return conversionRUB(value, resultCurrency)
	}
	return 0.0
}

func conversionUSD(value float64, resultCurrency string) float64 {
	if resultCurrency == "RUB" {
		return (value * usdToRub)
	}
	return (value * usdToEur)
}

func conversionEUR(value float64, resultCurrency string) float64 {
	if resultCurrency == "USD" {
		return (value / usdToEur)
	}
	return (value * eurToRub)
}

func conversionRUB(value float64, resultCurrency string) float64 {
	if resultCurrency == "USD" {
		return (value / usdToRub)
	}
	return (value / eurToRub)
}