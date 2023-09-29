:- set_prolog_flag(encoding, utf8).

:-['familytree.pl'].

remove([], _Elem, []). % предикат удаления элемента из списка
remove([Elem|T], Elem, TR) :- remove(T, Elem, TR), !.
remove([H|T], Elem, [H|TR]) :- remove(T, Elem, TR).

print_list([]). % предикат вывода элементов списка
print_list([H|T]) :-
    write(H),
    write('\n'),
    print_list(T).

freefromlist([H|_], H).

findfather(X, Y) :-
    findall(F, parents(X, F, _), L),
    freefromlist(L, Y).

findmother(X, Y) :-
    findall(M, parents(X, _, M), L),
    freefromlist(L, Y).

findchild(X, Y) :-
    findall(C, parents(C, X, _), L1),
    findall(D, parents(D, _, X), L2),
    append(L1, L2, Y).

findbrosis(X, Y) :-
    findfather(X, F),
    findchild(F, R1),
    remove(R1, X, Y).

findunclesaunts(X, Y) :-
    findfather(X, F),
    findbrosis(F, L1),
    findmother(X, M),
    findbrosis(M, L2),
    append(L1, L2, Y).

printcousins([H]) :-
    findchild(H, L),
    print_list(L).
printcousins([H|T]) :-
    findchild(H, L),
    print_list(L),
    printcousins(T).

findcousins(X) :-
    findunclesaunts(X, Y),
    printcousins(Y).

relative(brother/sister, X, Y) :-
    parents(X, F, M),
    parents(Y, F, M).

relative(father, X, Y) :-
    parents(X, Y, _).

relative(mother, X, Y) :-
    parents(X, _, Y).

relative(grandmother, X, Y) :-
    parents(X, F, _),
    parents(F, _, Y).

relative(grandmother, X, Y) :-
    parents(X, _, M),
    parents(M, _, Y).

relative(grandfather, X, Y) :-
    parents(X, F, _),
    parents(F, Y, _).

relative(grandfather, X, Y) :-
    parents(X, _, F),
    parents(F, Y, _).

relative(child, X, Y) :-
    relative(father, Y, X); relative(mother, Y, X).

relative(grandchild, X, Y) :-
    relative(grandfather, Y, X); relative(grandmother, Y, X).
