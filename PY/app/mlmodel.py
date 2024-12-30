from entities.catrgories import categories

import joblib


class MlModel:
    def __init__(self, path_to_model: str):
        self.model = joblib.load(path_to_model)

    def predict(self, area, duration, budget, time, type_l):
        return categories[int(self.model.predict([[area, duration, budget, time, type_l]]))]
