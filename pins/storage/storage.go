package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"pins/bins"
	"pins/file"
)

func SaveStorage(storage bins.BinList) {
	context, err := json.Marshal(storage)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Create(storage.FilePath)
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

func ReadStorage() (*bins.BinList) {
	storage := bins.NewBinList()
	data, err := file.ReadFile(storage.FilePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(data, &storage)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return storage
}
