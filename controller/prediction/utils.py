from PIL import Image
from torchvision import transforms


def transform_image(image: Image.Image):
    my_transforms = transforms.Compose(
        [
            transforms.Resize(224),
            transforms.CenterCrop(224),
            transforms.ToTensor(),
            transforms.Normalize([0.485, 0.456, 0.406], [0.229, 0.224, 0.225]),
        ]
    )
    return my_transforms(image).unsqueeze(0)
