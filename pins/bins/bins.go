package bins

import (
	"encoding/json"
	"fmt"
	"pins/api"
	"pins/file"
	"time"
)

type Db interface {
	Read () ([]byte, error)
	Save ([]byte) ()
}

type Bin struct {
	Id        string 	`json:"id"`
	Private   bool		`json:"private"`
	CreatedAt time.Time	`json:"createdAt"`
	Name      string	`json:"name"`
}

type respond struct {
	Metadata struct 
	{
		Id        	string 		`json:"id"`
		Private   	bool		`json:"private"`
		CreatedAt 	time.Time	`json:"createdAt"`
		Name      	string		`json:"name"`
		ParentId	string 		`json:"parentId"`
	} `json:"metadata"`
	Record struct {

	} `json:"record"`
}


type BinList struct {
	Bins []Bin
}

type BinListwithDb struct {
	BinList
	Db Db
}

func (binList *BinListwithDb) NewBin (filename string, name string) (error) {
	data, err := file.ReadFile(filename)
	if err != nil {
		return err
	}
	resp, err := api.NewBin(data, name)
	if err != nil {
		return err
	}
	var newBin Bin
	var respStruct respond
	json.Unmarshal(resp, &respStruct)
	newBin.Id = respStruct.Metadata.Id
	newBin.CreatedAt = respStruct.Metadata.CreatedAt
	newBin.Private = respStruct.Metadata.Private
	newBin.Name = respStruct.Metadata.Name
	binList.BinList.Bins = append(binList.BinList.Bins, newBin)
	return nil
}

func NewBinList(db Db) BinListwithDb {
	data, err := db.Read()
	if err != nil || data == nil {
		return BinListwithDb{
			BinList: BinList {
				Bins: 		[]Bin{},
			},
			Db: db,
		}
	}
	var binList BinListwithDb
	err = json.Unmarshal(data, &binList.BinList)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error opening local Binlist, creating new one")
		return BinListwithDb{
			BinList: BinList{
				Bins: 		[]Bin{},
			},
			Db: db,
		}
	}
	return BinListwithDb{
		BinList: binList.BinList,
		Db: db,
	}
	}

func (binList *BinListwithDb) ToBytes () ([]byte, error) {
	file, err := json.Marshal(binList.BinList)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (binList *BinListwithDb) UpdateBin (filename string, id string) error {
	data, err := file.ReadFile(filename)
	if err != nil {
		return err
	}
	err = api.PutBin(data, id)
	if err != nil {
		return err
	}
	return nil 
}

func (binList *BinListwithDb) DeleteBin (id string) (error) {
	err := api.DeleteBin(id)
	if err != nil {
		return err
	}
	for i, elem := range binList.Bins {
		if elem.Id == id {
			binList.Bins = append(binList.Bins[:i], binList.Bins[i+1:]...)
			break
		}
	}
	return nil
}

func (binList *BinListwithDb) GetBin (id string) (error) {
	data, err := api.GetBin(id)
	if err != nil {
		return err
	}
	fmt.Println("Bin id ", id, ":")
	fmt.Println(string(data))
	return nil
}

func (binList *BinListwithDb) OutputList () {
	fmt.Println("outputlist")
}

func (binList *BinListwithDb) CheckBin (data string) (bool) {
	for _, elem := range binList.Bins {
		if elem.Id == data || elem.Name == data {
			return true
		}
	}
	return false
}