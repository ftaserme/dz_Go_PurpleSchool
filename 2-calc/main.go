package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)


func main() {
	operations := map[string] func([]int) float64 {
		"AVG": getAvg,
		"SUM": getSum,
		"MED": getMed,
	}
	fmt.Println("Добро пожаловать в калькулятор!")
	fmt.Println("\nВыберите операцию")
	menuChoise := getMenuChoise()
	numbers := getNumbers()
	operations[menuChoise](numbers)
}

func getMenuChoise() string {
	fmt.Println("AVG - среднее арифметическое")
	fmt.Println("SUM - сумма")
	fmt.Println("MED - медиана")
	var userInput string
	for {
		fmt.Scanln(&userInput)
		userInput = strings.ToUpper(userInput) // возводим всё в заглавные символы для единообразия и простоты отсеивания
		if userInput != "AVG" && userInput != "SUM" && userInput != "MED" {
			fmt.Print("\nОшибочный ввод, попробуйте снова: ")	
		} else {
			return userInput
		}
	}
}

func getNumbers() []int  {
	var onePosition string
	var numbers []int 
	for {
		fmt.Println("Введите числа через запятую(считаются целые числа, отличные от 0):")
		input := bufio.NewReader(os.Stdin)
		inputStr, err := input.ReadString('\n') // Читаем всю строку
		if err != nil {
			fmt.Println("Ошибка чтения, попробуйте снова")
			return nil
		}
		inputStr = strings.ReplaceAll(inputStr, " ", "") // убираем пробелы
		inputStr = strings.TrimSpace(inputStr) // убираем лишнее в начале и конце
		var inputGarbage []string
		for len(inputStr) > 0 {
			onePosition, inputStr, _ = strings.Cut(inputStr, ",")
			number, err := strconv.Atoi(onePosition) // преобразуем в int
			if err != nil {
				inputGarbage = append(inputGarbage, onePosition) // если это не число - то добавляем в "мусорку", чтоб потом сказать что это не принято
				continue
			}
			numbers = append(numbers, number) // добавляем отрезанный и преобразованный элемент в слайс
		}
		if len(inputGarbage) != 0 {
			fmt.Printf("\nЧасть элементов удалена, так как это не числа: %v\n", inputGarbage)
		}
		if len(numbers) == 0 {
			fmt.Println("Ошибка ввода, список чисел пустой, попробуйте снова")//если буквы или пустой ввод или нули - повторная попытка ввода
			continue //проверку на пустой массив выполняю тут, при вводе
		}
		return numbers
	}
}

func getAvg (numbers []int) float64 {
	sum := getSum(numbers)
	var result = sum / float64(len(numbers)) // высчитываем среднее арифметическое
	fmt.Printf("\n\nСреднее арифметическое введённых чисел - %.2f \n\n", result)
	return result
}

func getSum (numbers []int ) float64 {
	var sum = 0
	for _, elem := range numbers {
		sum += elem
	}
	fmt.Printf("\n\nСумма введённых чисел - %v \n\n", sum)
	return float64(sum)
}

func getMed (numbers []int) float64 {
	sort.Ints(numbers)
	var med float64
	if len(numbers) % 2 == 0 {
		med = (float64(numbers[((len(numbers) / 2)) - 1]) + float64(numbers[len(numbers) / 2])) / 2
	} else {
		med = float64(numbers[((len(numbers) - 1) / 2)])
	}
	fmt.Printf("\nМедиана введённых чисел - %.2f\n\n", med)
	return med
}