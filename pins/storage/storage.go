package storage

import (
	"fmt"
	"pins/file"
)

type JsonStorage struct {
	filename string
}

func NewJsonStorage() *JsonStorage {
	return &JsonStorage{
		filename: "StorageBin.json",
	}
}

func (db *JsonStorage) Save (data []byte) () {
	file.WriteFile(data, db.filename)
}

func (db *JsonStorage) Read() ([]byte, error) {
	data, err := file.ReadFile(db.filename)
	if err != nil || data == nil {
		fmt.Println(err)
		return nil, nil
	}
	return data, nil
}
