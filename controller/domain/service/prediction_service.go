package service

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/repository"
)

type PredictionService interface {
	ExecPrediction(user model.User, predictionModel model.PredictionModel, file multipart.FileHeader) (model.PredictionResult, error)
}

type predictionService struct {
	predictionResultRepository repository.PredictionResultRepository
}

func NewPredictionService(
	predictionResultRepository repository.PredictionResultRepository,
) PredictionService {
	return predictionService{
		predictionResultRepository,
	}
}

func (p predictionService) ExecPrediction(user model.User, predictionModel model.PredictionModel, file multipart.FileHeader) (model.PredictionResult, error) {
	img, err := file.Open()
	defer img.Close()
	if err != nil {
		return model.PredictionResult{}, err
	}

	// 予測を実行する
	client := resty.New()
	resp, err := client.R().
		SetFileReader("image", file.Filename, img).
		SetFormData(map[string]string{
			"net_name":   predictionModel.NetworkName,
			"param_path": predictionModel.ParamPath,
		}).
		SetFormDataFromValues(url.Values{
			"labels": predictionModel.Labels.ToSlice(),
		}).
		Post("http://predictor:5000/predictor/predict")

	if err != nil {
		return model.PredictionResult{}, err
	}

	// respは {"class":"健常"} のような形式で返ってくる
	// "健常"の部分を抜き出す
	data := map[string]string{}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return model.PredictionResult{}, err
	}
	res := data["class"]
	if res == "" {
		return model.PredictionResult{}, fmt.Errorf("failed to prediction")
	}

	// imgを保存する
	// 保存先は /app/_var/{user_id}/{DateTime}_{file_name}.jpg
	saveDir := fmt.Sprintf("/app/_var/%s", user.ID)
	err = os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		return model.PredictionResult{}, err
	}
	newFilePath := fmt.Sprintf("%s/%s_%s", saveDir, model.DateTimeNow().Format(model.DateTimeBarFormat), file.Filename)
	newFile, err := os.Create(newFilePath)
	if err != nil {
		return model.PredictionResult{}, err
	}
	defer newFile.Close()
	_, err = img.Seek(0, io.SeekStart)
	if err != nil {
		return model.PredictionResult{}, err
	}
	_, err = io.Copy(newFile, img)
	if err != nil {
		return model.PredictionResult{}, err
	}

	// 予測結果を保存する
	predictionResult := model.NewPredictionResult(user, predictionModel, res, newFilePath)
	err = p.predictionResultRepository.Create(predictionResult)
	if err != nil {
		return model.PredictionResult{}, err
	}

	return predictionResult, err
}
