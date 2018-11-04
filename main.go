package main

// update 03 Nov 2018 Update 5
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
	var channeltoken = `LZANYzm90f3h6eUyUl4PNce0sTjKQeojp7/E+F3NAyBW/ltsD3UCqsZdTxcyiEk+76+IXhfmzPKjoP16JDCMnzVI/EI5nQarf3h5ngNBDctTepFQAhtP/25yj5EV72OkNtIFPYxzgQPLwssXgQ8gpgdB04t89/1O/w1cDnyilFU=`

	client := &http.Client{}
	bot, err := linebot.New("24ee6cbebfd41457136563b41f8184d6", channeltoken, linebot.WithHTTPClient(client))
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
		fmt.Println("Work start callback")
		events, err := bot.ParseRequest(c.Request)

		fmt.Println("Work 1.1")
		if err != nil {
			fmt.Println("Error 1.1")
			log.Print(err)
		}
		for _, event := range events {
			fmt.Println("Work 1.2")

			// fmt.Println(event.Source.UserID)
			// fmt.Println(event.Source.GroupID)
			// fmt.Println(event.Source.RoomID)
			// fmt.Println(event.ReplyToken)

			if event.Type == linebot.EventTypeMessage {
				fmt.Println("Work EventTypeMessage")
				if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(fmt.Sprint("hello, iam good ", event.Message))).Do(); err != nil {
					log.Print(err)
				}

			}
			if event.Type == linebot.EventTypeFollow {
				fmt.Println("Work EventTypeFollow")

				if Testdata, err := bot.GetProfile(event.Source.UserID).Do(); err != nil {
					log.Print(err)
				} else {
					// fmt.Println("UserID -> ", Testdata.UserID)
					// fmt.Println("DisplayName -> ", Testdata.DisplayName)
					// fmt.Println("PictureURL -> ", Testdata.PictureURL)
					// fmt.Println("StatusMessage -> ", Testdata.StatusMessage)
					// Message
					var message = fmt.Sprint("Selamat Datang ", Testdata.DisplayName, "  ")
					if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(message)).Do(); err != nil {
						log.Print(err)
					}

					// Sticker
					if _, err := bot.PushMessage(event.Source.UserID, linebot.NewStickerMessage("1", "106")).Do(); err != nil {
						log.Print(err)
					}

					// Button Template
					leftBtn := linebot.NewMessageAction("Start Bot", "Start Bot")
					rightBtn := linebot.NewMessageAction("Information", "Information")
					template := linebot.NewConfirmTemplate("Hi..", leftBtn, rightBtn)
					if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTemplateMessage("Bot call you..", template)).Do(); err != nil {
						log.Print(err)
					}

				}
			}
			// if event.Type == linebot.EventTypeUnfollow {
			// 	fmt.Println("Work EventTypeUnfollow")

			// }
			// if event.Type == linebot.EventTypeJoin {
			// 	fmt.Println("Work EventTypeJoin")

			// }
			// if event.Type == linebot.EventTypeLeave {
			// 	fmt.Println("Work EventTypeLeave")

			// }
			// if event.Type == linebot.EventTypeBeacon {
			// 	fmt.Println("Work EventTypeBeacon")

			// }
		}
	})

	cin.GET("/menu", func(c *gin.Context) {
		// var iduser = c.Param("iduser")
		// fmt.Println("Work start menu")

		leftBtn := linebot.NewMessageAction("Start Bot", "Start Bot")
		rightBtn := linebot.NewMessageAction("Information", "Information")
		template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)
		// message := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)

		if _, err := bot.PushMessage("U77e1544ac9ae112f2bde7542bd61df65", linebot.NewTemplateMessage("Hellow ", template)).Do(); err != nil {
			log.Print(err)
		}
	})

	cin.GET("/pushmessage", func(c *gin.Context) {
		fmt.Println("Work start pushmessage")

		if _, err := bot.PushMessage("U77e1544ac9ae112f2bde7542bd61df65", linebot.NewTextMessage("hello, iam from golang")).Do(); err != nil {
			log.Print(err)
		}
	})

	cin.GET("/multi", func(c *gin.Context) {
		fmt.Println("Work start multi message")
		// var to = `["U77e1544ac9ae112f2bde7542bd61df65","U77e1544ac9ae112f2bde7542bd61df65"]`
		var to []string
		to = append(to, "U77e1544ac9ae112f2bde7542bd61df65")
		to = append(to, "U77e1544ac9ae112f2bde7542bd61df65")

		if _, err := bot.Multicast(to, linebot.NewTextMessage("hello, iam from multi message")).Do(); err != nil {
			log.Print(err)
		}
	})

	cin.GET("/flex1", func(c *gin.Context) {
		fmt.Println("Work start flex1")

		contentsCarousel := &linebot.CarouselContainer{
			Type: linebot.FlexContainerTypeCarousel,
			Contents: []*linebot.BubbleContainer{
				&linebot.BubbleContainer{
					Type: linebot.FlexContainerTypeBubble,
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "First bubble",
							},
						},
					},
				},
				&linebot.BubbleContainer{
					Type: linebot.FlexContainerTypeBubble,
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "Second bubble",
							},
						},
					},
				},
			},
		}

		if _, err := bot.PushMessage("U77e1544ac9ae112f2bde7542bd61df65", linebot.NewFlexMessage("Flex message Carousel", contentsCarousel)).Do(); err != nil {
			log.Print(err)
		}

	})

	cin.GET("/flex2", func(c *gin.Context) {
		fmt.Println("Work start flex2")

		contentsBubbleButton := &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Header: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "Tittle,",
					},
				},
			},
			Hero: &linebot.ImageComponent{
				Type:   linebot.FlexComponentTypeImage,
				URL:    "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Image_created_with_a_mobile_phone.png/1200px-Image_created_with_a_mobile_phone.png",
				Size:   linebot.FlexImageSizeTypeFull,
				Action: &linebot.URIAction{"Label", "https://dark-asylum-97180.herokuapp.com/flex1"},
			},
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
			Footer: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "Footer,",
					},
					&linebot.ButtonComponent{
						Type:   linebot.FlexComponentTypeButton,
						Action: &linebot.URIAction{"Label", "https://dark-asylum-97180.herokuapp.com/flex1"},
					},
				},
			},
		}

		if _, err := bot.PushMessage("U77e1544ac9ae112f2bde7542bd61df65", linebot.NewFlexMessage("Flex message Bubble", contentsBubbleButton)).Do(); err != nil {
			log.Print(err)
		}

	})

	cin.GET("/flex3", func(c *gin.Context) {
		fmt.Println("Work start flex3")

		contentsBubbleButton := &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Header: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "Tittle,",
					},
				},
			},
			Hero: &linebot.ImageComponent{
				Type:   linebot.FlexComponentTypeImage,
				URL:    "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Image_created_with_a_mobile_phone.png/1200px-Image_created_with_a_mobile_phone.png",
				Size:   linebot.FlexImageSizeTypeFull,
				Action: &linebot.URIAction{"Label", "https://dark-asylum-97180.herokuapp.com/flex1"},
			},
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
			Footer: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeVertical,
				Contents: []linebot.FlexComponent{
					&linebot.ButtonComponent{
						Type:   linebot.FlexComponentTypeButton,
						Action: &linebot.URIAction{"Ok", "https://dark-asylum-97180.herokuapp.com/flex1"},
					},
					&linebot.ButtonComponent{
						Type:   linebot.FlexComponentTypeButton,
						Action: &linebot.URIAction{"Cancel", "https://dark-asylum-97180.herokuapp.com/flex1"},
					},
				},
			},
		}

		if _, err := bot.PushMessage("U77e1544ac9ae112f2bde7542bd61df65", linebot.NewFlexMessage("Flex message Bubble", contentsBubbleButton)).Do(); err != nil {
			log.Print(err)
		}

	})

	cin.GET("/flex4", func(c *gin.Context) {
		fmt.Println("Work start flex4")

		contentsBubbleButton1 := &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Header: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "Tittle,",
					},
				},
			},
			Hero: &linebot.ImageComponent{
				Type:   linebot.FlexComponentTypeImage,
				URL:    "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Image_created_with_a_mobile_phone.png/1200px-Image_created_with_a_mobile_phone.png",
				Size:   linebot.FlexImageSizeTypeFull,
				Action: &linebot.URIAction{"Label", "https://dark-asylum-97180.herokuapp.com/flex1"},
			},
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
			Footer: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeVertical,
				Contents: []linebot.FlexComponent{
					&linebot.ButtonComponent{
						Type:   linebot.FlexComponentTypeButton,
						Action: &linebot.URIAction{"Ok", "https://dark-asylum-97180.herokuapp.com/flex1"},
					},
					&linebot.ButtonComponent{
						Type:   linebot.FlexComponentTypeButton,
						Action: &linebot.URIAction{"Cancel", "https://dark-asylum-97180.herokuapp.com/flex1"},
					},
				},
			},
		}

		contentsBubbleButton2 := &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Header: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "Tittle,",
					},
				},
			},
			Hero: &linebot.ImageComponent{
				Type:   linebot.FlexComponentTypeImage,
				URL:    "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Image_created_with_a_mobile_phone.png/1200px-Image_created_with_a_mobile_phone.png",
				Size:   linebot.FlexImageSizeTypeFull,
				Action: &linebot.URIAction{"Label", "https://dark-asylum-97180.herokuapp.com/flex1"},
			},
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
			Footer: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeVertical,
				Contents: []linebot.FlexComponent{
					&linebot.ButtonComponent{
						Type:   linebot.FlexComponentTypeButton,
						Action: &linebot.URIAction{"Ok", "https://dark-asylum-97180.herokuapp.com/flex1"},
					},
					&linebot.ButtonComponent{
						Type:   linebot.FlexComponentTypeButton,
						Action: &linebot.URIAction{"Cancel", "https://dark-asylum-97180.herokuapp.com/flex1"},
					},
				},
			},
		}

		contentsCarousel := &linebot.CarouselContainer{
			Type: linebot.FlexContainerTypeCarousel,
			Contents: []*linebot.BubbleContainer{
				contentsBubbleButton1,
				contentsBubbleButton2,
			},
		}

		if _, err := bot.PushMessage("U77e1544ac9ae112f2bde7542bd61df65", linebot.NewFlexMessage("Flex message Bubble", contentsCarousel)).Do(); err != nil {
			log.Print(err)
		}

	})

	cin.GET("/profileapi", func(c *gin.Context) {
		fmt.Println("Work start profileapi")

		// type data
		// type UserProfileResponse struct {
		// 	UserID        string `json:"userId"`
		// 	DisplayName   string `json:"displayName"`
		// 	PictureURL    string `json:"pictureUrl"`
		// 	StatusMessage string `json:"statusMessage"`
		// }

		if Testdata, err := bot.GetProfile("U77e1544ac9ae112f2bde7542bd61df65").Do(); err != nil {
			log.Print(err)
		} else {
			fmt.Println("UserID -> ", Testdata.UserID)
			fmt.Println("DisplayName -> ", Testdata.DisplayName)
			fmt.Println("PictureURL -> ", Testdata.PictureURL)
			fmt.Println("StatusMessage -> ", Testdata.StatusMessage)
			var data = fmt.Sprint(Testdata.UserID, "  ", Testdata.DisplayName, "  ", Testdata.PictureURL)
			if _, err := bot.PushMessage("U77e1544ac9ae112f2bde7542bd61df65", linebot.NewTextMessage(data)).Do(); err != nil {
				log.Print(err)
			}
		}

	})

	cin.GET("/contentapi", func(c *gin.Context) {
		fmt.Println("Work start contentapi")

	})

	cin.GET("/grouproom", func(c *gin.Context) {
		fmt.Println("Work start getmemberroom")

	})

	cin.Run(":" + port)
}
