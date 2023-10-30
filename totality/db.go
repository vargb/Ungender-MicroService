package main

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

var UserDB Users

type DBConnection interface {
	Connection()
}

type Connection struct {
	Location string
}

func parsingDB(file string) (*Users, error) {
	var users Users
	userFile, err := os.Open(file)
	if err != nil {
		logrus.Error("Error in opening the file", err)
		return nil, err
	}
	defer userFile.Close()
	jsonParser := json.NewDecoder(userFile)
	jsonParser.Decode(&users)
	return &users, nil
}

func (con *Connection) Connection() *Users {
	users, err := parsingDB(con.Location)
	if err != nil {
		logrus.Error("Could not connect to MockDB!")
		return nil
	}
	UserDB = *users
	logrus.Infoln("MockDB is connected")
	return users
}

func (con *Connection) GetDBData() *Users {
	return &UserDB
}
