# Prediction SQLAlchemy Model

from model.id_model import IdModel
from sqlalchemy import Column, ForeignKey, String, Text
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class Prediction(Base, IdModel):
    __tablename__ = "predictions"
    image_path = Column(Text, nullable=False)
    result = Column(Text, nullable=False)
    user_id = Column(String(26), ForeignKey("users.id"))
    prediction_model_id = Column(String(26), ForeignKey("prediction_models.id"))

    def __init__(
        self, image_path: str, result: str, user_id: str, prediction_model_id: str
    ):
        super().__init__()
        self.image_path = image_path
        self.result = result
        self.user_id = user_id
        self.prediction_model_id = prediction_model_id
