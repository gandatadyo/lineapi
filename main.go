package main

// update 1 Nov 2018 1.1
import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

var cin *gin.Engine

func main() {
	// init bot
	client := &http.Client{}
	bot, err := linebot.New("24ee6cbebfd41457136563b41f8184d6", "LZANYzm90f3h6eUyUl4PNce0sTjKQeojp7/E+F3NAyBW/ltsD3UCqsZdTxcyiEk+76+IXhfmzPKjoP16JDCMnzVI/EI5nQarf3h5ngNBDctTepFQAhtP/25yj5EV72OkNtIFPYxzgQPLwssXgQ8gpgdB04t89/1O/w1cDnyilFU=", linebot.WithHTTPClient(client))
	if err != nil {
		fmt.Println("Not Work 1")
		log.Print(err)
	}

	// init port
	port := os.Getenv("PORT")
	if port == "" {
		port = "1000"
		fmt.Println("$PORT must be set")
		log.Print(err)
	}

	// init gin
	cin = gin.Default()
	cin.Delims("{$", "}")
	// cin.LoadHTMLGlob("src/templates/*")

	cin.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Start")
	})

	// event all
	cin.POST("/callback", func(c *gin.Context) {
		fmt.Println("Work start POST")
		events, err := bot.ParseRequest(c.Request)

		fmt.Println("Work 1.1")
		if err != nil {
			fmt.Println("Error 1.1")
			log.Print(err)
		}
		for _, event := range events {
			fmt.Println("Work 1.2")

			fmt.Println(event.Source.UserID)
			fmt.Println(event.Source.GroupID)
			fmt.Println(event.Source.RoomID)
			fmt.Println(event.ReplyToken)

			if event.Type == linebot.EventTypeMessage {
				fmt.Println("Work EventTypeMessage")
				if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("hello, iam good")).Do(); err != nil {
					log.Print(err)
				}

			}
			if event.Type == linebot.EventTypeFollow {
				fmt.Println("Work EventTypeFollow")

			}
			if event.Type == linebot.EventTypeUnfollow {
				fmt.Println("Work EventTypeUnfollow")

			}
			if event.Type == linebot.EventTypeJoin {
				fmt.Println("Work EventTypeJoin")

			}
			if event.Type == linebot.EventTypeLeave {
				fmt.Println("Work EventTypeLeave")

			}
			if event.Type == linebot.EventTypeBeacon {
				fmt.Println("Work EventTypeBeacon")

			}
		}

		// var valheader = c.Request.Header.Get("valheader")
		// buf := make([]byte, 1024)
		// num, _ := c.Request.Body.Read(buf)
		// valbody := string(buf[0:num])
		// c.HTML(http.StatusOK, "home.html", gin.H{"valheader": valheader, "valbody": string(valbody)})
	})
	cin.GET("/callback", func(c *gin.Context) {
		fmt.Println("Work start GET")
		events, err := bot.ParseRequest(c.Request)
		fmt.Println("Work 1.1")
		if err != nil {
			fmt.Println("Error 1.1")
			log.Print(err)
		}
		for _, event := range events {
			fmt.Println("Work 1.2")

			fmt.Println(event.Source.UserID)
			fmt.Println(event.Source.GroupID)
			fmt.Println(event.Source.RoomID)
			fmt.Println(event.ReplyToken)

			if event.Type == linebot.EventTypeMessage {
				fmt.Println("Work EventTypeMessage")

			}
			if event.Type == linebot.EventTypeFollow {
				fmt.Println("Work EventTypeFollow")

			}
			if event.Type == linebot.EventTypeUnfollow {
				fmt.Println("Work EventTypeUnfollow")

			}
			if event.Type == linebot.EventTypeJoin {
				fmt.Println("Work EventTypeJoin")

			}
			if event.Type == linebot.EventTypeLeave {
				fmt.Println("Work EventTypeLeave")

			}
			if event.Type == linebot.EventTypeBeacon {
				fmt.Println("Work EventTypeBeacon")

			}
		}

		// var valheader = c.Request.Header.Get("valheader")
		// // c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
		// c.HTML(http.StatusOK, "home.html", gin.H{"valheader": valheader, "valbody": " THIS IS METHOD GET"})
	})

	cin.GET("/pushmessage", func(c *gin.Context) {
		fmt.Println("Work start PUSH MESSAGE")

		if _, err := bot.PushMessage("U77e1544ac9ae112f2bde7542bd61df65", linebot.NewTextMessage("hello, iam from golang")).Do(); err != nil {
			log.Print(err)
		}
	})

	cin.GET("/pushmessage2", func(c *gin.Context) {
		fmt.Println("Work start PUSH MESSAGE")

		if _, err := bot.PushMessage("U77e1544ac9ae112f2bde7542bd61df65", linebot.NewTextMessage("hello, iam from golang pushmessage2")).Do(); err != nil {
			log.Print(err)
		}
	})

	cin.GET("/pushmessage3 ", func(c *gin.Context) {
		fmt.Println("Work start FLEX")

		containerBubble := &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Body: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "Hello,",
					},
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "World!",
					},
				},
			},
		}

		if _, err := bot.PushMessage("U77e1544ac9ae112f2bde7542bd61df65", linebot.NewFlexMessage("alt text", containerBubble)).Do(); err != nil {
			log.Print(err)
		}
	})

	cin.Run(":" + port)
}
