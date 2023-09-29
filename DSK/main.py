from aiogram import Bot, Dispatcher, types
from aiogram.contrib.fsm_storage.memory import MemoryStorage
from aiogram.dispatcher import FSMContext
from aiogram.dispatcher.filters.state import State, StatesGroup
from aiogram import executor

from tensorflow.keras.models import load_model

import logging
import yaml
import json
from datetime import datetime

ERROR_MESSAGE = 'Некорректные данные. Попробуйте снова или введите /start, чтобы начать заново'

# setting up logger
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger()

# getting token from config.yml
with open('config.yml', 'r') as config_file:
    config = yaml.safe_load(config_file)
    telegram_bot_token = config['token']

# setting up telegram bot
bot = Bot(token=telegram_bot_token)
storage = MemoryStorage()
dp = Dispatcher(bot, storage=storage)

# setting up model of the neural network and jsons
model = load_model('neural/model/flight_delay_model.h5')
airlines = json.load(open('json/airline_dict.json'))
dest_airports = json.load(open('json/destination_airport_dict.json'))
orig_airports = json.load(open('json/origin_airport_dict.json'))
airlines_list = json.load(open('json/airlines_list.json'))

def get_json_list(json_list):
    """
    get_json_list concatenates all pairs of key-value in the json_list
    """
    pairs = []
    for key, value in json_list.items():
        pair = f"{key} - {value}"
        pairs.append(pair)
    result = "\n".join(pairs)
    return result

def predict_delay(input_data):
    """
    Args:
        input_data: data given by the user

    Returns:
        output string based on result of neural network 
    """
    prediction = model.predict([input_data])[0][0]
    if prediction < 20:
        return 'Вылет будет вовремя. Не опаздывайте!'
    elif prediction < 50:
        return 'Предполагаемая задержка — 1-3 минуты'
    elif prediction < 100:
        return 'Предполагаемая задержка — 3-5 минут'
    elif prediction < 150:
        return 'Предполагаемая задержка — 5-10 минут'
    else:
        return 'Возможна задержка больше 10 минут'


class DataForm(StatesGroup):
    date = State()
    airline = State()
    origin_airport = State()
    destination_airport = State()
    scheduled_departure = State()


@dp.message_handler(commands=['start'], state='*')
async def start_command(message: types.Message):
    await message.reply('Привет! Я могу определить, на сколько задержится ваш авиарейс. Если нужна помощь, введите /help\n\nВведите дату вылета в формате ДД.ММ.ГГГГ')
    await DataForm.date.set()


@dp.message_handler(commands=['help'], state='*')
async def help_command(message: types.Message):
    await message.reply('Данный бот может предугадать задержку авиарейса на основе даты, планируемого времени вылета, а также авиалинии и аэропортов вылета и прибытия.\n' +
                        'Чтобы начать заново, введите /start\n' +
                        'Не стоит доверять боту слишком сильно. Все ошибаются.\n\n' +
                        'Поддерживаемые команды:\n /start — начать ввод заново\n' +
                        '/help — помощь\n' +
                        '/airlines — список поддерживаемых ботом авиалиний\n' +  
                        'Удачного использования!')


@dp.message_handler(commands=['airlines'], state='*')
async def airlines_command(message: types.Message):
    await message.reply('Список поддерживаемых авиалиний:\n' + get_json_list(airlines_list))


@dp.message_handler(state=DataForm.date, content_types=types.ContentTypes.TEXT)
async def handle_date(message: types.Message, state: FSMContext):
    
    # check if given message is correct
    def is_valid_date(date_to_valid):
        try:
            date = datetime.strptime(date_to_valid, '%d.%m.%Y')
            return date.weekday() + 1
        except ValueError:
            return None
    input_date = message.text
    day_of_week = is_valid_date(input_date)
    if day_of_week == None:
        await message.reply(ERROR_MESSAGE)
        return

    input_date = input_date.split('.')
    month = int(input_date[1])
    day = int(input_date[0])

    # going to the next state
    await state.update_data(date=(month, day, day_of_week))
    await message.reply("Введите двухбуквенный код авиалинии (можно узнать с помощью команды /airlines)")
    await DataForm.airline.set()


@dp.message_handler(state=DataForm.airline, content_types=types.ContentTypes.TEXT)
async def handle_airline(message: types.Message, state: FSMContext):

    # check if given message is correct
    if message.text not in airlines:
        await message.reply(ERROR_MESSAGE)
        return
    
    # going to the next state
    await state.update_data(airline=airlines[message.text])
    await message.reply('Введите трёхбуквенный код аэропорта вылета')
    await DataForm.origin_airport.set()


@dp.message_handler(state=DataForm.origin_airport, content_types=types.ContentTypes.TEXT)
async def handle_origin_airport(message: types.Message, state: FSMContext):

    # check if given message is correct
    if not message.text.isalpha():
        await message.reply(ERROR_MESSAGE)
        return
    if message.text not in orig_airports:
        await message.reply(ERROR_MESSAGE)
        return
    
    # going to the next state
    await state.update_data(origin_airport=orig_airports[message.text])
    await message.reply('Введите трёхбуквенный код аэропорта назначения')
    await DataForm.destination_airport.set()


@dp.message_handler(state=DataForm.destination_airport, content_types=types.ContentTypes.TEXT)
async def handle_destination_airport(message: types.Message, state: FSMContext):

    # check if given message is correct
    if not message.text.isalpha():
        await message.reply(ERROR_MESSAGE)
        return
    if message.text not in dest_airports:
        await message.reply(ERROR_MESSAGE)
        return
    
    # going to the next state
    await state.update_data(destination_airport=dest_airports[message.text])
    await message.reply('Введите время вылета в двадцатичетырёхчасовом формате ЧЧ:ММ')
    await DataForm.scheduled_departure.set()


@dp.message_handler(state=DataForm.scheduled_departure, content_types=types.ContentTypes.TEXT)
async def handle_scheduled_departure(message: types.Message, state: FSMContext):

    # check if given message is correct
    input_time = message.text.split(':')
    if len(input_time) != 2 or not input_time[0].isdigit() or not (input_time[1].isdigit() and len(input_time[1]) == 2):
        await message.reply(ERROR_MESSAGE)
        return
    h, m = map(int, input_time)
    if not (0 <= h <= 23 and 0 <= m <= 59):
        await message.reply(ERROR_MESSAGE)
        return
    
    # collecting all data given by user
    data = await state.get_data()
    month, day, day_of_week = data.get('date')
    airline = data.get('airline')
    origin_airport = data.get('origin_airport')
    destination_airport = data.get('destination_airport')
    scheduled_departure = h * 60 + m
    prediction_date = [month, day, day_of_week, airline, origin_airport, destination_airport, scheduled_departure]

    # getting a string with neural network prediction
    predicted_delay = predict_delay(prediction_date)
    await message.reply(predicted_delay)

    logger.info(f'date {day}.{month}, airline {airline}, origin_airport {origin_airport}, destination_airport {destination_airport}, scheduled_departure {scheduled_departure}, answer "{predicted_delay}"')
    await state.finish()


if __name__ == '__main__':
    executor.start_polling(dp, skip_updates=True)
