package sample

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"main/modules/logic"
)

// RequestType : 送られてくるリクエストの型
type RequestType struct {
	Message string `json:"message"`
}

// ParseRequest : Logicモジュールに向けて型を整形する
func ParseRequest(c echo.Context) error {
	// Logicモジュールに向けて型を整える
	post := new(RequestType)

	if err := c.Bind(post); err != nil {
		fmt.Println("POST data :", err)
        return err
    }

	// AIの返答をもらう
	return c.JSON(http.StatusOK, logic.RequestAI(post.Message))
}
