package router

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"tiny_vote/method"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func graphqlHandler() gin.HandlerFunc {
	queryFields := graphql.Fields{
		"query": &graphql.Field{
			Name:        "Query",
			Type:        graphql.Int,
			Description: "Query user's votes",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: method.Query,
		},
		"cas": &graphql.Field{
			Name:        "Query",
			Type:        graphql.String,
			Description: "Get valid ticket",
			Resolve:     method.Cas,
		},
	}

	mutationFields := graphql.Fields{
		"vote": &graphql.Field{
			Name:        "Mutation",
			Type:        graphql.Boolean,
			Description: "Vote for a user",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"ticket": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: method.Vote,
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "Query",
		Fields: queryFields,
	}
	rootMutation := graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: mutationFields,
	}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	// Create GraphQL HTTP handler
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	// 只需要通过Gin简单封装即可
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func SetRouter() {
	Router.GET("/graphql", graphqlHandler())
	Router.POST("/graphql", graphqlHandler())
}
