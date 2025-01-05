package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	if !strings.HasSuffix(name, ".json") {
		err := errors.New("Error: It's not a json file")
		return nil, err
	}
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile (content []byte, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Local BinList updated")
}