package config

import (
	"bufio"
	"os"
	"strings"
)

func loadEnv () (error) {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "$", "#")
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(parts[0], strings.Trim(parts[1], "\""))
		}
	}
	return scanner.Err()
}

func GetKey () string {
	err := loadEnv()
	if err != nil {
		panic("Its not .env file")
	}
	key := os.Getenv("MASTER_KEY")
	key = strings.ReplaceAll(key, "#", "$")
	if key == "" {
		panic("Its no KEY value in .env file")
	}
	return key
}