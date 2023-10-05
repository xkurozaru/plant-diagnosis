# User model SQLAlchemy class


from model.id_model import IdModel
from sqlalchemy import Column, String
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class User(Base, IdModel):
    __tablename__ = "users"
    login_id = Column(String(64), nullable=False)
    name = Column(String(256), nullable=False)

    def __init__(self, login_id: str, name: str):
        super().__init__()
        self.login_id = login_id
        self.name = name
