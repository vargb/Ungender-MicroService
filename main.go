package main

import (
	"hopeugetknowuwont/config"
	"hopeugetknowuwont/domain"
	"hopeugetknowuwont/postgres"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-pg/pg/v9"
	"github.com/sirupsen/logrus"
)

//when trying to run graphql code is breaking

func main() {
	config := config.Parse("config.json")

	DB := postgres.New(&pg.Options{
		User:     config.Psql.User,
		Password: config.Psql.Password,
		Database: config.Psql.Dbname,
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	logrus.Info("starting server")
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Port
	}

	userRepo := postgres.UsersRepo{DB: DB}

	_ = domain.NewDomain(userRepo, postgres.GarageRepo{DB: DB})
	//graph := graph.Config{Resolvers: {Domain: dom}}

	http.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	//http.Handle("/query",handler.GraphQL())

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
