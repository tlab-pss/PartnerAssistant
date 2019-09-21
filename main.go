package main

import (
	"github.com/labstack/echo"

	reqManager "main/modules/requestmanager"
)

func main() {
	e := echo.New()

	e.GET("/:module", reqManager.Controller)

	e.Logger.Fatal(e.Start(":8080"))
}
