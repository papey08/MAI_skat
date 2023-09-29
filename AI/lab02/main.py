from pyknow import *

# Visualisation: https://wbd.ms/share/v2/aHR0cHM6Ly93aGl0ZWJvYXJkLm1pY3Jvc29mdC5jb20vYXBpL3YxLjAvd2hpdGVib2FyZHMvcmVkZWVtLzJlZjU2MDliM2ViMTQxNDc4MWFmMDc0MzFiODYzNDNlX0JCQTcxNzYyLTEyRTAtNDJFMS1CMzI0LTVCMTMxRjQyNEUzRF8zM2NiMDljOC04MGI3LTQzM2QtOGY0Yy1mZGEwOGY0YTQxY2Q=


class Banket(KnowledgeEngine):
    result = []

    # c_restaurant (Не оч ресторан)
    @Rule(
        AND(
            OR(Fact(price='low'), Fact(price='mid')),
            Fact(event='after job'),
            OR(Fact(amount='1'), Fact(amount='2-3')),
        )
    )
    def c_restaurant(self):
        self.declare(Fact('c_restaurant'))

    # b_restaurant (Норм ресторан)
    @Rule(
        AND(
            OR(Fact(event='holiday'), Fact(event='date')), Fact(price='mid'),
            OR(Fact(amount='1'), Fact(amount='2-3'))
        )
    )
    def b_restaurant(self):
        self.declare(Fact('b_restaurant'))

    # a_restaurant (Крутой ресторан)
    @Rule(
        AND(
            OR(Fact(event='holiday'), Fact(event='date')),
            OR(Fact(amount='1'), Fact(amount='2-3')),
            Fact(price='high'),
        )
    )
    def a_restaurant(self):
        self.declare(Fact('a_restaurant'))

    # pizzeria
    @Rule(
        AND(
            OR(Fact(amount='2-3'), Fact(amount='more')),
            OR(Fact(event='after job'), Fact(event='holiday')),
            Fact(kitchen='Italy'),
        )
    )
    def pizzeria(self):
        self.declare(Fact('pizzeria'))

    # Pizza Hut

    @Rule(Fact('pizzeria'), Fact(price='low'), Fact(location='Вся Москва'))
    def pizza_hut(self):
        self.declare(Fact(banket='Pizza Hut'))

    # Zotman

    @Rule(Fact('pizzeria'), Fact(price='high'))
    @Rule(Fact('pizzeria'), Fact(price='high'), Fact(location='Вся Москва'))
    def zotman(self):
        self.declare(Fact(banket='Zotman'))

    # Dodo

    @Rule(Fact('pizzeria'), Fact(price='mid'))
    @Rule(Fact('pizzeria'), Fact(price='mid'), Fact(location='Вся Москва'))
    def dodo_pizza(self):
        self.declare(Fact(banket='Додо Пицца'))

    # Марукамэ
    @Rule(Fact(kitchen='Asian'), Fact('c_restaurant'))
    @Rule(Fact(kitchen='Asian'), Fact('c_restaurant'), OR(Fact(location='САО'), Fact(location='ЦАО'), Fact(location='ЮЗАО')))
    def marucame(self):
        self.declare(Fact(banket='Марукамэ'))

    # Тануки
    @Rule(Fact(kitchen='Asian'), Fact('b_restaurant'))
    @Rule(Fact(kitchen='Asian'), Fact('b_restaurant'), Fact(location='Вся Москва'))
    def tanuki(self):
        self.declare(Fact(banket='Тануки'))

    # MEGUmi
    @Rule(Fact(kitchen='Asian'), Fact('a_restaurant'))
    @Rule(Fact(kitchen='Asian'), Fact('a_restaurant'), Fact(location='ЦАО'))
    def megumi(self):
        self.declare(Fact(banket='MEGUmi'))

    # Бургер Кинг
    @Rule(Fact(kitchen='American'), Fact('c_restaurant'))
    @Rule(Fact(kitchen='American'), Fact('c_restaurant'), Fact(location='Вся Москва'))
    def burger_king(self):
        self.declare(Fact(banket='Burger King'))

    # FARШ
    @Rule(Fact(kitchen='American'), Fact('b_restaurant'))
    @Rule(Fact(kitchen='American'), Fact('b_restaurant'), Fact(location='Вся Москва'))
    def FARSH(self):
        self.declare(Fact(banket='FARSH'))

    # Goodman Prime
    @Rule(Fact(kitchen='American'), Fact('a_restaurant'))
    @Rule(Fact(kitchen='American'), Fact('a_restaurant'), Fact(location='ЦАО'))
    def goodman_prime(self):
        self.declare(Fact(banket='Goodman Prime'))

    # Elardgi
    @Rule(Fact(kitchen='Georgian'), Fact('b_restaurant'))
    @Rule(Fact(kitchen='Georgian'), Fact('b_restaurant'), Fact(location='ЦАО'))
    def elargi(self):
        self.declare(Fact(banket='Эларджи'))

    # Ткемали
    @Rule(Fact(kitchen='Georgian'), Fact('a_restaurant'))
    @Rule(Fact(kitchen='Georgian'), Fact('a_restaurant'), Fact(location='ЦАО'))
    def tkemali(self):
        self.declare(Fact(banket='Ткемали'))

    # Bella Pasta
    @Rule(Fact(kitchen='Italian'), Fact('b_restaurant'))
    @Rule(Fact(kitchen='Italian'), Fact('b_restaurant'), Fact(location='ЦАО'))
    def bella_pasta(self):
        self.declare(Fact(banket='Bella Pasta'))

    # Accenti
    @Rule(Fact(kitchen='Italian'), Fact('a_restaurant'))
    @Rule(Fact(kitchen='Italian'), Fact('a_restaurant'), Fact(location='ЦАО'))
    def accenti(self):
        self.declare(Fact(banket='Accenti'))

    # Ледокол
    @Rule(Fact(kitchen='Russian'), Fact('c_restaurant'))
    @Rule(Fact(kitchen='Russian'), Fact('c_restaurant'), Fact(location='САО'))
    def ledokol(self):
        self.declare(Fact(banket='Ледокол'))

    # Теремок
    @Rule(Fact(kitchen='Russian'), Fact('b_restaurant'))
    @Rule(Fact(kitchen='Russian'), Fact('b_restaurant'), Fact(location='Вся Москва'))
    def teremok(self):
        self.declare(Fact(banket='Теремок'))

    # Матрёшка
    @Rule(Fact(kitchen='Russian'), Fact('a_restaurant'))
    @Rule(Fact(kitchen='Russian'), Fact('a_restaurant'), Fact(location='ЦАО'))
    def matryoshka(self):
        self.declare(Fact(banket='Матрёшка'))

    # Le Petit Paris
    @Rule(Fact(kitchen='French'), Fact('c_restaurant'))
    @Rule(Fact(kitchen='French'), Fact('c_restaurant'), Fact(location='ЦАО'))
    def le_petit_paris(self):
        self.declare(Fact(banket='Le Petit Paris'))

    # Vаниль
    @Rule(Fact(kitchen='French'), Fact('b_restaurant'))
    @Rule(Fact(kitchen='French'), Fact('b_restaurant'), Fact(location='ЦАО'))
    def vanil(self):
        self.declare(Fact(banket='Vаниль'))

    # Клод Монэ
    @Rule(Fact(kitchen='French'), Fact('a_restaurant'))
    @Rule(Fact(kitchen='French'), Fact('a_restaurant'), Fact(location='ЦАО'))
    def Klod_Mone(self):
        self.declare(Fact(banket='Клод Монэ'))

    @Rule(Fact(banket=MATCH.a))
    def print_result(self, a):
        print(f'Банкет в {a}')
        

    def factz(self, l):
        for x in l:
            self.declare(x)
        


