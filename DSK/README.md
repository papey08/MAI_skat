# flights_delay

***Вместе со мной над проектом работали [someliiz](https://github.com/someliiz), [Stepan5024](https://github.com/Stepan5024), [haselius](https://github.com/haselius), [lumses](https://github.com/lumses).***

## Обучение нейросети

Чтобы обучить нейросеть локально, скачайте [архив](https://www.kaggle.com/datasets/usdot/flight-delays/download?datasetVersionNumber=1)
с датасетом и распакуйте файл `flights.csv` в `neural/dataset/flights.csv`.

## Запуск бота

### Настройка конфигураций

Вставьте свой токен в файл `config.yml` вместо строки `your-telegram-bot-token`

### Запуск бота локально

Чтобы запустить бота локально, установите зависимости:

```
pip3 install -r requirements.txt
```

Затем запустите main.py:

```
python3 main.py
```

Готово!

### Запуск бота в docker контейнере

```
docker-compose up
```

Готово!
