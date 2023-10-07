package router

import (
	"github.com/labstack/echo/v4"
	"github.com/xkurozaru/plant-diagnosis/controller/interface/handler"
)

func InitAccountRouter(e *echo.Echo, handler handler.AccountHandler) {
	api := e.Group("/api/v1")

	api.POST("/sign-up", handler.SignUp())
	api.POST("/sign-in", handler.SignIn())
}
