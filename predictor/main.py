from fastapi import FastAPI, HTTPException, UploadFile
from fastapi.middleware.cors import CORSMiddleware
from PIL import Image
from prediction import get_prediction

app = FastAPI()

# CORSミドルウェアの設定
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/")
async def index():
    return {"message": "Hello World"}


@app.post("/api/predictor/predict")
async def predict(image: UploadFile):
    if not image:
        raise HTTPException(status_code=400, detail="No image found")

    img = Image.open(image.file)
    predict_class, predict_prob = get_prediction(img)
    return {"class": predict_class, "probability": predict_prob}
