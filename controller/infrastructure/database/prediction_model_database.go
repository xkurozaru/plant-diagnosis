package database

import (
	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type predictionModelDatabase struct {
	db *gorm.DB
}

func NewPredictionModelDatabase(db *gorm.DB) repository.PredictionModelRepository {
	return &predictionModelDatabase{db: db}
}

func (p predictionModelDatabase) Create(predictionModel model.PredictionModel) error {
	predictionModelE := NewPredictionModelEntity(predictionModel)

	err := p.db.Create(&predictionModelE).Error
	if err != nil {
		return err
	}

	return nil
}

func (p predictionModelDatabase) FindAll() ([]model.PredictionModel, error) {
	predictionModelEs := []PredictionModelEntity{}
	err := p.db.Preload(clause.Associations).Find(&predictionModelEs).Error
	if err != nil {
		return nil, err
	}

	predictionModels := []model.PredictionModel{}
	for _, predictionModelE := range predictionModelEs {
		predictionModels = append(predictionModels, predictionModelE.ToModel())
	}

	return predictionModels, nil

}

func (p predictionModelDatabase) Find(ID model.ULID) (model.PredictionModel, error) {
	predictionModelE := PredictionModelEntity{}
	err := p.db.Preload(clause.Associations).Where("id = ?", ID.ToString()).First(&predictionModelE).Error
	if err != nil {
		return model.PredictionModel{}, err
	}

	return predictionModelE.ToModel(), nil
}
