package sample

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/sskmy1024/PartnerAssistant/modules/logic"
)

// RequestType : 送られてくるリクエストの型
type RequestType struct {
	Message string `json:"message"`
}

// ExecuteProcess :  Logicモジュールにリクエストを送り、処理結果をクライアントに返す
func ExecuteProcess(c echo.Context) error {
	// Logicモジュールに向けて型を整える
	post := new(RequestType)

	if err := c.Bind(post); err != nil {
		fmt.Println("POST data :", err)
		return err
	}

	payload := &logic.LogicPayload{
		UserMessage: post.Message,
	}

	response, err := payload.ExecuteLogic()
	if err != nil {
		return err
	}

	// AIの返答をもらう
	return c.JSON(http.StatusOK, response)
}
