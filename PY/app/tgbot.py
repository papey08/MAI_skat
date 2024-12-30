import app.service as service
from entities.locations import locations

import telebot
from telebot.types import ReplyKeyboardMarkup, KeyboardButton, ReplyKeyboardRemove


class Tgbot:
    def __init__(self, api_key: str, service: service.Service):
        self.token = api_key
        self.bot = telebot.TeleBot(self.token)
        self.service = service
        self.user_data = {}


    def start_handler(self):
        @self.bot.message_handler(commands=['start'])
        def start_message(message):
            chat_id = message.chat.id
            username = message.from_user.username or "Гость"  # Если username отсутствует, используем заглушку
            self.user_data[chat_id] = {'username': username}
            self.bot.send_message(chat_id, "Привет, я могу подсказать, куда можно сходить в Москве. Приступим?")
            self.ask_area(message)


    def run(self):
        self.start_handler()
        self.bot.polling(none_stop=True, interval=0)

    def ask_area(self, message):
        markup = ReplyKeyboardMarkup(resize_keyboard=True, one_time_keyboard=True)
        markup.add(KeyboardButton("Улица"), KeyboardButton("Помещение"))
        self.bot.send_message(
            message.chat.id,
            "Где бы вы хотели отдохнуть? Выберите один из вариантов:",
            reply_markup=markup,
        )
        self.bot.register_next_step_handler(message, self.process_area)

    def process_area(self, message):
        if message.text == "Улица":
            self.user_data[message.chat.id]['area'] = 1
        elif message.text == "Помещение":
            self.user_data[message.chat.id]['area'] = 2
        else:
            self.bot.send_message(message.chat.id, "Пожалуйста, выберите один из вариантов.")
            return self.ask_area(message)
        self.ask_duration(message)

    def ask_duration(self, message):
        markup = ReplyKeyboardMarkup(resize_keyboard=True, one_time_keyboard=True)
        markup.add(KeyboardButton("1 час"), KeyboardButton("3 часа"), KeyboardButton("6 часов"))
        self.bot.send_message(
            message.chat.id,
            "Сколько времени вы хотите отдыхать?",
            reply_markup=markup,
        )
        self.bot.register_next_step_handler(message, self.process_duration)

    def process_duration(self, message):
        duration_map = {"1 час": 1, "3 часа": 3, "6 часов": 6}
        if message.text in duration_map:
            self.user_data[message.chat.id]['duration'] = duration_map[message.text]
        else:
            self.bot.send_message(message.chat.id, "Пожалуйста, выберите один из вариантов.")
            return self.ask_duration(message)
        self.ask_budget(message)

    def ask_budget(self, message):
        markup = ReplyKeyboardMarkup(resize_keyboard=True, one_time_keyboard=True)
        markup.add(KeyboardButton("До 1000 рублей"), KeyboardButton("Больше 1000 рублей"))
        self.bot.send_message(
            message.chat.id,
            "Какой у вас бюджет на отдых?",
            reply_markup=markup,
        )
        self.bot.register_next_step_handler(message, self.process_budget)

    def process_budget(self, message):
        if message.text == "До 1000 рублей":
            self.user_data[message.chat.id]['budget'] = 1000
        elif message.text == "Больше 1000 рублей":
            self.user_data[message.chat.id]['budget'] = 10000
        else:
            self.bot.send_message(message.chat.id, "Пожалуйста, выберите один из вариантов.")
            return self.ask_budget(message)
        self.ask_time(message)

    def ask_time(self, message):
        markup = ReplyKeyboardMarkup(resize_keyboard=True, one_time_keyboard=True)
        markup.add(KeyboardButton("Утро"), KeyboardButton("День"), KeyboardButton("Вечер"), KeyboardButton("Ночь"))
        self.bot.send_message(
            message.chat.id,
            "Какое время суток вам подходит?",
            reply_markup=markup,
        )
        self.bot.register_next_step_handler(message, self.process_time)

    def process_time(self, message):
        time_map = {"Утро": 1, "День": 2, "Вечер": 3, "Ночь": 4}
        if message.text in time_map:
            self.user_data[message.chat.id]['time'] = time_map[message.text]
        else:
            self.bot.send_message(message.chat.id, "Пожалуйста, выберите один из вариантов.")
            return self.ask_time(message)
        self.ask_type(message)

    def ask_type(self, message):
        markup = ReplyKeyboardMarkup(resize_keyboard=True, one_time_keyboard=True)
        markup.add(KeyboardButton("Активный"), KeyboardButton("Пассивный"))
        self.bot.send_message(
            message.chat.id,
            "Какой тип отдыха вы предпочитаете?",
            reply_markup=markup,
        )
        self.bot.register_next_step_handler(message, self.process_type)

    def process_type(self, message):
        if message.text == "Активный":
            self.user_data[message.chat.id]['type_l'] = 1
        elif message.text == "Пассивный":
            self.user_data[message.chat.id]['type_l'] = 2
        else:
            self.bot.send_message(message.chat.id, "Пожалуйста, выберите один из вариантов.")
            return self.ask_type(message)
        self.ask_location(message)

    def ask_location(self, message):
        markup = ReplyKeyboardMarkup(resize_keyboard=True, one_time_keyboard=True)
        markup.add(
            KeyboardButton("ЦАО"), 
            KeyboardButton("САО"),
            KeyboardButton("СВАО"), 
            KeyboardButton("ВАО"),
            KeyboardButton("ЮВАО"), 
            KeyboardButton("ЮАО"),
            KeyboardButton("ЮЗАО"), 
            KeyboardButton("ЗАО"),
            KeyboardButton("СЗАО"),
        )
        self.bot.send_message(
            message.chat.id,
            "В каком округе Москвы предпочитаете отдохнуть?",
            reply_markup=markup,
        )
        self.bot.register_next_step_handler(message, self.process_location)

    def process_location(self, message):
        if message.text not in locations:
            self.bot.send_message(message.chat.id, "Пожалуйста, выберите один из вариантов.")
            return self.ask_location(message)
        else:
            self.user_data[message.chat.id]['location'] = message.text
        self.make_prediction(message)

    def make_prediction(self, message):
        user_answers = self.user_data[message.chat.id]
        username = user_answers['username']
        area = user_answers['area']
        duration = user_answers['duration']
        budget = user_answers['budget']
        time = user_answers['time']
        type_l = user_answers['type_l']
        location = user_answers['location']

        # Вызов find_place
        recommendation = self.service.find_place(username, area, duration, budget, time, type_l, location)
        self.bot.send_message(
            message.chat.id,
            f"Рекомендация: {recommendation.name}\n\nСсылка: {recommendation.url}",
            reply_markup=ReplyKeyboardRemove(),
        )
