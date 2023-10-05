# Prediction Class SQLAlchemy Model

from model.id_model import IdModel
from sqlalchemy import Column, ForeignKey, Integer, String
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class PredictionClass(Base, IdModel):
    __tablename__ = "prediction_classes"
    index = Column(Integer, nullable=False)
    label = Column(String(256), nullable=False)
    prediction_model_id = Column(String(26), ForeignKey("prediction_models.id"))

    def __init__(self, index: int, label: str, prediction_model_id: str):
        super().__init__()
        self.index = index
        self.label = label
        self.prediction_model_id = prediction_model_id


def new_prediction_classes(prediction_model_id, classes):
    return [PredictionClass(i, c, prediction_model_id) for i, c in enumerate(classes)]
