package bins

import (
	"fmt"
	"strings"
	"time"
)

type Db interface {
	Read () (*BinListwithDb, error)
	Save (BinListwithDb)
}

type Bin struct {
	Id        int 	`json:"id"`
	Private   bool		`json:"private"`
	CreatedAt time.Time	`json:"createdAt"`
	Name      string	`json:"name"`
}

type BinList struct {
	Bins []Bin
}

type BinListwithDb struct {
	BinList
	db Db
}

func NewBin(binList *BinListwithDb) *BinListwithDb{
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

func NewBinList(db Db) *BinListwithDb {
	binList, err := db.Read()
	if err != nil {
		return &BinListwithDb{
			BinList: BinList{
				Bins: 		[]Bin{},
			},
			db: db,
		}
	}
	return binList
	}

