package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Добро пожаловать в калькулятор!")
	var userChoise string
	fmt.Println("\nВыберите операцию")
	userChoise = getChoiseOperation()
	inputNumbers := getNumbers()
	switch userChoise {
	case "AVG":
		getAvg(&inputNumbers)
	case "SUM":
		getSum(&inputNumbers)
	case "MED":
		getMed(&inputNumbers, inputNumbers[0])
	}
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
	var userInput, elemStr string
	var numbers []int 
	for {
		fmt.Println("Введите числа через запятую(считаются целые числа, отличные от 0):")
		in := bufio.NewReader(os.Stdin)
		userInput, _ = in.ReadString('\n') // Читаем всю строку
		userInput = strings.ReplaceAll(userInput, " ", "") // убираем пробелы
		userInput = strings.TrimSpace(userInput) // убираем лишнее в начале и конце
		for len(userInput) > 0 {
			elemStr, userInput, _ = strings.Cut(userInput, ",")
			elemInt, _ := strconv.Atoi(elemStr) // преобразуем в int
			if elemInt == 0 { //отсеиваем все лишние элементы - буквы и нули
				continue
			}
			numbers = append(numbers, elemInt) // добавляем отрезанный и преобразованный элемент в слайс
		}
		if len(numbers) == 0 {
			fmt.Println("Ошибка ввода, попробуйте снова")//если буквы или пустой ввод или нули - повторная попытка ввода
			continue
		}
		return numbers
	}
}

func getAvg (numbers *[]int) {
	var sum = 0
	for _, elem := range *numbers { // высчитываем сумму элементов
		sum += elem
	}
	var result float64 = float64(sum) / float64(len(*numbers)) // высчитываем среднее арифметическое
	fmt.Printf("\n\nСреднее арифметическое введённых чисел - %.2f \n\n", result)
}

func getSum (numbers *[]int ) {
	var sum = 0
	for _, elem := range *numbers {
		sum += elem
	}
	fmt.Printf("\n\nСумма введённых чисел - %v \n\n", sum)
}


func getMed (numbers *[]int, baseNum int) {
	var maxNum, minNum = baseNum, baseNum //без BaseNum ругается, что слайс может быть пустой. Проверку на пустоту сделали при вводе числе - поэтому здесь она не нужна
	var differ float64
	for _, elem := range *numbers { // находим максимальное и минимальное числа
		if elem > maxNum {
			maxNum = elem
		}
		if elem < minNum {
			minNum = elem
		}
	} 
	med := baseNum // ставим медианой первый элемент, для того чтобы отталкиваться
	perfectMed := (float64(maxNum + minNum)) / 2 // находим идеальную медиану
	differMed := difference(perfectMed, float64(med)) //разница первого элемента и идеальной медианы, для дальнейшего сравнения
	CHECK:
	for _, elem := range *numbers { // проходимся по слайсу и находим разность между реальными числами и идеальной медианой
		differ = difference(perfectMed, float64(elem))
		if perfectMed == float64(elem) { // если число совпадает с идеальной медианой - прерываем цикл
			med = elem 
			break CHECK
		}
		if differ < differMed { //сравниваем удалённость текущего числа от медианы и удалённость текущей выбранной медианы
			med = elem
			differMed = differ
		}
	} // решил решить собственным алгоритмом, прямолинейным. Алгоритм quickselect пока сложновато
	fmt.Printf("\nМедиана введённых чисел - %d\n\n", med)
}

func difference (middleNumber float64, elem float64) float64 { //для учтния случаев разницы знаков и отрицательных чисел
	if middleNumber > elem {
		return middleNumber - elem
	}
	return elem - middleNumber
}