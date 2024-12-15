package bins

import (
	"encoding/json"
	"fmt"
	"pins/file"
	"strings"
	"time"
)

type Bin struct {
	Id        int 	`json:"id"`
	Private   bool		`json:"private"`
	CreatedAt time.Time	`json:"createdAt"`
	Name      string	`json:"name"`
}

type BinList struct {
	Bins []Bin
	FilePath string
}

func NewBin(binList *BinList) *BinList{
	fmt.Println("Введите id: ")
	var id int
	fmt.Scan(&id)
	for id <=0 {
		fmt.Printf("Ошибка чтения, попробуйте снова: ")
		fmt.Scan(&id)
	}
	input := ""
	for {
		fmt.Println("Приватность(Y/N): ")
		fmt.Scan(&input)
		input = strings.ToUpper(input)
		if input == "Y" || input == "N" {
			break
		}
		fmt.Println("Ошибка, попробуйте снова")
	}
	private := true
	if input == "N" {
		private = false
	}
	fmt.Printf("Введите имя bin: ")
	var name string
	fmt.Scan(&name)
	bin := Bin {
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	binList.Bins = append(binList.Bins, bin)
	return binList
}

func NewBinList() *BinList {
	baseStoragePath := "StorageBin.json"
	file, err := file.ReadFile(baseStoragePath)
	if err != nil {
		return &BinList{
			Bins: 		[]Bin{},
			FilePath: 	baseStoragePath,
		}
	}
	var binList BinList
	err = json.Unmarshal(file, &binList)
	if err != nil {
		fmt.Printf("Не удалось разобрать файл %v\n", baseStoragePath)
		return  &BinList {
			Bins: 		[]Bin{},
			FilePath: 	baseStoragePath,
		}
	}
	return &binList
	}

