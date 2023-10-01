import torch.nn as nn
from torchvision.models import resnet18


class EggplantNet(nn.Module):
    def __init__(self):
        super(EggplantNet, self).__init__()
        self.model = resnet18(pretrained=True)
        self.model.fc = nn.Linear(512, 7)

    def forward(self, x):
        return self.model(x)
