:- set_prolog_flag(encoding, utf8).

agent('Даша').

object('шоколад').
object('деньги').

verbform('любить', 'любит').
verbform('лежать', 'лежат').

loc('').

condition(W, CW, T, CW:T:L) :- member(W, L).

getElem([],_,_).
getElem([H|T],T,H).

split(Q, W, Verb, Ad) :-
    getElem(Q, Q1, W), 
    getElem(Q1, Q2, Verb),
    getElem(Q2,_,Ad).

find_agent2(W,S) :- W == 'Кто'; agent(S).
find_agent1(W) :- W == 'Кто'.

find_object2(W,S) :- W == 'Что'; object(S).
find_object1(W) :- W == 'Что'.

find_loc(W,S) :- W == 'Где', loc(S).
find_loc1(W) :- W == 'Где'.
find_loc2(S) :- loc(S).

an_q(Q, _) :- 
    write("X="),
    split(Q, W, V1, S), 
    verbform(V, V1), write(V), 
    (
        find_agent2(W,S), !, write("(agent("), 
        (
            find_agent1(W), !, write("Y),"); write(S), write("),")
        ), 
        (
            find_loc(W,S), !, write("loc("),
            (
                find_loc1(W), !, write("Y))"); write(S), write("))")
            );
            write("object("),
            (
                find_object1(W), !, write("Y))");
                write(S),write("))")
            )
        );
        write("(object("),
        (
            find_object1(W), !, write("Y),"); write(S), write("),")
        ), 
        (
            write("loc("),
            (
                find_loc1(W), !, write("X))"); write(S), write("))")
            )
        )
    ).
