package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"

	reqManager "main/modules/requestmanager"
)

func main() {
	err := godotenv.Load(fmt.Sprintf("envfiles/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.POST("/:module", reqManager.Controller)

	e.Logger.Fatal(e.Start(":8080"))
}
