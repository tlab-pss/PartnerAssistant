// Package linebot : LINE Botを使用する際のModulePlugin
package linebot

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/sskmy1024/PartnerAssistant/modules/logic"
)

// ExecuteProcess : Logicモジュールにリクエストを送り、処理結果をクライアントに返す
func ExecuteProcess(c echo.Context) error {

	bot, err := linebot.New(
		os.Getenv("LINEBOT_CHANNEL_SECRET"),
		os.Getenv("LINEBOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		fmt.Println("Initialize Error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	received, err := bot.ParseRequest(c.Request())
	if err != nil {
		fmt.Println("Receive Error:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	for _, event := range received {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				fmt.Println("Get message:", message.Text)

				payload := &logic.LogicPayload{
					UserMessage: message.Text,
				}
				watsonResponse, err := payload.ExecuteLogic()
				if err != nil {
					fmt.Println("Watson error:", err)
				}
				resMessage := linebot.NewTextMessage(watsonResponse.Message)
				if _, err = bot.ReplyMessage(event.ReplyToken, resMessage).Do(); err != nil {
					fmt.Println("Reply error:", err)
				}
			case *linebot.ImageMessage:
				payload := &logic.LogicPayload{
					ImagePath: message.OriginalContentURL,
				}
				response, err := payload.ExecuteLogic()
				if err != nil {
					fmt.Println("Watson error:", err)
				}
				resMessage := linebot.NewTextMessage(response.Message)
				if _, err = bot.ReplyMessage(event.ReplyToken, resMessage).Do(); err != nil {
					fmt.Println("Reply error:", err)
				}

				// imgResponse, err := http.Get(message.OriginalContentURL)
				// if err != nil {
				// 	log.Fatal(err)
				// }
				// defer imgResponse.Body.Close()

				// fileData, _ := ioutil.ReadAll(imgResponse.Body)
				// base64img := base64.StdEncoding.EncodeToString(fileData)
				// fmt.Printf("<img src=\"data:image/jpg;base64,%s\">", base64img)
			}
		}
	}

	return c.JSON(http.StatusOK, "success")
}
