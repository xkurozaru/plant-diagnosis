from typing import List
from fastapi import FastAPI, HTTPException, UploadFile, Form
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


@app.post("/predictor/predict")
async def predict(image: UploadFile, net_name: str = Form(), param_path: str = Form(), labels: List[str] = Form()):
    if not image:
        raise HTTPException(status_code=400, detail="No image found")

    img = Image.open(image.file)
    predict_class = get_prediction(img, net_name, param_path, labels)
    img.close()
    return {"class": predict_class}
