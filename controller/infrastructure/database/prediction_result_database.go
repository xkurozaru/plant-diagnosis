package database

import (
	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type predictionResultDatabase struct {
	db *gorm.DB
}

func NewPredictionResultDatabase(db *gorm.DB) repository.PredictionResultRepository {
	return &predictionResultDatabase{db: db}
}

func (p predictionResultDatabase) Create(predictionResult model.PredictionResult) error {
	predictionResultE := NewPredictionResultEntity(predictionResult)

	err := p.db.Create(&predictionResultE).Error
	if err != nil {
		return err
	}

	return nil
}

func (p predictionResultDatabase) FindByUserID(userID model.ULID) ([]model.PredictionResult, error) {
	predictionResultEs := []PredictionResultEntity{}
	err := p.preloads().Where("user_id = ?", userID.ToString()).Find(&predictionResultEs).Error
	if err != nil {
		return nil, err
	}

	predictionResults := []model.PredictionResult{}
	for _, predictionResultE := range predictionResultEs {
		predictionResults = append(predictionResults, predictionResultE.ToModel())
	}

	return predictionResults, nil
}

func (p predictionResultDatabase) Find(ID model.ULID) (model.PredictionResult, error) {
	predictionResultE := PredictionResultEntity{}
	err := p.preloads().First(&predictionResultE, ID.ToString()).Error
	if err != nil {
		return model.PredictionResult{}, err
	}

	return predictionResultE.ToModel(), nil
}

func (p predictionResultDatabase) preloads() *gorm.DB {
	return p.db.Preload(clause.Associations).Preload("PredictionModel.PredictionLabels")
}
