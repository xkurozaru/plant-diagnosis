package router

import (
	"github.com/labstack/echo/v4"
	"github.com/xkurozaru/plant-diagnosis/controller/interface/handler"
)

func InitAccountRouter(e *echo.Echo, handler handler.AccountHandler) {
	e.POST("/account/sign-up", handler.SignUp())
	e.POST("/account/sign-in", handler.SignIn())
}
