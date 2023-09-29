# Отчет по лабораторной работе №1
## Работа со списками и реляционным представлением данных
## по курсу "Логическое программирование"

### студент: Попов Матвей Романович, группа М8О-208Б-20, вариант 18

## Результат проверки

| Преподаватель     | Дата         |  Оценка       |
|-------------------|--------------|---------------|
| Сошников Д.В. |              |               |
| Левинская М.А.|              |               |

> *Комментарии проверяющих*


## Введение
Списки в прологе сильно отличаются от того, что обычно называют списками в императивных языках программирования. В прологе любой список разделён на голову и хвост, причём хвост всегда представлен только одним элементом. Работа со списками в прологе немного напоминает работу с очередями (так как обход списка чаще всего делается через отделение головного элемента от остальной части) или массивами.


## Задание 1.1: Предикат обработки списка

### Стандартные предикаты для работы со списками:

```prolog
length_([], 0). % предикат получения длины списка
length_([_|Y], N) :- length_(Y, N1), N is N1 + 1.

member(A, [A|_]). % предикат проверки вхождения элемента в список 
member(A, [_|Z]) :- member(A, Z).

append([], List2, List2). % предикат объединения списков
append([H|T], List2, [H|TR]) :- append(T, List2, TR).

remove([], _Elem, []). % предикат удаления элемента из списка
remove([Elem|T], Elem, TR) :- remove(T, Elem, TR), !.
remove([H|T], Elem, [H|TR]) :- remove(T, Elem, TR).

permute([],[]). % предикат проверки перестановки списка
permute(L,[X|T]) :- remove(L,X,R), permute(R,T).

sub_start([], _List). % предикат проверки подсписка
sub_start([H|TSub], [H|TList]) :- sub_start(TSub, TList).
sublist(Sub, List) :- sub_start(Sub, List), !.
sublist(Sub, [_H|T]) :- sublist(Sub, T).
```

`shift1(List, Res)` — циклический сдвиг списка вправо, основанный на стандартных предикатах.  
`shift2(List)` — циклический сдвиг списка вправо без стандартных предикатов.

### Реализация:
```prolog
% Task 1.1
% Циклический сдвиг списка вправо

shift1(List, Res) :- % предикат, основанный на стандартных предикатах
    reverse(List, [Res1H|Res1T]),
    append([], [Res1H], Res2),
    reverse(Res1T, Res3),
    append(Res2, Res3, Res).

print_last([_|T]) :-
    print_last(T).
print_last([T]) :-
    write(T),
    write(" ").
print_first([_|[]]).
print_first([H|T]) :-
    write(H),
    write(" "),
    print_first(T).

shift2([H|T]) :- % предикат без стандартных предикатов
    print_last(T),
    write(H),
    write(' '),
    print_first(T).


/* Tests:
?- shift1([1, 2, 3], R).
?- shift2([1, 2, 3]).
*/
```

Описание работы предиката `shift1`: переворачиваем исходный список, голову перевёрнутого списка кладём в результат, а хвост переворачиваем и кладём в результат.  
Описание работы предиката `shift2`: вызываем рекурсивные предикаты `print_last` (выводит последний элемент) и `print_first`(выводит все элементы, кроме последнего), в результате чего выведется список с циклическим сдвигом. 

## Задание 1.2: Предикат обработки числового списка

`min([MinElem], MinElem)` — вычисление минимального элемента.

### Реализация:
```prolog
% Task 1.2
% Вычисление минимального элемента

min([MinElem], MinElem). % предикат из задания
min([H|T], MinElem) :- min(T, TMinElem), TMinElem < H, !, MinElem = TMinElem; MinElem = H.

/* Tests:
?- min([1, 2, 3], N).
?- min([0, 0, 0], N).
?- min([-2, -1, 0, 1, 2], N).
*/
```
Описание работы предиката: рекурсивно обходим список, находя минимальный элемент для каждого из хвостов списка, результатом работы предиката будет либо головной элемент, либо минимальный элемент хвоста.

## Задание 2: Реляционное представление данных

