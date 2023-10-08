package database

import "github.com/xkurozaru/plant-diagnosis/controller/domain/model"

type PredictionResultEntity struct {
	Model
	UserID            string
	User              UserEntity `gorm:"foreignKey:UserID"`
	PredictionModelID string
	PredictionModel   PredictionModelEntity `gorm:"foreignKey:PredictionModelID"`
	Result            string
	FilePath          string
}

func NewPredictionResultEntity(predictionResult model.PredictionResult) PredictionResultEntity {
	return PredictionResultEntity{
		Model:             Model{ID: predictionResult.ID.ToString(), CreatedAt: predictionResult.PredictedAt.ToTime()},
		UserID:            predictionResult.PredictedBy.ID.ToString(),
		PredictionModelID: predictionResult.PredictionModel.ID.ToString(),
		Result:            predictionResult.Result,
		FilePath:          predictionResult.FilePath,
	}
}

func (p PredictionResultEntity) ToModel() model.PredictionResult {
	return model.PredictionResult{
		ID:              model.ULID(p.ID),
		PredictedAt:     model.DateTime(p.CreatedAt),
		PredictedBy:     p.User.ToModel(),
		PredictionModel: p.PredictionModel.ToModel(),
		Result:          p.Result,
		FilePath:        p.FilePath,
	}
}
