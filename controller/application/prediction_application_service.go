package application

import (
	"mime/multipart"

	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/repository"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/service"
)

type PredictionApplicationService interface {
	CreatePredictionModel(modelName string, networkName string, paramPath string, labels []string) error
	GetPredictionModels() ([]model.PredictionModel, error)
	GetPredictionModel(modelID model.ULID) (model.PredictionModel, error)
	Predict(userID model.ULID, modelID model.ULID, file multipart.FileHeader) (model.PredictionResult, error)
	GetPredictionResults(userID model.ULID) ([]model.PredictionResult, error)
}

type predictionApplicationService struct {
	predictionModelRepository  repository.PredictionModelRepository
	predictionResultRepository repository.PredictionResultRepository
	userRepository             repository.UserRepository
	predictionService          service.PredictionService
}

func NewPredictionApplicationService(
	predictionModelRepository repository.PredictionModelRepository,
	predictionResultRepository repository.PredictionResultRepository,
	userRepository repository.UserRepository,
	predictionService service.PredictionService,
) PredictionApplicationService {
	return predictionApplicationService{
		predictionModelRepository:  predictionModelRepository,
		predictionResultRepository: predictionResultRepository,
		userRepository:             userRepository,
		predictionService:          predictionService,
	}
}

func (p predictionApplicationService) CreatePredictionModel(modelName string, networkName string, paramPath string, labels []string) error {
	predictionModel := model.NewPredictionModel(modelName, networkName, paramPath, labels)
	err := p.predictionModelRepository.Create(predictionModel)
	if err != nil {
		return err
	}

	return nil
}

func (p predictionApplicationService) GetPredictionModels() ([]model.PredictionModel, error) {
	predictionModels, err := p.predictionModelRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return predictionModels, nil
}

func (p predictionApplicationService) GetPredictionModel(modelID model.ULID) (model.PredictionModel, error) {
	predictionModel, err := p.predictionModelRepository.Find(modelID)
	if err != nil {
		return model.PredictionModel{}, err
	}

	return predictionModel, nil
}

func (p predictionApplicationService) Predict(userID model.ULID, modelID model.ULID, file multipart.FileHeader) (model.PredictionResult, error) {
	user, err := p.userRepository.Find(userID)
	if err != nil {
		return model.PredictionResult{}, err
	}

	predictionModel, err := p.predictionModelRepository.Find(modelID)
	if err != nil {
		return model.PredictionResult{}, err
	}

	predictionResult, err := p.predictionService.ExecPrediction(user, predictionModel, file)
	if err != nil {
		return model.PredictionResult{}, err
	}

	return predictionResult, nil
}

func (p predictionApplicationService) GetPredictionResults(userID model.ULID) ([]model.PredictionResult, error) {
	user, err := p.userRepository.Find(userID)
	if err != nil {
		return nil, err
	}

	predictionResults, err := p.predictionResultRepository.FindByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	return predictionResults, nil
}
