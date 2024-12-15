package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

//	"fmt"
//	"pins/bins"
//	"pins/storage"

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не удалось найти env файл")
	}
}
