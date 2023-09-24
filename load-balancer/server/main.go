package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleHomeRequest(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "successful",
	})
}

func main() {
	fmt.Println("This is a sample server")

	router := gin.Default()

	router.Use(corsMiddleware())

	router.GET("/", handleHomeRequest)

	router.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		g.Writer.Header().Set("Access-control-allow-methods", "GET")
		g.Writer.Header().Set("Access-control-allow-headers", "Content-Type, mode, authorization")
		g.Writer.Header().Set("Access-control-allow-origin", "*")

		if g.Request.Method == http.MethodOptions {
			g.AbortWithStatus(http.StatusOK)
			return
		}

		g.Next()
	}
}
