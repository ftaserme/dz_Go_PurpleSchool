package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"pins/config"
	"strings"
)

func NewBin (data []byte, name string) ([]byte, error) {
	respData, err := sendReq("POST", data, name)
	if err != nil {
		return nil, err
	}
	return respData, nil
}

func PutBin (data []byte, id string) (error) {
	_, err := sendReq("PUT", data, id)
	if err != nil {
		return err
	}
	return nil
}

func GetBin (id string) ([]byte, error) {
	respData, err := sendReq("GET", nil, id)
	if err != nil {
		return nil, err
	}
	return respData, nil
}

func DeleteBin (id string) (error) {
	_, err := sendReq("DELETE", nil, id)
	if err != nil {
		return err
	}
	return nil
}

func sendReq (reqType string, data []byte, binDate string) ([]byte, error) {
	url := "https://api.jsonbin.io/v3/b/"
	if reqType != "POST" {
		url = url + binDate
	}
	req, _ := http.NewRequest(reqType, url, strings.NewReader(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", config.GetKey())
	if reqType == "POST" {
		req.Header.Set("X-Bin-Private", getPrivate())
		req.Header.Set("X-Bin-Name", binDate)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	var checkRes struct {
		Message string `json:"message"`
	}
	json.Unmarshal(respData, &checkRes)
	if checkRes.Message != "" && checkRes.Message != "Bin deleted successfully" {
		return nil, errors.New(checkRes.Message)
	}
	return respData, err
}

func getPrivate() string {
	var userInput string
	fmt.Print("Is private(Y/N): ")
	for {
		fmt.Scan(&userInput)
		userInput = strings.ToUpper(userInput)
		if userInput == "" {
			fmt.Print("Error, try again: ")
			continue
		}
		break
	}
	if userInput == "N" {
		return "false"
	}
	return "true"
}