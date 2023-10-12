package router

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/xkurozaru/plant-diagnosis/controller/interface/handler"
)

func InitAccountRouter(e *echo.Echo, accountHandler handler.AccountHandler) {
	api := e.Group("/api/v1")

	api.POST("/sign-up", accountHandler.SignUp())
	api.POST("/sign-up-admin", accountHandler.SignUpAdmin())
	api.POST("/sign-in", accountHandler.SignIn())

	u := api.Group("/users")
	u.Use(echojwt.WithConfig(handler.JwtConfig))

	u.GET("", accountHandler.GetUser())
}
