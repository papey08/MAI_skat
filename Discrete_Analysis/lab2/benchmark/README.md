# Что с этим делать?

В первую очередь компилируем файлы lab.cpp и map.cpp в lab и map соответственно.

С помощью testsGenerator.py генерируем тесты. Подробнее о его работе читайте в комментариях.
Например, введя в терминале это, создадим файл test.txt с 1000 пар на вставку, 1000 пар на поиск и 1000 пар на удаление (в сумме 3000 команд):

```()
python3 testsGenerator.py 1000
```

С помощью этой команды в терминале получим файл labRes.txt, в котором будет весь вывод программы lab, в конце файла будет время выполнения вставки, поиска и удаления:

```()
./lab < test.txt > labRes.txt
```

С помощью этой команды в терминале получим файл mapRes.txt, в котором будет весь вывод программы map, в конце файла будет время выполнения вставки, поиска и удаления:

```()
./map < test.txt > mapRes.txt
```