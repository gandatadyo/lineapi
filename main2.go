package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

var cin1 *gin.Engine

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Print("poemicu ganda tt")
	client := &http.Client{}
	bot, err := linebot.New("<24ee6cbebfd41457136563b41f8184d6>", "<LZANYzm90f3h6eUyUl4PNce0sTjKQeojp7/E+F3NAyBW/ltsD3UCqsZdTxcyiEk+76+IXhfmzPKjoP16JDCMnzVI/EI5nQarf3h5ngNBDctTepFQAhtP/25yj5EV72OkNtIFPYxzgQPLwssXgQ8gpgdB04t89/1O/w1cDnyilFU=>", linebot.WithHTTPClient(client))
	events, err := bot.ParseRequest(r)
	// if err != nil {
	// 	fmt.Println("Not Work 1")
	// }

	fmt.Print("poemicu ganda hh")
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		fmt.Print("poemicu ganda 1")
		if event.Type == linebot.EventTypeMessage {
			fmt.Print("poemicu ganda 2")
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			}
		}

		if event.Type == linebot.EventTypeMessage {
			// Do Something...
			fmt.Print("poemicu ganda")
			// userID := event.Source.UserID
			// groupID := event.Source.GroupID
			// RoomID := event.Source.RoomID

			// replyToken := event.ReplyToken

			// // Create Message
			// leftBtn := linebot.NewMessageAction("left", "left clicked")
			// rightBtn := linebot.NewMessageAction("right", "right clicked")

			// template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)

			// message := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)

			// // SendMessage
			// var messages []linebot.Message

			// append some message to messages
			// _, err := bot.PushMessage(ID, messages).Do()
			// if err != nil {
			// 	// Do something when some bad happened
			// }
		}
	}

}

func main1() {

	fmt.Println("Server Running")
	http.HandleFunc("/callback", login)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
