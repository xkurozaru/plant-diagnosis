package model

type PredictionResult struct {
	ID              ULID
	PredictedAt     DateTime
	PredictedBy     User
	PredictionModel PredictionModel
	Result          string
	FileName        string
}

func NewPredictionResult(predictedBy User, predictionModel PredictionModel, result string, fileName string) PredictionResult {
	return PredictionResult{
		ID:              NewULID(),
		PredictedAt:     DateTimeNow(),
		PredictedBy:     predictedBy,
		PredictionModel: predictionModel,
		Result:          result,
		FileName:        fileName,
	}
}
