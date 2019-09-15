package main

import (
    "net/http"
    "github.com/labstack/echo"
)

type User struct {
    ID int `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    e := echo.New()
    e.GET("/", hello)
    e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
    user := &User{
        ID: 100,
        Name: "sample user",
        Email: "sample@test.com",
    }
    return c.JSON(http.StatusOK, user)
}
