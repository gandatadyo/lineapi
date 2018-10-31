package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

var cin *gin.Engine

func main() {
	// bot, err := linebot.New(
	// 	os.Getenv("977b5cf3276bce134af72ad4e6e45707"), //CHANNEL_SECRET
	// 	os.Getenv("nFAf7VS/RulMCFabBKJfgIHvDxfKDpENufTZ4lWQ1M0AGUTKMJG+bVMw+uA/Twon2qYS/EDq/B3AHy/JnVZSnq25RG7SPpEF8QL87YFqDzzne6ZHYA/bhmu6jE3k6/1DG/+a4sp5Bo3+7ZOAwoSLbQdB04t89/1O/w1cDnyilFU="), //CHANNEL_TOKEN
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	client := &http.Client{}
	bot, err := linebot.New("<977b5cf3276bce134af72ad4e6e45707>", "<nFAf7VS/RulMCFabBKJfgIHvDxfKDpENufTZ4lWQ1M0AGUTKMJG+bVMw+uA/Twon2qYS/EDq/B3AHy/JnVZSnq25RG7SPpEF8QL87YFqDzzne6ZHYA/bhmu6jE3k6/1DG/+a4sp5Bo3+7ZOAwoSLbQdB04t89/1O/w1cDnyilFU=>", linebot.WithHTTPClient(client))

	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}

	// port := os.Getenv("PORT")

	// if port == "" {
	// 	port = "1000"
	// 	// log.Fatal("$PORT must be set")
	// }

	// cin = gin.Default()
	// cin.Delims("{$", "}")
	// cin.LoadHTMLGlob("src/templates/*")
	// cin.POST("/", func(c *gin.Context) {
	// 	var valheader = c.Request.Header.Get("valheader")

	// 	buf := make([]byte, 1024)
	// 	num, _ := c.Request.Body.Read(buf)
	// 	valbody := string(buf[0:num])
	// 	c.HTML(http.StatusOK, "home.html", gin.H{"valheader": valheader, "valbody": string(valbody)})
	// })
	// cin.GET("/", func(c *gin.Context) {
	// 	var valheader = c.Request.Header.Get("valheader")

	// 	// c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	// 	c.HTML(http.StatusOK, "home.html", gin.H{"valheader": valheader, "valbody": " THIS IS METHOD GET"})
	// })
	// cin.Run(":" + port)
}
