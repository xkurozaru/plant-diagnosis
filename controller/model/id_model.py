# id_model SQLAlchemy class

from cuid2 import Cuid
from sqlalchemy import Column, DateTime, String
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.sql import func

Base = declarative_base()


class IdModel(Base):
    __abstract__ = True

    id = Column(String(26), primary_key=True)
    created_at = Column(DateTime(timezone=True), server_default=func.now())
    updated_at = Column(DateTime(timezone=True), onupdate=func.now())
    deleted_at = Column(DateTime(timezone=True), nullable=True)

    def __init__(self):
        self.id = Cuid().generate(26)
        self.created_at = func.now()
        self.updated_at = func.now()
        self.deleted_at = None
