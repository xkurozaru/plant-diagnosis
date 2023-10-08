package database

import "github.com/xkurozaru/plant-diagnosis/controller/domain/model"

type PredictionLabelEntity struct {
	Model
	PredictionModelID string
	Name              string
	Index             uint
}

func NewPredictionLabelEntity(predictionLabel model.PredictionLabel) PredictionLabelEntity {
	return PredictionLabelEntity{
		Model:             Model{ID: predictionLabel.ID.ToString()},
		PredictionModelID: predictionLabel.PredictionModelID.ToString(),
		Name:              predictionLabel.Name,
		Index:             predictionLabel.Index,
	}
}

func (p PredictionLabelEntity) ToModel() model.PredictionLabel {
	return model.PredictionLabel{
		ID:                model.ULID(p.ID),
		PredictionModelID: model.ULID(p.PredictionModelID),
		Name:              p.Name,
		Index:             p.Index,
	}
}
