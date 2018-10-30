package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var cin *gin.Engine

func main() {
	cin = gin.Default()
	cin.Delims("{$", "}")
	cin.LoadHTMLGlob("src/templates/*")
	cin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})
	cin.Run(":1000")
}
