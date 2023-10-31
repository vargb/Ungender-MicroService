package main

import (
	"hopeugetknowuwont/ungender/config"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/youtube-golang-graphql-tutorial-master/postgres"
)

func main() {
	config := config.Parse("config.json")

	psqlconn := postgres.New(&pg.Options{
		User:     config.Psql.User,
		Password: config.Psql.Password,
		Database: config.Psql.Dbname,
	})

	defer psqlconn.Close()

	logrus.Info("starting server")
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Port
	}

	// dom := domain.NewDomain(postgres.UsersRepo{DB: DB}, postgres.GarageRepo{DB: DB})
	// graph := graph.Config{Resolvers: {Domain: dom}}

	// http.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	// //http.Handle("/query",handler.GraphQL())

	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
