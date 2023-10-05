package messages

import (
	"mime/multipart"

	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
)

////////////////////////
// Request & Response //
////////////////////////

type CreatePredictionModelRequest struct {
	ModelName   string   `json:"model_name"`
	NetworkName string   `json:"network_name"`
	ParamPath   string   `json:"param_path"`
	Labels      []string `json:"labels"`
}
type CreatePredictionModelResponse struct{}

type GetPredictionModelsRequest struct{}
type GetPredictionModelsResponse struct {
	PredictionModels []PredictionModel `json:"prediction_models"`
}

type GetPredictionModelRequest struct {
	ModelID string `param:"model_id"`
}
type GetPredictionModelResponse struct {
	PredictionModel PredictionModel `json:"prediction_model"`
}

type PredictRequest struct {
	ModelID string                `param:"model_id"`
	Image   *multipart.FileHeader `form:"image"`
}
type PredictResponse struct {
	PredictionResult PredictionResult `json:"prediction_result"`
}

type GetPredictionResultsRequest struct{}
type GetPredictionResultsResponse struct {
	PredictionResults []PredictionResult `json:"prediction_results"`
}

///////////
// Model //
///////////

type PredictionModel struct {
	ID        string   `json:"id"`
	ModelName string   `json:"model_name"`
	Labels    []string `json:"labels"`
}

func NewPredictionModel(p model.PredictionModel) PredictionModel {
	return PredictionModel{
		ID:        p.ID.ToString(),
		ModelName: p.Name,
		Labels:    p.Labels.ToSlice(),
	}
}

type PredictionResult struct {
	ID          string `json:"id"`
	PredictedAt string `json:"predicted_at"`
	Result      string `json:"result"`
}

func NewPredictionResult(p model.PredictionResult) PredictionResult {
	return PredictionResult{
		ID:          p.ID.ToString(),
		PredictedAt: p.PredictedAt.Format(model.DateTimeFormat),
		Result:      p.Result,
	}
}
