% USE UTF-8


% Стандартные предикаты:

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

/* Tests:
?- length([1, 2, 3], N).
?- member(1, [1, 2, 3]).
?- append([1, 2], [3, 4], L).
?- remove([1, 2, 3], 3, L).
?- permute([1, 2, 3], [2, 1, 3]).
?- sublist([1, 2], [1, 2, 3]). 
*/


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


% Task 1.2
% Вычисление минимального элемента

min([MinElem], MinElem). % предикат из задания
min([H|T], MinElem) :- 
    min(T, TMinElem), 
    TMinElem < H, 
    !, 
    MinElem = TMinElem; 
    MinElem = H.

/* Tests:
?- min([1, 2, 3], X).
*/
