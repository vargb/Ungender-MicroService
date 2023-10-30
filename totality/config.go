package main

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Port      string    `json:"port"`
	UserValid UserValid `json:"userValidation"`
}

type UserValid struct {
	Userid   string `json:"userid"`
	Password string `json:"password"`
}

func Parse(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		logrus.Error(err)
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	logrus.Info("parsed config file")
	return config
}
