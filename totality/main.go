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
	router.GET("/getusers", userConnection.GetAll)
	router.GET("/getusers/:id", userConnection.GetById)
	router.POST("/postuser", userConnection.Post)
	router.Run(":" + config.Port)
}
