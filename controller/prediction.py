import json

import torch
from models.eggplant import EggplantNet
from torchvision import transforms


def get_prediction(image):
    tensor = transform_image(image)

    model = EggplantNet()
    model.load_state_dict(
        torch.load("./models/eggplant.pth", map_location=torch.device("cpu"))
    )
    model.eval()
    outputs = model.forward(tensor)
    prob, pred = outputs.max(1)
    predicted_idx = str(pred.item())
    probability = str(prob.item())

    del model, outputs, prob, pred
    class_index = json.load(open("./index/eggplant_class_index.json"))
    return class_index[predicted_idx], probability


def transform_image(image):
    my_transforms = transforms.Compose(
        [
            transforms.Resize(224),
            transforms.CenterCrop(224),
            transforms.ToTensor(),
            transforms.Normalize([0.485, 0.456, 0.406], [0.229, 0.224, 0.225]),
        ]
    )
    return my_transforms(image).unsqueeze(0)
