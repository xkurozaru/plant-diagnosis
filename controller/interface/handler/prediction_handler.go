package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xkurozaru/plant-diagnosis/controller/application"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
	"github.com/xkurozaru/plant-diagnosis/controller/interface/messages"
)

type PredictionHandler interface {
	CreatePredictionModel() echo.HandlerFunc
	GetPredictionModels() echo.HandlerFunc
	GetPredictionModel() echo.HandlerFunc
	Predict() echo.HandlerFunc
	DeletePredictionModel() echo.HandlerFunc
}

type predictionHandler struct {
	predictionApplicationService application.PredictionApplicationService
}

func NewPredictionHandler(
	predictionApplicationService application.PredictionApplicationService,
) PredictionHandler {
	return predictionHandler{
		predictionApplicationService,
	}
}

func (p predictionHandler) CreatePredictionModel() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req messages.CreatePredictionModelRequest
		var res messages.CreatePredictionModelResponse

		userID, err := GetUserID(ctx)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		err = ctx.Bind(&req)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
		}

		err = p.predictionApplicationService.CreatePredictionModel(userID, req.ModelName, req.NetworkName, req.ParamPath, req.Labels)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		return ctx.JSON(http.StatusOK, res)
	}
}

func (p predictionHandler) GetPredictionModels() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var res messages.GetPredictionModelsResponse

		userID, err := GetUserID(ctx)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		predictionModels, err := p.predictionApplicationService.GetPredictionModels(userID)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		for _, predictionModel := range predictionModels {
			res.PredictionModels = append(res.PredictionModels, messages.NewPredictionModel(predictionModel))
		}

		return ctx.JSON(http.StatusOK, res)
	}
}

func (p predictionHandler) GetPredictionModel() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req messages.GetPredictionModelRequest
		var res messages.GetPredictionModelResponse

		userID, err := GetUserID(ctx)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		err = ctx.Bind(&req)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
		}

		predictionModel, err := p.predictionApplicationService.GetPredictionModel(userID, model.ULID(req.ModelID))
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		res.PredictionModel = messages.NewPredictionModel(predictionModel)

		return ctx.JSON(http.StatusOK, res)
	}
}

func (p predictionHandler) Predict() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req messages.PredictRequest
		var res messages.PredictResponse

		userID, err := GetUserID(ctx)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		err = ctx.Bind(&req)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
		}

		req.File, err = ctx.FormFile("image")
		if err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
		}

		predictionResult, err := p.predictionApplicationService.Predict(userID, model.ULID(req.ModelID), *req.File)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		res.PredictionResult = messages.NewPredictionResult(predictionResult)

		return ctx.JSON(http.StatusOK, res)
	}
}

func (p predictionHandler) DeletePredictionModel() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req messages.DeletePredictionModelRequest
		var res messages.DeletePredictionModelResponse

		userID, err := GetUserID(ctx)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		err = ctx.Bind(&req)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
		}

		err = p.predictionApplicationService.DeletePredictionModel(userID, model.ULID(req.ModelID))
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		return ctx.JSON(http.StatusOK, res)
	}
}
