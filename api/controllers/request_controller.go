package controllers

import (
	"net/http"

	"github.com/sskmy1024/PartnerAssistant/modules/client/linebot"
	"github.com/sskmy1024/PartnerAssistant/modules/client/sample"

	"github.com/labstack/echo"
)

// errorMessage : エラーを返すための型
type errorMessage struct {
	Message string `json:"message"`
}

// RequestController : 送られてきたリクエストに対し、モジュールに振り分ける
func RequestController(c echo.Context) error {
	module := c.Param("module")

	switch module {
	// どうにかする（後で考える）
	case "linebot":
		return linebot.ExecuteProcess(c)
	case "sample":
		return sample.ExecuteProcess(c)
	default:
		return c.JSON(http.StatusBadRequest, &errorMessage{
			Message: "The requested client does not exist",
		})
	}
}