user_ans = []
questions = {
    'Когда?': [['1. after job', '2. holiday', '3. date'], '1. После работы  2. Праздник  3. Свидание'],
    'Сколько людей?': [['1. 1', '2. 2-3', '3. more'], '1. один  2. два-три  3. Много'],
    'Бюджет?': [['1. low', '2. mid', '3. high'], '1. мало  2. средний  3. много'],
    'Кухня?': [['1. French', '2. Russian', '3. Georgian', '4. Italian', '5. American', '6. Asian'], '1. французская    2. русская    3. грузинская    4.итальянская    5.американская   6. японская'],
    'Район?': [['1. Вся Москва', "2. САО", "3. ЦАО", "4. ЮЗАО"], '1. Вся Москва  2. САО  3. ЦАО  4. ЮЗАО']
}

for question in questions:
    print(question)
    print(questions[question][1])
    user_answer = input()
    answer = ''
    user_answer_list = []
    for option in questions[question][0]:
        if option.startswith(user_answer):
            answer = option.split('.')[1].strip()
    user_ans.append(answer)

ex1 = Banket()
ex1.reset()

facts_my = [
    Fact(event=user_ans[0]),
    Fact(amount=user_ans[1]),
    Fact(price=user_ans[2]),
    Fact(kitchen=user_ans[3]),
    Fact(location=user_ans[4])
]

ex1.factz(facts_my)
ex1.run()
