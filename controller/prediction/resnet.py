import torch.nn as nn
from torchvision.models import resnet18


class Resnet18(nn.Module):
    def __init__(self, num_classes=7):
        super(Resnet18, self).__init__()
        self.model = resnet18(pretrained=True)
        self.model.fc = nn.Linear(512, num_classes)

    def forward(self, x):
        return self.model(x)
