import os

from fastapi import Depends, FastAPI, HTTPException, UploadFile
from fastapi.middleware.cors import CORSMiddleware
from fastapi.security import OAuth2PasswordBearer
from jose import JWTError, jwt
from model import authentication, prediction, prediction_model, user
from PIL import Image
from sqlalchemy import create_engine
from sqlalchemy.orm import scoped_session, sessionmaker

DB_USER = os.getenv("DB_USER")
PASSWORD = os.getenv("PASSWORD")
HOST = os.getenv("HOST")
DATABASE = os.getenv("DATABASE")
engine = create_engine(f"mysql+mysqlconnector://{DB_USER}:{PASSWORD}@{HOST}/{DATABASE}")
session = scoped_session(sessionmaker(autocommit=False, autoflush=False, bind=engine))
session1 = session()

app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="sign-in")


@app.get("/")
async def index():
    return {"message": "Hello World"}


@app.post("/api/sign-up/")
async def sign_up(login_id: str, name: str, password: str):
    if not login_id or not name or not password:
        raise HTTPException(status_code=400, detail="Invalid request")

    u = user.User(login_id, name)
    a = authentication.Authentication(login_id, password)

    try:
        session1.add(u)
        session1.add(a)
        session1.commit()
    except Exception as e:
        session1.rollback()
        raise HTTPException(status_code=400, detail=e)

    return {"message": "success"}


@app.post("/api/sign-in/")
async def sign_in(login_id: str, password: str):
    if not login_id or not password:
        raise HTTPException(status_code=400, detail="Invalid request")

    u = session1.query(user.User).filter(user.User.login_id == login_id).first()
    if not u:
        raise HTTPException(status_code=400, detail="Invalid login_id or password")

    a = (
        session1.query(authentication.Authentication)
        .filter(authentication.Authentication.user_id == u.id)
        .first()
    )
    if not a:
        raise HTTPException(status_code=400, detail="Invalid login_id or password")

    if not a.verify_password(password):
        raise HTTPException(status_code=400, detail="Invalid login_id or password")

    token = a.generate_token()
    return {"token": token}


async def current_user(token: str = Depends(oauth2_scheme)):
    try:
        payload = jwt.decode(token, "secret", algorithms=["HS256"])
        user_id = payload.get("user_id")
        if user_id is None:
            raise HTTPException(status_code=401, detail="Invalid token")
    except JWTError:
        raise HTTPException(status_code=401, detail="Invalid token")

    u = session1.query(user.User).filter(user.User.id == user_id).first()
    if not u:
        raise HTTPException(status_code=401, detail="Invalid token")

    return u


@app.get("/api/prediction-models/")
async def get_prediction_models(user: user.User = Depends(current_user)):
    prediction_models = session1.query(prediction_model.PredictionModel).all()
    return {"prediction_models": prediction_models}


@app.post("/api/prediction_models/{prediction_model_id}")
async def predict(
    prediction_model_id: str, image: UploadFile, user: user.User = Depends(current_user)
):
    if not image:
        raise HTTPException(status_code=400, detail="No image found")

    img = Image.open(image.file)

    predict_model = (
        session1.query(prediction_model.PredictionModel)
        .filter(prediction_model.PredictionModel.id == prediction_model_id)
        .first()
    )
    if not predict_model:
        raise HTTPException(status_code=400, detail="Invalid prediction_model_id")

    predict_class, predict_prob = predict_model.predict(img)

    save_path = f"images/{image.filename}"
    p = prediction.Prediction(save_path, predict_class, user.id, predict_model.id)

    try:
        session1.add(p)
        session1.commit()
        img.save(save_path)
    except Exception as e:
        session1.rollback()
        raise HTTPException(status_code=400, detail=e)

    return {"class": predict_class, "probability": predict_prob}
