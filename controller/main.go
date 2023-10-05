package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xkurozaru/plant-diagnosis/controller/infrastructure"
	"github.com/xkurozaru/plant-diagnosis/controller/interface/router"
	"github.com/xkurozaru/plant-diagnosis/controller/registry"
)

func main() {
	infrastructure.InitDB()
	reg := registry.NewRegistry(infrastructure.GetDB())

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	router.InitAccountRouter(e, reg.NewAccountHandler())

	e.Logger.Fatal(e.Start(":8000"))
}
