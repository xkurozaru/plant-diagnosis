package database

import "github.com/xkurozaru/plant-diagnosis/controller/domain/model"

type PredictionModelEntity struct {
	Model
	Name             string
	NetworkName      string
	ParamPath        string
	PredictionLabels []PredictionLabelEntity `gorm:"foreignKey:PredictionModelID"`
}

func NewPredictionModelEntity(predictionModel model.PredictionModel) PredictionModelEntity {
	labelEs := []PredictionLabelEntity{}
	for _, label := range predictionModel.Labels {
		labelEs = append(labelEs, NewPredictionLabelEntity(label))
	}

	return PredictionModelEntity{
		Model:            Model{ID: predictionModel.ID.ToString()},
		Name:             predictionModel.Name,
		NetworkName:      predictionModel.NetworkName,
		ParamPath:        predictionModel.ParamPath,
		PredictionLabels: labelEs,
	}
}

func (p PredictionModelEntity) ToModel() model.PredictionModel {
	labelModels := []model.PredictionLabel{}
	for _, labelE := range p.PredictionLabels {
		labelModels = append(labelModels, labelE.ToModel())
	}

	return model.PredictionModel{
		ID:          model.ULID(p.ID),
		Name:        p.Name,
		NetworkName: p.NetworkName,
		Labels:      labelModels,
	}
}
