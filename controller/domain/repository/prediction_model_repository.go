package repository

import "github.com/xkurozaru/plant-diagnosis/controller/domain/model"

type PredictionModelRepository interface {
	Create(model.PredictionModel) error
	FindAll() ([]model.PredictionModel, error)
	Find(ID model.ULID) (model.PredictionModel, error)
}
