package sample

import (
	"github.com/labstack/echo"

	"main/modules/logic"
)

// ParseRequest : Logicモジュールに向けて型を整形する
func ParseRequest(c echo.Context) string {
	// Logicモジュールに向けて型を整える
	return logic.RequestAI()
}
