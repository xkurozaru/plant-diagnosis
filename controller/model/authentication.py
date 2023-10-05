# authentication SQLAlchemy model

from datetime import datetime, timedelta
from hashlib import sha256

from jose import jwt
from model.id_model import IdModel
from sqlalchemy import Column, ForeignKey, String
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class Authentication(Base, IdModel):
    __tablename__ = "authentications"
    hash = Column(String(256), nullable=False)
    user_id = Column(String(26), ForeignKey("users.id"))

    def __init__(self, user_id: str, password: str):
        super().__init__()
        self.user_id = user_id
        self.password = sha256(password.encode("utf-8")).hexdigest()

    def verify_password(self, password: str) -> bool:
        return self.password == sha256(password.encode("utf-8")).hexdigest()

    def generate_token(self) -> str:
        expire = datetime.utcnow() + timedelta(minutes=30)
        data = {"user_id": self.user_id, "exp": expire}
        encoded_jwt = jwt.encode(data)
        return encoded_jwt
