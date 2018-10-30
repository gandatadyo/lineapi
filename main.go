package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var cin *gin.Engine

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	cin = gin.Default()
	cin.Delims("{$", "}")
	cin.LoadHTMLGlob("src/templates/*")
	cin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})
	cin.Run(":" + port)
}
