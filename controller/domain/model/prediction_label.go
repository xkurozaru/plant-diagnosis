package model

type PredictionLabel struct {
	ID                ULID
	PredictionModelID ULID
	Name              string
	Index             uint
}

type PredictionLabels []PredictionLabel

func NewPredictionLabels(labels []string, predictionModelID ULID) PredictionLabels {
	predictionLabels := make(PredictionLabels, len(labels))
	for i, label := range labels {
		predictionLabels[i] = PredictionLabel{
			ID:                NewULID(),
			PredictionModelID: predictionModelID,
			Name:              label,
			Index:             uint(i),
		}
	}
	return predictionLabels
}

func (p PredictionLabels) NameOf(index uint) string {
	for _, label := range p {
		if label.Index == index {
			return label.Name
		}
	}
	return ""
}

func (p PredictionLabels) ToSlice() []string {
	labels := make([]string, len(p))
	for _, label := range p {
		labels[label.Index] = label.Name
	}
	return labels
}
