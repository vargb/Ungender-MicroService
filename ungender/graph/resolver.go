package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func GraphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		logrus.Error("could not retrieve gin.Context")
		return nil, errors.New("count not retrieve gin.Context")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		logrus.Error("gin.Context has wrong type")
		return nil, errors.New("gin.Context has wrong type")
	}
	return gc, nil
}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
