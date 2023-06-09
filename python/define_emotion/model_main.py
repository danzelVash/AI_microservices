# from ai_models import OurModel, HuggingModel
import cv2


import torch.nn.functional as F
import numpy as np
import torchvision.models as models
import torch.nn as nn
import torch
from transformers import AutoFeatureExtractor, AutoModelForImageClassification
import cv2


class OurModel:

    @staticmethod
    def predict_images(model, image_path):
        img = cv2.imread(image_path, cv2.IMREAD_COLOR)
        img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
        img = img.astype(np.float32)
        img = cv2.resize(img, (128, 128), interpolation=cv2.INTER_AREA)

        image_tensor = torch.from_numpy(img).unsqueeze(0).unsqueeze(0)
        with torch.no_grad():
            model.eval()
            output = model(image_tensor)
            prediction = F.softmax(output, dim=1).argmax(dim=1).item()

        return prediction

    @staticmethod
    def get_prediction(img_path):

        model_path = 'define_emotion/saved_epochs/emoji_last1_model_epoch_50.pth'
        model = models.resnet18()
        num_features = model.fc.in_features
        model.fc = nn.Linear(num_features, 7)
        model.conv1 = nn.Conv2d(1, 64, kernel_size=(7, 7), stride=(2, 2), padding=(3, 3), bias=False)
        model.load_state_dict(torch.load(model_path))
        model.eval()

        prediction = OurModel.predict_images(model, image_path=img_path)

        labels = ['Злость', 'Отвращение', 'Страх', 'Счастье', 'Грусть', 'Удивление', 'Нейтрально']

        return {f'{labels[prediction]}': 100}


class HuggingModel:

    @staticmethod
    def get_prediction(image_path):
        extractor = AutoFeatureExtractor.from_pretrained(
            "kdhht2334/autotrain-diffusion-emotion-facial-expression-recognition-40429105176")

        model = AutoModelForImageClassification.from_pretrained(
            "kdhht2334/autotrain-diffusion-emotion-facial-expression-recognition-40429105176")

        image = cv2.imread(image_path)

        image = cv2.cvtColor(image, cv2.COLOR_BGR2RGB)

        image_np = np.array(image)

        features = extractor(images=image_np, return_tensors="pt")
        outputs = model(**features)
        probabilities = F.softmax(outputs.logits, dim=1)
        new_prob = probabilities.detach().numpy()
        tensor_flat = np.ravel(new_prob)

        labels = ['Злость', 'Отвращение', 'Страх', 'Счастье', 'Нейтрально', 'Грусть', 'Удивление']
        emotions = {}
        for i in range(len(tensor_flat)):
            emotions[labels[i]] = round(round(tensor_flat[i], 2) * 100)

        return emotions


def detect_faces(pathToImg: str) -> bool:
    face_cascade = cv2.CascadeClassifier(cv2.data.haarcascades + 'haarcascade_frontalface_default.xml')
    image = cv2.imread(pathToImg)
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    faces = face_cascade.detectMultiScale(gray, scaleFactor=1.1, minNeighbors=5, minSize=(30, 30))
    if len(faces) > 0:
        return True
    return False


def describe_photo(pathToImg: str) -> str:
    # return OurModel.get_prediction(img_path=pathToImg)
    emotions = HuggingModel.get_prediction(image_path=pathToImg)

    max_weight, max_emotion = 0, ""
    for key in emotions.keys():
        if emotions[key] > max_weight:
            max_weight = emotions[key]
            max_emotion = key

    return max_emotion
