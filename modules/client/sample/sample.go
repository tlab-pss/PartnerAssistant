package sample

import (
	"github.com/labstack/echo"

	"main/modules/logic"
)

// ParseRequest : Logicモジュールに向けて型を整形する
func ParseRequest(c echo.Context) interface{} {
	// Logicモジュールに向けて型を整える

	// AIの返答をもらう
	return logic.RequestAI()
}
