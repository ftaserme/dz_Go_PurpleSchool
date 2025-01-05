package main

import (
	"flag"
	"fmt"
	"pins/bins"
	"pins/storage"
	"strings"
)

func main() {
	isCreate := flag.Bool("create", false, "Создание Bin")
	isUpdate := flag.Bool("update", false, "Обновление Bin")
	isDelete := flag.Bool("delete", false, "Удаление Bin")
	isGet := flag.Bool("get", false, "Просмотр Bin")
	isList := flag.Bool("list", false, "Вывод Storage")
	file := flag.String("file", "", "Файл Bin")
	name := flag.String("name", "", "Имя Bin")
	id := flag.String("id", "", "ID Bin")
	flag.Parse()
	binsStorage := bins.NewBinList(storage.NewJsonStorage()) //открываем локальное хранилище, если нет - создаём
	switch {
	case *isCreate:
		createBin(&binsStorage, *file, *name)
	case *isUpdate:
		updateBin(&binsStorage, *file, *id)
	case *isDelete:
		deleteBin(&binsStorage, *id)
	case *isGet:
		getBin(&binsStorage, *id)
	case *isList:
		getList(&binsStorage)
	default:
		fmt.Println("Error: no action")//от JsonBin ошибки приходят на en, поэтому решил всё прописывать на en для единообразия
	}
}

func createBin(binStorage *bins.BinListwithDb, filename string, name string) {
	if filename == "" {
		fmt.Println("Error: no file for creating Bin")
		return
	}
	if name == "" {
		name, _ = strings.CutSuffix(filename, ".json")// если не введено имя - по умолчанию используем имя файла
	}
	if binStorage.CheckBin(name) {
		fmt.Println("Error: Bin ", name, " already exist")// проверка на задвоение имён бин
		return
	}
	err := binStorage.NewBin(filename, name)
	if err != nil {
		fmt.Println("Error creating Bin: ")
		fmt.Println(err)
		return
	}
	data, err := binStorage.ToBytes()
	if err != nil {
		fmt.Println("Error of saving local BinList")
		return
	}
	binStorage.Db.Save(data)
	fmt.Println("Bin created")
}

func updateBin (binStorage *bins.BinListwithDb, filename string, id string) {
	if !binStorage.CheckBin(id) { // проверка есть ли что обновлять
	 	fmt.Println("Error: no such Bin in database")
	 	return
	}
	err := binStorage.UpdateBin(filename, id)
	if err != nil {
		fmt.Println("Error update Bin:")
		fmt.Println(err)
		return
	}
	fmt.Println("Bin ", id, " updated")
}

func deleteBin (binStorage *bins.BinListwithDb, id string) {
	if !binStorage.CheckBin(id) {
		fmt.Println("Error: no such Bin in database") // проверка есть ли что удалять 
		return
   }
   err := binStorage.DeleteBin(id)
   if err != nil {
		fmt.Println("Error deleting Bin:")
		fmt.Println(err)
		return
   }
   data, err := binStorage.ToBytes()
   if err != nil {
	   fmt.Println("Error of saving local BinList")
	   return
   }
   binStorage.Db.Save(data)
   fmt.Println("Bin ", id, " deleted")
}

func getBin (binStorage *bins.BinListwithDb, id string) {
	if !binStorage.CheckBin(id) {
		fmt.Println("Error: no such Bin in database")
		return
   }
   err := binStorage.GetBin(id)
   if err != nil {
		fmt.Println("Error geting Bin:")
		fmt.Println(err)
		return
   }
}

func getList (binStorage *bins.BinListwithDb) {
	fmt.Println("List of created Bins:")
	for i, elem := range binStorage.Bins {
		fmt.Println("\nBin № ", i+1)
		fmt.Println("Name: ", elem.Name)
		fmt.Println("ID: ", elem.Id)
	}
}