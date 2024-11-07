package main

import (
	"fmt"
	"strings"
)

const usdToEur = 0.93
const usdToRub = 98.29
var eurToRub = usdToEur * usdToRub

func main() {
	fmt.Println("Добро пожаловать в конвертер валют!")
	originalCurrency := getInput("Введите начальную валюту", "")
	value := getInputValue()
	resultCurrency := getInput("Введите итоговую валюту", originalCurrency)
	value = conversion(originalCurrency, value, resultCurrency)
	fmt.Printf("Результат конвертации - %.2f", value)
}

func getInput(output string, usingCurrency string) string {
	var userInput string
	var outputCurrency = "(USD, EUR, RUB)"
	switch usingCurrency {
	case "USD":
		outputCurrency = "(EUR, RUB)"
	case "EUR":
		outputCurrency = "(USD, RUB)"
	case "RUB":
		outputCurrency = "(USD, EUR)"
	}
	fmt.Printf(output + " " + outputCurrency +": ")
	fmt.Scan(&userInput)
	userInput = strings.ToUpper(userInput)
	if userInput == usingCurrency && userInput != "" {
		return getInput("Ошибка, эта валюта уже используется, введите другую", usingCurrency)
	}
	if userInput != "USD" && userInput != "EUR" && userInput != "RUB" {
		return getInput("Ошибка при вводе, попробуйте снова", usingCurrency)
	}
	return userInput
}

func getInputValue () float64 {
	var value float64
	fmt.Printf("Введите количество валюты: ")
	fmt.Scan(&value)
	if value <= 0 {
		fmt.Println("Ошибка ввода, попробуйте снова")
		return getInputValue()
	}
	fmt.Println(value)
	return value
}

func conversion (originalСurrency string, value float64, resultCurrency string) float64 {
	switch {
	case originalСurrency == "USD" :
		if resultCurrency == "RUB" {
			return (value * usdToRub)
		}
		return value * usdToEur
	case originalСurrency == "EUR" :
		if resultCurrency == "USD" {
			return (value / usdToEur)
		}
		return value * eurToRub
	case originalСurrency == "RUB":
		if resultCurrency == "EUR" {
			return (value / eurToRub)
		}
		
	}
	return (value / usdToRub)
}