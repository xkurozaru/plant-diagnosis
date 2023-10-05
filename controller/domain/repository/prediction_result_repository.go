package repository

import "github.com/xkurozaru/plant-diagnosis/controller/domain/model"

type PredictionResultRepository interface {
	Create(model.PredictionResult) error
	Find(ID model.ULID) (model.PredictionResult, error)
	FindByUserID(userID model.ULID) ([]model.PredictionResult, error)
}
