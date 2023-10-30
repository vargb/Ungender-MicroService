package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (con *Connection) GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, con.GetDBData())
}

func (con *Connection) GetById(c *gin.Context) {
	id := c.Param("id")
	m := make(map[string]int)
	for i := range id {
		if _, err := strconv.Atoi(string(id[i])); err == nil {
			m[string(id[i])] = i
		}
	}
	users := con.GetDBData()
	flag := false
	var userList []User
	for _, user := range users.Users {
		if _, ok := m[strconv.Itoa(user.Id)]; ok {
			userList = append(userList, user)
			flag = true
		}
	}
	if !flag {
		c.IndentedJSON(http.StatusNotFound, gin.H{"HeadsUp": "ID not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, userList)
}

func (con *Connection) Post(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	users := con.GetDBData()
	users.Users = append(users.Users, newUser)
	file, err := json.Marshal(users)
	if err != nil {
		logrus.Error("Error in writing to DB")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"HeadsUp": "Error in posting data"})
		return
	}
	_ = os.WriteFile("mockdb.json", file, 0644)
	c.IndentedJSON(http.StatusCreated, gin.H{"HeadsUp": "New User Created"})
}
