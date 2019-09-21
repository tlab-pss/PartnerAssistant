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

	ab := os.Getenv("TEST")
	log.Println(ab)

	e := echo.New()

	e.GET("/:module", reqManager.Controller)

	e.Logger.Fatal(e.Start(":8080"))
}
