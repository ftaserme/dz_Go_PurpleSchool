package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"pins/bins"
	"pins/file"
)

type JsonStorage struct {
	filename string
}

func NewJsonStorage(name string) *JsonStorage {
	return &JsonStorage{
		filename: name,
	}
}

func (db *JsonStorage) Save (storage bins.BinListwithDb) {
	context, err := json.Marshal(storage.BinList)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.Write(context)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}

func (db *JsonStorage) Read() (*bins.BinListwithDb, error) {
	storage := bins.NewBinList(db)
	data, err := file.ReadFile(db.filename)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	err = json.Unmarshal(data, &storage)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return storage, nil
}
