package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	if !strings.HasSuffix(name, ".json") {
		err := errors.New("ошибка - это не json файл")
		return nil, err
	}
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return nil, err
	}
	return data, nil
}