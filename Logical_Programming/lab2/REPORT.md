## Отчет по лабораторной работе №2
## по курсу "Логическое программирование"

## Решение логических задач

### студент: Попов Матвей Романович, группа М8О-208Б-20

## Результат проверки

| Преподаватель     | Дата         |  Оценка       |
|-------------------|--------------|---------------|
| Сошников Д.В. |              |               |
| Левинская М.А.|              |               |

> *Комментарии проверяющих *


## Введение

Решение логических задач — довольно нетипичная задача программирования, однако с ней легко справляются логические языки программирования, в частности Prolog, ведь в этом языке уже есть все необходимые инструменты для этого.

## Задание

Вариант 18: В нашем городе обувной магазин закрывается каждый понедельник, хозяйственный каждый вторник, продовольственный каждый четверг, а парфюмерный магазин работает только по понедельникам, средам и пятницам. В воскресенье все магазины закрыты. Однажды подруги Ася, Ира, Клава и Женя отправились за покупками, причем каждая в свой магазин и притом в один. По дороге они обменивались такими замечаниями. Ася. Женя и я хотели пойти вместе еще раньше на этой неделе, но не было такого дня, чтобы мы обе могли сделать наши покупки. Ира. Я не хотела идти сегодня, но завтра я уже не смогу купить то, что мне нужно. Клава. А я могла бы пойти в магазин и вчера и позавчера. Женя. А я могла бы пойти и вчера и завтра. Скажите, кому какой магазин нужен?

## Принцип решения

Для начала выпишем предикаты, обозначающие, в какой день недели тот или иной магазин открыт (или закрыт). Например, предикат `shop_is_closed(shoes, 1).` означает, что обувной магазин закрыт в первый день недели, то есть в понедельник. По условию задачи все четыре подруги отправились по магазинам в один день, то есть в день, когда открыты все четыре магазина. С помощью предиката `shop_is_open_day(X)` определим этот день. Это третий день недели, а именно среда. Далее используя условие задачи определим, в какие магазины идут подруги, отталкиваясь от того факта, что действие задачи происходит в среду. Определим предикат `solve` для каждой подруги, тем самым решим задачу.

### Исходный код файла `solution.pl`:
```
shop_is_closed(shoes, 1). % Вариант 18
shop_is_open(shoes, 2).
shop_is_open(shoes, 3).
shop_is_open(shoes, 4).
shop_is_open(shoes, 5).
shop_is_open(shoes, 6).
shop_is_closed(shoes, 7).
shop_is_open(house, 1).
shop_is_closed(house, 2).
shop_is_open(house, 3).
shop_is_open(house, 4).
shop_is_open(house, 5).
shop_is_open(house, 6).
shop_is_closed(house, 7).
shop_is_open(food, 1).
shop_is_open(food, 2).
shop_is_open(food, 3).
shop_is_closed(food, 4).
shop_is_open(food, 5).
shop_is_open(food, 6).
shop_is_closed(food, 7).
shop_is_open(perf, 1).
shop_is_closed(perf, 2).
shop_is_open(perf, 3).
shop_is_closed(perf, 4).
shop_is_open(perf, 5).
shop_is_closed(perf, 6).
shop_is_closed(perf, 7).

remove([], _Elem, []). % предикат удаления элемента из списка
remove([Elem|T], Elem, TR) :- remove(T, Elem, TR), !.
remove([H|T], Elem, [H|TR]) :- remove(T, Elem, TR).

list_of_shop_is_open_days(Shop, L) :-
    findall(X, shop_is_open(Shop, X), L).
    
member(A, [A|_]). % предикат проверки вхождения элемента в список 
member(A, [_|Z]) :- member(A, Z).

list_conjunction([],_,[]).
list_conjunction([H|T],Y,[H|R]) :- member(H,Y),list_conjunction(T,Y,R), !.
list_conjunction([_|T],Y,R) :- list_conjunction(T,Y,R), !.

head_return([H|_], H).

shop_is_open_day(X) :-
    list_of_shop_is_open_days(shoes, S),
    list_of_shop_is_open_days(house, H),
    list_of_shop_is_open_days(food, F),
    list_of_shop_is_open_days(perf, P),
    list_conjunction(S, H, R1),
    list_conjunction(F, P, R2),
    list_conjunction(R1, R2, R),
    head_return(R, X).

klava(X, Shop) :- % могла пойти в магазин и вчера, и позавчера
    X1 is X - 1,
    findall(R1, shop_is_open(R1, X1), L1),
    X2 is X - 2,
    findall(R2, shop_is_open(R2, X2), L2),
    list_conjunction(L1, L2, R),
    head_return(R, Shop).

zhenya(X, Shop) :- % могла пойти и вчера, и завтра
    X1 is X - 1,
    findall(R1, shop_is_open(R1, X1), L1),
    X2 is X + 1,
    findall(R2, shop_is_open(R2, X2), L2),
    list_conjunction(L1, L2, R),
    head_return(R, Shop).

ira(X, Shop) :- % завтра её магазин будет закрыт
    X1 is X + 1,
    findall(R, shop_is_closed(R, X1), L),
    solve(klava, KS),
    solve(zhenya, ZS),
    remove(L, KS, R1),
    remove(R1, ZS, R2),
    head_return(R2, Shop).

solve(zhenya, Shop) :-
    shop_is_open_day(X),
    zhenya(X, Shop).

solve(klava, Shop) :-
    shop_is_open_day(X),
    klava(X, Shop).

solve(ira, Shop) :-
    shop_is_open_day(X),
    ira(X, Shop).

solve(asya, Shop) :-
    solve(zhenya, ZS),
    solve(klava, KS),
    solve(ira, IS),
    remove([food, perf, shoes, house], ZS, R1),
    remove(R1, KS, R2),
    remove(R2, IS, R3),
    head_return(R3, Shop).
```
### Пример работы программы:
```
1 ?- solve(asya, X).
X = house.

2 ?- solve(zhenya, X).
X = shoes.

3 ?- solve(ira, X).
X = perf.

4 ?- solve(klava, X).
X = food.

5 ?- solve(X, house).
X = asya.

6 ?- solve(X, shoes).
X = zhenya .

7 ?- solve(X, perf).
X = ira .

8 ?- solve(X, food).
X = klava .
```
## Выводы

Проделав лабораторную  работу, я составил программу для решения логической задачи, получил непротиворечивое и единственное решение этой задачи, тем самым я научился решать логические задачи с помощью, что вполне логично, логического программирования.  
Ответ на задачу из варианта №18: Ася идёт в хозяйственный магазин, Женя идёт в обувной магазин, Ира идёт в парфюмерный магазин, Клава идёт в продуктовый магазин.



