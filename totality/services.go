package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, GetDBData())
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	m := make(map[string]int)
	for i := range id {
		if _, err := strconv.Atoi(string(id[i])); err == nil {
			m[string(id[i])] = i
		}
	}
	users := GetDBData()
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
