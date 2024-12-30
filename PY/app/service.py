import app.mlmodel as mlmodel
import app.db as db


class Service:
    def __init__(self, neuro: mlmodel.MlModel, db: db.Db):
        self.neuro = neuro
        self.db = db

    def find_place(self, username: str, area, duration, budget, time, type_l, location):
        """
        :params:
            username: ник в тг
            area: улица - 1, помещение - 2
            duration: 1 час = 1, 3 часа = 3, 6 часов = 6
            budget: до 1000 рублей = 1000, больше 1000 = 10000
            time: утро = 1, день = 2, вечер = 3, ночь = 4
            type_l: активный = 1, пассивный = 2
            location: трехбуквенное обозначение округа (ЦАО, САО и тд)
        :returns:
            place: entities.Place
        """
        category = self.neuro.predict(area, duration, budget, time, type_l)
        place = self.db.find_place(category, location)
        self.db.add_user(username)
        self.db.add_response(username, place.id)
        return place
