package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sskmy1024/PartnerAssistant/api/controllers"
)

// Router : ルーティング
func Router(e *echo.Echo) {
	g := e.Group("/api")

	g.Use(middleware.CORS())

	// リクエストを受けるエンドポイント
	g.POST("/:module", controllers.RequestController)
}
