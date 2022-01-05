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