### Исходный код
```prolog
% USE UTF-8


% представление three.pl, вариант с заадниями 1

:- set_prolog_flag(encoding, utf8).

:- ['three.pl'].


% 1) Получить таблицу групп и средний балл по каждой из групп

print_list([]). % предикат вывода элементов списка
print_list([H|T]) :-
    write(H),
    write('\n'),
    print_list(T).

append1([], List2, List2). % предикат объединения списков
append1([Head|Tail], List2, [Head|TailResult]):-
   append1(Tail, List2, TailResult).

group_grades_list(Num, L) :- % предикат получения списка со всеми оценками группы
    findall(X, student(Num, _, [grade('LP',X),grade('MTH',_),grade('FP',_),grade('INF',_),grade('ENG',_),grade('PSY',_)]), LP),
    findall(X, student(Num, _, [grade('LP',_),grade('MTH',X),grade('FP',_),grade('INF',_),grade('ENG',_),grade('PSY',_)]), MTH),
    findall(X, student(Num, _, [grade('LP',_),grade('MTH',_),grade('FP',X),grade('INF',_),grade('ENG',_),grade('PSY',_)]), FP),
    findall(X, student(Num, _, [grade('LP',_),grade('MTH',_),grade('FP',_),grade('INF',X),grade('ENG',_),grade('PSY',_)]), INF),
    findall(X, student(Num, _, [grade('LP',_),grade('MTH',_),grade('FP',_),grade('INF',_),grade('ENG',X),grade('PSY',_)]), ENG),
    findall(X, student(Num, _, [grade('LP',_),grade('MTH',_),grade('FP',_),grade('INF',_),grade('ENG',_),grade('PSY',X)]), PSY),
    append1(LP, MTH, A),
    append1(FP, INF, B),
    append1(ENG, PSY, C),
    append1(A, B, D),
    append1(C, D, L).

sum_list([],0). % предикат суммы списка
sum_list([H|T], S) :-
	sum_list(T, S1), 
	S is H + S1.

group_grades_sum(Num, Sum) :- % предикат нахождения суммы оценок группы
    group_grades_list(Num, L),
    sum_list(L, Sum).

group_grades_av(Num, Av) :- % предикат нахождения средней оценки группы
    group_grades_sum(Num, Sum),
    group_grades_list(Num, L),
    length(L, Le),
    Av is Sum / Le.

all_groups_grades_av() :- % предикат вывода средних оценок всех групп
    write("Average of 101 is "),
    group_grades_av(101, A),
    write(A),
    write('\n'),
    write("Average of 102 is "),
    group_grades_av(102, B),
    write(B),
    write('\n'),
    write("Average of 103 is "),
    group_grades_av(103, C),
    write(C),
    write('\n'),
    write("Average of 104 is "),
    group_grades_av(104, D),
    write(D),
    write('\n').

group_list(Num, L) :- % предикат нахождения списка группы 
    findall(X, student(Num, X, _), L). 

group_table() :- % предикат вывода списков всех групп
    write("\nGroup 101:\n"),
    group_list(101, A),
    print_list(A),
    write("\nGroup 102:\n"),
    group_list(102, B),
    print_list(B),
    write("\nGroup 103:\n"),
    group_list(103, C),
    print_list(C),
    write("\nGroup 104:\n"),
    group_list(104, D),
    print_list(D),
    write('\n').

task1() :- % предикат, выполняющий первое подзадание из трёх
    group_table(),
    all_groups_grades_av().


% 2) Для каждого предмета получить список студентов, не сдавших экзамен (grade=2)

lp_grade2_list(L) :- % предикат нахождения списка студентов, получивших 2 по логпроге (ниже по остальным 5 предметам)
    findall(X, student(_, X, [grade('LP',2),grade('MTH',_),grade('FP',_),grade('INF',_),grade('ENG',_),grade('PSY',_)]), L).

mth_grade2_list(L) :-
    findall(X, student(_, X, [grade('LP',_),grade('MTH',2),grade('FP',_),grade('INF',_),grade('ENG',_),grade('PSY',_)]), L).

fp_grade2_list(L) :-
    findall(X, student(_, X, [grade('LP',_),grade('MTH',_),grade('FP',2),grade('INF',_),grade('ENG',_),grade('PSY',_)]), L).

inf_grade2_list(L) :-
    findall(X, student(_, X, [grade('LP',_),grade('MTH',_),grade('FP',_),grade('INF',2),grade('ENG',_),grade('PSY',_)]), L).

eng_grade2_list(L) :-
    findall(X, student(_, X, [grade('LP',_),grade('MTH',_),grade('FP',_),grade('INF',_),grade('ENG',2),grade('PSY',_)]), L).

psy_grade2_list(L) :-
    findall(X, student(_, X, [grade('LP',_),grade('MTH',_),grade('FP',_),grade('INF',_),grade('ENG',_),grade('PSY',2)]), L).

task2() :- % предикат вывода списков несдавших (предикат, выполняющий второе подзадание из трёх)
    write("\nStudents haven't passed logical programming:\n"),
    lp_grade2_list(A),
    print_list(A),
    write("\nStudents haven't passed mathematical analysis:\n"),
    mth_grade2_list(B),
    print_list(B),
    write("\nStudents haven't passed functional programming:\n"),
    fp_grade2_list(C),
    print_list(C),
    write("\nStudents haven't passed informatics:\n"),
    inf_grade2_list(D),
    print_list(D),
    write("\nStudents haven't passed English:\n"),
    eng_grade2_list(E),
    print_list(E),
    write("\nStudents haven't passed psychology:\n"),
    psy_grade2_list(F),
    print_list(F).


% 3) Найти количество не сдавших студентов в каждой из групп

delete([], _Elem, []):-!. % предикат удаления элемента из списка
delete([Elem|Tail], Elem, ResultTail):-
   delete(Tail, Elem, ResultTail), !.
delete([Head|Tail], Elem, [Head|ResultTail]):-
   delete(Tail, Elem, ResultTail).

delete_list([], _, []). % предикат удаления всех элементов списка 2 из списка 1
delete_list(L, [], L).
delete_list(L, [H2|T2], Res) :-
    delete(L, H2, Res2),
    delete_list(Res2, T2, Res).

group_amount(Num, Am) :- % предикат нахождения числа сдавших все экзамены в группе
    group_list(Num, L),
    lp_grade2_list(LP),
    mth_grade2_list(MTH),
    fp_grade2_list(FP),
    inf_grade2_list(INF),
    eng_grade2_list(ENG),
    psy_grade2_list(PSY),
    append1(LP, MTH, A),
    append1(FP, INF, B),
    append1(ENG, PSY, C),
    append1(A, B, D),
    append1(C, D, N),
    delete_list(L, N, R),
    length(R, Am).

task3() :- % предикат, выполняющий третье подзадание из трёх
    write("\nAmount of students haven't passed in 101 is "),
    group_amount(101, Am1),
    group_list(101, L1),
    length(L1, A1),
    X is A1 - Am1,
    write(X),
    write("\nAmount of students haven't passed in 102 is "),
    group_amount(102, Am2),
    group_list(102, L2),
    length(L2, A2),
    Y is A2 - Am2,
    write(Y),
    write("\nAmount of students haven't passed in 103 is "),
    group_amount(103, Am3),
    group_list(103, L3),
    length(L3, A3),
    Z is A3 - Am3,
    write(Z),
    write("\nAmount of students haven't passed in 104 is "),
    group_amount(104, Am4),
    group_list(104, L4),
    length(L4, A4),
    W is A4 - Am4,
    write(W),
    write('\n').


/* Tests:
?- task1().
?- task2().
?- task3().
*/

```
### Список реализованных предикатов:  
* `print_list([])` — предикат вывода элементов списка
* `append1([], List2, List2)` — предикат объединения списков
* `group_grades_list(Num, L)` — предикат получения списка со всеми оценками группы
* `sum_list([],0)` — предикат суммы списка
* `group_grades_sum(Num, Sum)` — предикат нахождения суммы оценок группы
* `group_grades_av(Num, Av)` — предикат нахождения средней оценки группы
* `all_groups_grades_av()` — предикат вывода средних оценок всех групп
* `group_list(Num, L)` — предикат нахождения списка группы
* `group_table()` — предикат вывода списков всех групп
* `task1()` — предикат, выполняющий первое подзадание из трёх
* `task2()` — предикат вывода списков несдавших (предикат, выполняющий второе подзадание из трёх)
* `delete([], _Elem, [])` — предикат удаления элемента из списка
* `delete_list([], _, [])` — предикат удаления всех элементов списка 2 из списка 1
* `group_amount(Num, Am)` — предикат нахождения числа сдавших все экзамены в группе
* `task3()` — предикат, выполняющий третье подзадание из трёх

