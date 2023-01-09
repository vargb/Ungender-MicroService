package config

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Port   string `json:"port"`
	Psql   Psql   `json:"psql"`
	Auth   Auth   `json:"auth"`
	Garage Garage `json:"garageservice"`
}

type Psql struct {
	Host     string `json:"host"`
	Sqlport  string `json:"sqlport"`
	User     string `json:"user"`
	Password string `json:"pass"`
	Dbname   string `json:"dbname"`
}

type Auth struct {
	SignUp   string `json:"signup"`
	Login    string `json:"login"`
	Logout   string `json:"logout"`
	SezCheck string `json:"sezcheck"`
}

type Garage struct {
	Get    string `json:"get"`
	Book   string `json:"book"`
	Return string `json:"return"`
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
