import torch
from models.resnet18 import ResNet18
from torchvision import transforms


def get_prediction(image, net_name, param_path, labels):
    tensor = transform_image(image)

    model = net_name2model(net_name, param_path, labels)
    model.eval()

    outputs = model.forward(tensor)
    _, pred = outputs.max(1)
    predicted_idx = int(pred.item())

    del model, outputs, pred
    return labels[predicted_idx]


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


def net_name2model(net_name, param_path, labels):
    if net_name == "ResNet18":
        model = ResNet18(len(labels))
        model.load_state_dict(
            torch.load(param_path, map_location=torch.device("cpu"))
        )
        return model
    else:
        raise ValueError("Invalid net name")
