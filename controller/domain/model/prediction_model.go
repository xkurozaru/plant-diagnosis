package model

type PredictionModel struct {
	ID          ULID
	Name        string
	NetworkName string
	ParamPath   string
	Labels      PredictionLabels
}

func NewPredictionModel(name string, networkName string, paramPath string, labels []string) PredictionModel {
	ID := NewULID()
	return PredictionModel{
		ID:          ID,
		Name:        name,
		NetworkName: networkName,
		ParamPath:   paramPath,
		Labels:      NewPredictionLabels(labels, ID),
	}
}
