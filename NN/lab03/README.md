# Генеративные нейросети

## Задание

Необходимо натренировать и сравнить качество нескольких генеративных текстовых моделей на одном из заданных текстовых датасетов.

Необходимо исследовать следующие нейросетевые архитектуры:

1. Simple RNN с посимвольной и по-словной токенизацией
1. Однонаправленная однослойная и многослойная LSTM c посимвольной токенизацией и токенизацией по словам и [на основе BPE](https://keras.io/api/keras_nlp/tokenizers/byte_pair_tokenizer/)
1. Двунаправленная LSTM
1. (На хорошую оценку) трансформерная архитектура (GPT) "с нуля" [[пример](https://keras.io/examples/generative/text_generation_gpt/)] 
1. (На отличную оценку) до-обучение предобученной GPT-сети [[пример 1](https://github.com/ZotovaElena/RuGPT3_finetuning)]

Работа выполняется в группах по 3-5 человек, с таким расчетом, чтобы один участник группы отвечал за тренировку своей модели. 

## Датасеты

Рекомендуется использовать один из следующих датасетов, распределив их таким образом, чтобы все команды в группе использовали разные датасеты:

1. Английская литература с сайта [Project Gutenberg](https://www.gutenberg.org/)
1. Русская литература с сайта [lib.ru](http://lib.ru)
1. Архивы выборочных конференций сети FIDONet (можно найти на [archive.org](https://archive.org/download/usenet-fido7.ru) или по [magnet-ссылке](magnet:?xt=urn:btih:fa52bf91bbc33a6ce64d7e272f7c25ba252dba70&dn=usenet-fido7.ru&tr=http%3a%2f%2fbt1.archive.org%3a6969%2fannounce&tr=http%3a%2f%2fbt2.archive.org%3a6969%2fannounce&ws=http%3a%2f%2farchive.org%2fdownload%2f&ws=http%3a%2f%2fia601008.us.archive.org%2f2%2fitems%2f&ws=http%3a%2f%2fia601305.us.archive.org&ws=http%3a%2f%2fia801305.us.archive.org))
1. Текст книги [Гарри Поттер и методы рационального мышления](https://hpmor.ru/)
1. Англоязычные книги с Wikibooks ([датасет](https://www.kaggle.com/datasets/dhruvildave/wikibooks-dataset))
1. Русскоязычные книги с Wikibooks ([датасет](https://www.kaggle.com/datasets/dhruvildave/wikibooks-dataset))
1. Статьи с medium ([датасет](https://www.kaggle.com/datasets/fabiochiusano/medium-articles))
1. Субтитры фильмов ([датасет](https://www.kaggle.com/datasets/adiamaan/movie-subtitle-dataset))

> На отличную оценку можно также попробовать решить задачу генерации последовательности, отличной от текста, например музыкальных файлов. Музыкальный формат MIDI по сути дела содержит последовательность нот, которую можно генерировать.

## Отчет

Отчет приведите в файле [Report.md](Report.md). Также приложите к репозиторию набор Jupyter-ноутбуков, демонстрирующих процесс обучения моделей и результаты текстовой генерации.
