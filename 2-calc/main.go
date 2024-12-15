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
	operations := map[string] func([]int) {
		"AVG": getAvg,
		"SUM": getSum,
		"MED": getMed,
	}
	fmt.Println("Добро пожаловать в калькулятор!")
	fmt.Println("\nВыберите операцию")
	userChoise := getChoiseOperation()
	inputNumbers := getNumbers()
	operations[userChoise](inputNumbers)
}

func getChoiseOperation() string {
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
	var elemStr string
	var numbers []int 
	for {
		fmt.Println("Введите числа через запятую(считаются целые числа, отличные от 0):")
		in := bufio.NewReader(os.Stdin)
		userInput, err := in.ReadString('\n') // Читаем всю строку
		if err != nil {
			fmt.Println("Ошибка чтения, попробуйте снова")
			return nil
		}
		userInput = strings.ReplaceAll(userInput, " ", "") // убираем пробелы
		userInput = strings.TrimSpace(userInput) // убираем лишнее в начале и конце
		var inputGarbage []string
		for len(userInput) > 0 {
			elemStr, userInput, _ = strings.Cut(userInput, ",")
			elemInt, err := strconv.Atoi(elemStr) // преобразуем в int
			if err != nil {
				inputGarbage = append(inputGarbage, elemStr) // если это не число - то добавляем в "мусорку", чтоб потом сказать что это не принято
				continue
			}
			numbers = append(numbers, elemInt) // добавляем отрезанный и преобразованный элемент в слайс
		}
		if len(inputGarbage) != 0 {
			fmt.Printf("\nЧасть элементов удалена, так как это не числа: %v\n", inputGarbage)
		}
		if len(numbers) == 0 {
			fmt.Println("Ошибка ввода, список числе пустой, попробуйте снова")//если буквы или пустой ввод или нули - повторная попытка ввода
			continue
		}
		return numbers
	}
}

func getAvg (numbers []int) {
	var sum = 0
	for _, elem := range numbers { // высчитываем сумму элементов
		sum += elem
	}
	var result float64 = float64(sum) / float64(len(numbers)) // высчитываем среднее арифметическое
	fmt.Printf("\n\nСреднее арифметическое введённых чисел - %.2f \n\n", result)
}

func getSum (numbers []int ) {
	var sum = 0
	for _, elem := range numbers {
		sum += elem
	}
	fmt.Printf("\n\nСумма введённых чисел - %v \n\n", sum)
}

func getMed (numbers []int) {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	var med float64
	if len(numbers) % 2 == 0 {
		med = (float64(numbers[((len(numbers) / 2)) - 1]) + float64(numbers[len(numbers) / 2])) / 2
	} else {
		med = float64(numbers[((len(numbers) - 1) / 2)])
	}
	fmt.Printf("\nМедиана введённых чисел - %v\n\n", med)
}
// 	if middleNumber > elem {
// 		return middleNumber - elem
// 	}
// 	return elem - middleNumber
// }