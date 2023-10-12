package application

import (
	"fmt"
	"mime/multipart"

	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/repository"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/service"
)

type PredictionApplicationService interface {
	CreatePredictionModel(userID model.ULID, modelName string, networkName string, paramPath string, labels []string) error
	GetPredictionModels(userID model.ULID) ([]model.PredictionModel, error)
	GetPredictionModel(userID model.ULID, modelID model.ULID) (model.PredictionModel, error)
	Predict(userID model.ULID, modelID model.ULID, file multipart.FileHeader) (model.PredictionResult, error)
	GetPredictionResults(userID model.ULID) ([]model.PredictionResult, error)
	DeletePredictionModel(userID model.ULID, modelID model.ULID) error
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

func (p predictionApplicationService) CreatePredictionModel(userID model.ULID, modelName string, networkName string, paramPath string, labels []string) error {
	user, err := p.userRepository.Find(userID)
	if err != nil {
		return err
	}

	if !user.Role.HasPermission(model.CreatePredictionModelPermission) {
		return fmt.Errorf("Permission denied to create prediction model")
	}

	predictionModel := model.NewPredictionModel(modelName, networkName, paramPath, labels)
	err = p.predictionModelRepository.Create(predictionModel)
	if err != nil {
		return err
	}

	return nil
}

func (p predictionApplicationService) GetPredictionModels(userID model.ULID) ([]model.PredictionModel, error) {
	user, err := p.userRepository.Find(userID)
	if err != nil {
		return nil, err
	}

	if !user.Role.HasPermission(model.ReadPredictionModelPermission) {
		return nil, fmt.Errorf("Permission denied to read prediction models")
	}

	predictionModels, err := p.predictionModelRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return predictionModels, nil
}

func (p predictionApplicationService) GetPredictionModel(userID model.ULID, modelID model.ULID) (model.PredictionModel, error) {
	user, err := p.userRepository.Find(userID)
	if err != nil {
		return model.PredictionModel{}, err
	}

	if !user.Role.HasPermission(model.ReadPredictionModelPermission) {
		return model.PredictionModel{}, fmt.Errorf("Permission denied to read prediction model")
	}

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

	if !user.Role.HasPermission(model.PredictionPermission) {
		return model.PredictionResult{}, fmt.Errorf("Permission denied to predict")
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

func (p predictionApplicationService) DeletePredictionModel(userID model.ULID, modelID model.ULID) error {
	user, err := p.userRepository.Find(userID)
	if err != nil {
		return err
	}

	if !user.Role.HasPermission(model.DeletePredictionModelPermission) {
		return fmt.Errorf("Permission denied to delete prediction model")
	}

	predictionModel, err := p.predictionModelRepository.Find(modelID)
	if err != nil {
		return err
	}

	err = p.predictionModelRepository.Delete(predictionModel)
	if err != nil {
		return err
	}

	return nil
}
