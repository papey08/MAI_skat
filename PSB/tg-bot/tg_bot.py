from servise import Servise

import telebot
from telebot.types import ReplyKeyboardMarkup, KeyboardButton, ReplyKeyboardRemove


class Tgbot:
    def __init__(self, api_key: str, service: Servise):
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
            self.bot.send_message(chat_id, "Привет, я умею анализировать отзывы. Приступим?")
            self.сhoose_category(message)

    def run(self):
        self.start_handler()
        self.bot.polling(none_stop=True, interval=0)

    def сhoose_category(self, message):
        markup = ReplyKeyboardMarkup(resize_keyboard=True, one_time_keyboard=True)
        markup.add(KeyboardButton("жалоба"), KeyboardButton("предложение"), KeyboardButton("благодарность"))
        self.bot.send_message(
            message.chat.id,
            "Выберите категорию:",
            reply_markup=markup,
        )
        self.bot.register_next_step_handler(message, self.process_change_category)

    def process_change_category(self, message):
        markup = ReplyKeyboardMarkup(resize_keyboard=True, one_time_keyboard=True)
        markup.add(KeyboardButton("<-"), KeyboardButton("Выбрать категорию"), KeyboardButton("->"))
        response = self.service.get_current_response_in_category(message.text, 1)
        if not response:
            self.сhoose_category(message)
            return
        else:
            self.bot.send_message(
                message.chat.id,
                f"Категория: {response.resp_category}\n\n"
                f"Текст: {response.original_text}\n\n"
                f"Ссылка: https://www.banki.ru/services/responses/bank/response/{response.id}/",
                reply_markup=markup,
            )
            self.user_data[0] = response
            self.bot.register_next_step_handler(message, self.process_next)
            return

    def process_next(self, message):
        if message.text == "Выбрать категорию":
            self.сhoose_category(message)
            return
        markup = ReplyKeyboardMarkup(resize_keyboard=True, one_time_keyboard=True)
        markup.add(KeyboardButton("<-"), KeyboardButton("Выбрать категорию"), KeyboardButton("->"))
        response = self.user_data[0]
        last_response = response
        if message.text == "<-":
            response = self.service.get_previous_response_in_category(category=response.resp_category,
                                                                      current_id=response.current_index)
        elif message.text == "->":
            response = self.service.get_next_response_in_category(category=response.resp_category,
                                                                  current_id=response.current_index)
        if not response:
            response = last_response
        self.bot.send_message(
            message.chat.id,
            f"Категория: {response.resp_category}\n\n"
            f"Текст: {response.original_text}\n\n"
            f"Ссылка: https://www.banki.ru/services/responses/bank/response/{response.id}/",
            reply_markup=markup,
        )
        self.user_data[0] = response
        self.bot.register_next_step_handler(message, self.process_next)
        return