### Результат работы предикатов:
```
3 ?- task1().

Group 101:
Петровский
Сидоров
Мышин
Безумников
Густобуквенникова

Group 102:
Петров
Ивановский
Биткоинов
Шарпин
Эксель
Текстописов
Криптовалютников
Азурин
Круглотличников

Group 103:
Сидоркин
Эфиркина
Сиплюсплюсов
Программиро
Клавиатурникова
Решетников
Текстописова
Вебсервисов

Group 104:
Иванов
Запорожцев
Джаво
Фулл
Круглосчиталкин
Блокчейнис

Average of 101 is 4
Average of 102 is 3.814814814814815
Average of 103 is 3.7708333333333335
Average of 104 is 3.888888888888889
true.

3 ?- task2().

Students haven't passed logical programming:

Students haven't passed mathematical analysis:
Клавиатурникова
Блокчейнис
Азурин

Students haven't passed functional programming:
Сидоркин
Мышин
Шарпин

Students haven't passed informatics:
Клавиатурникова
Текстописов
Блокчейнис

Students haven't passed English:
Петровский
Круглосчиталкин

Students haven't passed psychology:
Запорожцев
Сидоров
Мышин
true.

4 ?- task3().

Amount of students haven't passed in 101 is 3
Amount of students haven't passed in 102 is 3
Amount of students haven't passed in 103 is 2
Amount of students haven't passed in 104 is 3
true .
```

Основное преимущество реляционного представления заключается в упорядоченности и, как следствие, в иерархизме данных. Одним из  немногих недостатков является громоздкость данных. 
Преимущество представления three.pl: все данные о студенте (группа, имя, оценки) собраны в одном месте. Недостаток: необходимость обработки лишних данных в некоторых предикатах.

## Выводы

Выполняя данную лабораторную работу, я познакомился с концепцией декларативного пограммирования в общем и со списками в прологе в частности и сделал вывод, что в логическом программировании в первую очередь важно понимать, что из себя представляет ответ на поставленную задачу, а не процесс решения, как это принято в императивном программировании. Таким образом, выполняя лабораторную работу, я убедился, что средствами логического программирования можно эффективно работать с базами данных.
