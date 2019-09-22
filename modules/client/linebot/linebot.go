// Package linebot : LINE Botを使用する際のModulePlugin
package linebot

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

// ParseRequest : Logicモジュールに向けて型を整形する
func ParseRequest(c echo.Context) error {

	bot, err := linebot.New(
		os.Getenv("linebot_channel_secret"),
		os.Getenv("linebot_channel_token"),
	)
	if err != nil {
		fmt.Println("Initialize Error:", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	received, err := bot.ParseRequest(c.Request())
	if err != nil {
		fmt.Println("Receive Error:", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	for _, event := range received {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				fmt.Println(message)
				// 送ったメッセージをおうむ返ししているだけ
				resMessage := linebot.NewTextMessage(message.Text)
				if _, err = bot.ReplyMessage(event.ReplyToken, resMessage).Do(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	return c.JSON(http.StatusOK, "success")
}
