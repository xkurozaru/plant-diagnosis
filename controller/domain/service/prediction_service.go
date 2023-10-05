package service

import (
	"fmt"
	"mime/multipart"

	"github.com/go-resty/resty/v2"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/repository"
)

type PredictionService interface {
	ExecPrediction(user model.User, predictionModel model.PredictionModel, file multipart.FileHeader) (model.PredictionResult, error)
}

type predictionService struct {
	predictionResultRepository repository.PredictionResultRepository
}

func NewPredictionService(
	predictionResultRepository repository.PredictionResultRepository,
) PredictionService {
	return predictionService{
		predictionResultRepository,
	}
}

func (p predictionService) ExecPrediction(user model.User, predictionModel model.PredictionModel, file multipart.FileHeader) (model.PredictionResult, error) {
	img, err := file.Open()
	defer img.Close()
	if err != nil {
		return model.PredictionResult{}, err
	}

	client := resty.New()
	resp, err := client.R().
		SetFileReader("image", file.Filename, img).
		Get(fmt.Sprintf("http://localhost:5000/predictor/predict/%s", predictionModel.Name))

	if err != nil {
		return model.PredictionResult{}, err
	}

	return model.NewPredictionResult(user, predictionModel, resp.String(), file.Filename), nil
}
