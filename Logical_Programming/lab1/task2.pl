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
