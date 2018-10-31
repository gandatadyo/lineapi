package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var cin *gin.Engine

// Channel ID
// 	1617989157
// Channel secret
// 	bcce677e544f65cb0ad30e2c2949b435

func lineapi() {
	// defer req.Body.Close()
	// body, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	// ...
	// }
	// decoded, err := base64.StdEncoding.DecodeString(req.Header.Get("X-Line-Signature"))
	// if err != nil {
	// 	// ...
	// }
	// hash := hmac.New(sha256.New, []byte("<channel secret>"))
	// hash.Write(body)

	// Compare decoded signature and `hash.Sum(nil)` by using `hmac.Equal`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "1000"
		// log.Fatal("$PORT must be set")
	}

	cin = gin.Default()
	cin.Delims("{$", "}")
	cin.LoadHTMLGlob("src/templates/*")
	cin.POST("/", func(c *gin.Context) {
		var valheader = c.Request.Header.Get("valheader")

		buf := make([]byte, 1024)
		num, _ := c.Request.Body.Read(buf)
		valbody := string(buf[0:num])
		c.HTML(http.StatusOK, "home.html", gin.H{"valheader": valheader, "valbody": string(valbody)})
	})
	cin.GET("/", func(c *gin.Context) {
		var valheader = c.Request.Header.Get("valheader")

		// c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
		c.HTML(http.StatusOK, "home.html", gin.H{"valheader": valheader, "valbody": " THIS IS METHOD GET"})
	})
	cin.Run(":" + port)
}
