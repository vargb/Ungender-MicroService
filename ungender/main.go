package main

import (
	"hopeugetknowuwont/ungender/config"
	"hopeugetknowuwont/ungender/graph"
	potgres "hopeugetknowuwont/ungender/pgres"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	config := config.Parse("config.json")
	router := gin.Default()
	_, err := potgres.Init(config)
	if err != nil {
		logrus.Info(err)
		return
	}
	router.GET("/getall", potgres.GetPqHandler().GetAll)
	router.POST("/postcar", potgres.GetPqHandler().PostGarage)
	router.POST("postuser", potgres.GetPqHandler().PostUser)
	//psqlconn := postgres.Connect()
	router.POST("/query", graph.GraphqlHandler())
	router.GET("/", graph.PlaygroundHandler())

	router.Run(":" + config.Port)
}
