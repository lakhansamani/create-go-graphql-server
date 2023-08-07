package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/lakhansamani/create-go-graphql-server/graph"
	"github.com/lakhansamani/create-go-graphql-server/internal/middleware"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/v1/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the router for the server
func New() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middleware.DefaultStructuredLogger())
	router.Use(middleware.GinContextToContextMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Recovery())
	// version 1
	apiV1 := router.Group("/v1")
	apiV1.GET("/", playgroundHandler())
	apiV1.POST("/graphql", graphqlHandler())
	return router
}
