package main

import (
	"database/sql"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sammychinedu2ky/synthesis/graph"
	"github.com/sammychinedu2ky/synthesis/graph/generated"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	db, err := sql.Open("mysql", "test:test@tcp(db:3306)/stars?parseTime=true")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	
	if err != nil{
		panic("error connecting to DB")
	}

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Db:db}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Setting up Gin
	r := gin.Default()
	r.Use(gin.Recovery(),gin.Logger())
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":"+port)
}
