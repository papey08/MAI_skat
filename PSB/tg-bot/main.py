from servise import Servise
from tg_bot import Tgbot
from repository import Repository


def main():
    api_key = "7624701922:AAFQfhExmM4OPiHr4PaWftK-XAQmSsZTSw0"
    db = Repository('localhost',
                    5501,
                    'root',
                    'root',
                    'psb-case')
    service = Servise(db)

    # Создайте и запустите бота
    bot = Tgbot(api_key=api_key, service=service)
    bot.run()


if __name__ == "__main__":
    main()