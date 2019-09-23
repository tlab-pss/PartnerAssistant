package requestmanager

import (
	"net/http"

	"github.com/labstack/echo"

	"main/modules/client/sample"
)

// errorMessage : エラーを返すための型
type errorMessage struct {
	Message string `json:"message"`
}

// Controller : 送られてきたリクエストに対し、モジュールに振り分ける
func Controller(c echo.Context) error {
	module := c.Param("module")

	switch module {
	// どうにかする（後で考える）
	case "sample":
		return sample.ParseRequest(c)
	default:
		return c.JSON(http.StatusBadRequest, &errorMessage{
			Message: "The requested client does not exist",
		})
	}
}
