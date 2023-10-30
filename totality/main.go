package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	config := Parse("config.json")
	userConnection := Connection{Location: "mockdb.json"}
	users := userConnection.Connection()
	if users == nil {
		return
	}
	router := gin.Default()
	router.GET("/getusers", GetAll)
	router.GET("/getusers/:id", GetById)
	// router.POST("/postuser", Post)
	router.Run(":" + config.Port)
}
