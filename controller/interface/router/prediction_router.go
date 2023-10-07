package router

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/xkurozaru/plant-diagnosis/controller/interface/handler"
)

func InitPredictionRouter(e *echo.Echo, predictionHandler handler.PredictionHandler) {
	api := e.Group("/api/v1")
	api.Use(echojwt.WithConfig(handler.JwtConfig))

	api.POST("/prediction/models", predictionHandler.CreatePredictionModel())
	api.GET("/prediction/models", predictionHandler.GetPredictionModels())
	api.GET("/prediction/models/:model_id", predictionHandler.GetPredictionModel())
	api.POST("/prediction/predict/:model_id", predictionHandler.Predict())
}
