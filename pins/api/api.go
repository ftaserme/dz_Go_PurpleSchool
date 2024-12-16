package api

import (
	"pins/config"
)

func GetConfig () string {
	config := config.NewConfig()
	return config.Key
}
