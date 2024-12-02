package main

import (
	"fmt"
	"strings"
)

func main() {
	conversionValues := map[string]map[string]float64 {
		"USD": {"EUR":0.96, "RUB":104.35},
		"EUR": {"USD":1.04, "RUB":108.77},
		"RUB": {"USD":0.0096, "EUR":0.0092},
	}
	fmt.Println("Добро пожаловать в конвертер валют!")
	originalCurrency := getInput("Введите начальную валюту", "")
	value := getInputValue()
	resultCurrency := getInput("Введите итоговую валюту", originalCurrency)
	value = conversion(originalCurrency, value, resultCurrency, &conversionValues)
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
		fmt.Printf(output + " " + outputCurrency +": ")
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
		fmt.Printf("Введите количество валюты: ")
		fmt.Scan(&value)
		if value <= 0 {
			fmt.Println("Ошибка ввода, попробуйте снова")
			continue
		}
		return value
	}
}
func conversion (originalCurrency string, value float64, resultCurrency string, conversionValues *map[string]map[string]float64) float64 {
	value *= (*conversionValues)[originalCurrency][resultCurrency]
	return value
}