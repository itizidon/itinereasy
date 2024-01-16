package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("HELLO WORLD")
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})

	users := r.Group("/users")
	{
		// Handle /users route
		users.GET("/")
	}

	// Start the server
	r.Run(":8080")
}
