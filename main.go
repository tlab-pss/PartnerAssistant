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

type Response struct {
    Message string `json:"message"`
}

type LinebotResponse struct {
    Token string `json:"token"`
    Message string `json:"message"`
}

func main() {
    e := echo.New()
    
    e.GET("/", hello)
    e.GET("/:module", moduleController)

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

func moduleController(c echo.Context) error {
    module := c.Param("module")
    switch module {
    case "linebot": return linebotHandler(c)
    // どうにかする（後で考える）
    default: return handler(c)
    }
} 

func handler(c echo.Context) error {
    res := &Response{
        Message: c.QueryParam("message"),
    }
    return c.JSON(http.StatusOK, res)
}

func linebotHandler(c echo.Context) error {
    res := &LinebotResponse{
        Message: c.QueryParam("message"),
        Token: c.Param("module"),
    }
    return c.JSON(http.StatusOK, res)
}
