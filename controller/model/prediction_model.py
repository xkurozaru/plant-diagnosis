# Prediction Model SQLAlchemy Model

import torch
from model.id_model import IdModel
from PIL import Image
from prediction.resnet import Resnet18
from prediction.utils import transform_image
from sqlalchemy import Column, String, Text
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import relationship

Base = declarative_base()


class PredictionModel(Base, IdModel):
    __tablename__ = "prediction_models"
    name = Column(String(256), nullable=False)
    weight_path = Column(Text, nullable=False)
    prediction_classes = relationship("PredictionClass", backref="prediction_model")

    def __init__(self, name: str, weight_path: str):
        super().__init__()
        self.name = name
        self.weight_path = weight_path

    def prediction(self, image: Image.Image):
        tensor = transform_image(image)

        model = Resnet18(num_classes=self.num_classes())
        model.load_state_dict(
            torch.load(self.weight_path, map_location=torch.device("cpu"))
        )
        model.eval()

        outputs = model.forward(tensor)
        prob, pred = outputs.max(1)
        predicted_idx = pred.item()
        probability = prob.item()
        del model, outputs, prob, pred

        return self.idx2class(predicted_idx), probability

    def idx2class(self, idx: int):
        for prediction_class in self.prediction_classes:
            if prediction_class.index == idx:
                return prediction_class.label
        return None

    def num_classes(self):
        return len(self.prediction_classes)
