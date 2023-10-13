package model

import "fmt"

type PredictionModel struct {
	ID          ULID
	Name        string
	NetworkName string
	ParamPath   string
	Labels      PredictionLabels
}

func NewPredictionModel(name string, networkName string, paramPath string, labels []string) (PredictionModel, error) {
	ID := NewULID()
	p := PredictionModel{
		ID:          ID,
		Name:        name,
		NetworkName: networkName,
		ParamPath:   paramPath,
		Labels:      NewPredictionLabels(labels, ID),
	}

	err := p.Validate()
	if err != nil {
		return PredictionModel{}, err
	}

	return p, nil
}

func (p PredictionModel) Validate() error {
	if len(p.Name) <= 0 {
		return fmt.Errorf("name must not be empty")
	}

	if len(p.NetworkName) <= 0 {
		return fmt.Errorf("network name must not be empty")
	}

	if len(p.ParamPath) <= 0 {
		return fmt.Errorf("param path must not be empty")
	}

	if len(p.Labels) <= 0 {
		return fmt.Errorf("labels must not be empty")
	}

	return nil
}
