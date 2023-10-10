package model

const (
	CreatePredictionModelPermission Permission = "CreatePredictionModel"
	ReadPredictionModelPermission   Permission = "ReadPredictionModel"
	DeletePredictionModelPermission Permission = "DeletePredictionModel"
	PredictionPermission            Permission = "Prediction"
)

type Permission string
