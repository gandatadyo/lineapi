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
	// var urlAddress = "https://dark-asylum-97180.herokuapp.com"

	// Content
	contentsBubbleTas1 := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			Type:   linebot.FlexComponentTypeImage,
			URL:    "https://www.static-src.com/wcsstore/Indraprastha/images/catalog/medium//92/MTA-2070809/webe_webe-tote-bag-ivy-ribbon-red_full03.jpg",
			Size:   linebot.FlexImageSizeTypeFull,
			Action: &linebot.URIAction{"Label", "https://www.static-src.com/wcsstore/Indraprastha/images/catalog/medium//92/MTA-2070809/webe_webe-tote-bag-ivy-ribbon-red_full03.jpg"},
		},
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "Tas Red Velvet / Rp. 350.000,",
				},
			},
		},
		Footer: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.MessageAction{"Beli", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.MessageAction{"Stock", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
			},
		},
	}

	contentsBubbleTas2 := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			Type:   linebot.FlexComponentTypeImage,
			URL:    "https://www.indonesiaharga.biz.id/details/1050/images/82/MTA-1602234/catenzo_catenzo-kh-014-mareuli-dan-marema-tas-perempuan---blue_full02.jpg",
			Size:   linebot.FlexImageSizeTypeFull,
			Action: &linebot.URIAction{"Label", "https://www.indonesiaharga.biz.id/details/1050/images/82/MTA-1602234/catenzo_catenzo-kh-014-mareuli-dan-marema-tas-perempuan---blue_full02.jpg"},
		},
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "Tas Blue Jeans / Rp. 378.000,",
				},
			},
		},
		Footer: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Beli", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Stock", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
			},
		},
	}

	contentsBubbleTas3 := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			Type:   linebot.FlexComponentTypeImage,
			URL:    "https://dynamic.zacdn.com/-C4Br6Rh3qjREbBw5pFBImqItXU=/fit-in/472x690/filters:quality(90):fill(ffffff)/http://static.id.zalora.net/p/elizabeth-bags-0764-8220171-2.jpg",
			Size:   linebot.FlexImageSizeTypeFull,
			Action: &linebot.URIAction{"Label", "https://dynamic.zacdn.com/-C4Br6Rh3qjREbBw5pFBImqItXU=/fit-in/472x690/filters:quality(90):fill(ffffff)/http://static.id.zalora.net/p/elizabeth-bags-0764-8220171-2.jpg"},
		},
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "Tas Ping Crown / Rp. 428.000,",
				},
			},
		},
		Footer: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Beli", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Stock", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
			},
		},
	}

	contentsBubbleTas4 := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			Type:   linebot.FlexComponentTypeImage,
			URL:    "https://www.penjualan.co/details/700/images/103/MTA-1726888/lansdeal_4pcs-women-pattern-leather-shoulder-bag-crossbody-bag-handbag-wallet-gray-_full04.jpg",
			Size:   linebot.FlexImageSizeTypeFull,
			Action: &linebot.URIAction{"Label", "https://www.penjualan.co/details/700/images/103/MTA-1726888/lansdeal_4pcs-women-pattern-leather-shoulder-bag-crossbody-bag-handbag-wallet-gray-_full04.jpg"},
		},
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "Tas Casual Gray / Rp. 330.000,",
				},
			},
		},
		Footer: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Beli", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Stock", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
			},
		},
	}

	contentsCarouseTas := &linebot.CarouselContainer{
		Type: linebot.FlexContainerTypeCarousel,
		Contents: []*linebot.BubbleContainer{
			contentsBubbleTas1,
			contentsBubbleTas3,
			contentsBubbleTas4,
			contentsBubbleTas2,
		},
	}

	contentsBubbleSepatu1 := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			Type:   linebot.FlexComponentTypeImage,
			URL:    "https://s4.bukalapak.com/img/4700393723/large/Produk_Sepatu_Basket_Anta_Kt3_Low.jpg",
			Size:   linebot.FlexImageSizeTypeFull,
			Action: &linebot.URIAction{"Label", "https://s4.bukalapak.com/img/4700393723/large/Produk_Sepatu_Basket_Anta_Kt3_Low.jpg"},
		},
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "Sepatu Blue Nike / Rp. 320.000,",
				},
			},
		},
		Footer: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Beli", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Stock", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
			},
		},
	}

	contentsBubbleSepatu2 := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			Type:   linebot.FlexComponentTypeImage,
			URL:    "https://cdn-images-1.medium.com/max/600/1*8PUwBv4Ijc9YCWqO13nu9g.jpeg",
			Size:   linebot.FlexImageSizeTypeFull,
			Action: &linebot.URIAction{"Label", "https://cdn-images-1.medium.com/max/600/1*8PUwBv4Ijc9YCWqO13nu9g.jpeg"},
		},
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "Sepatu Black Pearl/ Rp. 372.000,",
				},
			},
		},
		Footer: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Beli", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Stock", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
			},
		},
	}

	contentsCarouseSepatu := &linebot.CarouselContainer{
		Type: linebot.FlexContainerTypeCarousel,
		Contents: []*linebot.BubbleContainer{
			contentsBubbleSepatu1,
			contentsBubbleSepatu2,
		},
	}

	contentsBubblebajukaos1 := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			Type:   linebot.FlexComponentTypeImage,
			URL:    "https://www.lebih.co/cdn/700/cache/p/2018/04/28/19/kaos-distro-420-unisex-premium-quality_4555708_1_19398.jpg",
			Size:   linebot.FlexImageSizeTypeFull,
			Action: &linebot.URIAction{"Label", "https://www.lebih.co/cdn/700/cache/p/2018/04/28/19/kaos-distro-420-unisex-premium-quality_4555708_1_19398.jpg"},
		},
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "Black Izoo / Rp. 250.000,",
				},
			},
		},
		Footer: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Beli", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Stock", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
			},
		},
	}

	contentsBubblebajukaos2 := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			Type:   linebot.FlexComponentTypeImage,
			URL:    "https://apollo-singapore.akamaized.net/v1/files/hqna747vpkol1-ID/image;s=966x691;olx-st/_1_.jpg",
			Size:   linebot.FlexImageSizeTypeFull,
			Action: &linebot.URIAction{"Label", "https://apollo-singapore.akamaized.net/v1/files/hqna747vpkol1-ID/image;s=966x691;olx-st/_1_.jpg"},
		},
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "Sth Tyle 40 / Rp. 275.000,",
				},
			},
		},
		Footer: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Beli", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Stock", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
			},
		},
	}

	contentsBubblebajukaos3 := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Hero: &linebot.ImageComponent{
			Type:   linebot.FlexComponentTypeImage,
			URL:    "https://ecs7.tokopedia.net/img/cache/700/product-1/2016/12/14/603660/603660_73e9fa30-83d5-4a40-8c94-3a9b94871355.jpg",
			Size:   linebot.FlexImageSizeTypeFull,
			Action: &linebot.URIAction{"Label", "https://ecs7.tokopedia.net/img/cache/700/product-1/2016/12/14/603660/603660_73e9fa30-83d5-4a40-8c94-3a9b94871355.jpg"},
		},
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "420 VintBlack / Rp. 290.000,",
				},
			},
		},
		Footer: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Beli", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: &linebot.URIAction{"Stock", "https://dark-asylum-97180.herokuapp.com/flex1"},
				},
			},
		},
	}

	contentsCarouselBajuKaos := &linebot.CarouselContainer{
		Type: linebot.FlexContainerTypeCarousel,
		Contents: []*linebot.BubbleContainer{
			contentsBubblebajukaos1,
			contentsBubblebajukaos2,
			contentsBubblebajukaos3,
		},
	}

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
		if err != nil {
			fmt.Println("Error 1.1")
			log.Print(err)
		}
		for _, event := range events {
			// fmt.Println(event.Source.UserID)
			// fmt.Println(event.Source.GroupID)
			// fmt.Println(event.Source.RoomID)
			// fmt.Println(event.ReplyToken)

			if event.Type == linebot.EventTypeMessage {
				fmt.Println("Work EventTypeMessage")

				switch message := event.Message.(type) {
				case *linebot.TextMessage:

					// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
					// 	log.Print(err)
					// }

					switch datamessage := message.Text; datamessage {
					case "Produk Tas":
						{
							if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("Produk Tas", contentsCarouseTas)).Do(); err != nil {
								log.Print(err)
							}
						}
					case "Produk Sepatu":
						{
							if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("Produk Sepatu", contentsCarouseSepatu)).Do(); err != nil {
								log.Print(err)
							}
						}
					case "Produk Baju dan Kaos":
						{
							if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("Produk Sepatu", contentsCarouselBajuKaos)).Do(); err != nil {
								log.Print(err)
							}
							fmt.Println("Linux.")
						}
					case "Information":
						{
							if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("Example Bot Ecomerce")).Do(); err != nil {
								log.Print(err)
							}
						}
					default:
						fmt.Println("Tidak ada perintah")
					}

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
