movelr([X|L], [Y|R], L, [X, Y|R]) :- not(X = Y).
movelr([X|L], [], L, [X]).

movetr([X|T], [Y|R], T, [X, Y|R]) :- not(X = Y).
movetr([X|T], [], T, [X]).

movelt([X|L], [Y|T], L, [X, Y|T]).
movelt([X|L], [], L, [X]).

move1([L, T, R], [LN, TN, RN]) :- 
    T = TN, movelr(L, R, LN, RN);
    L = LN, movetr(T, R, TN, RN);
    R = RN, movelt(L, T, LN, TN).

move([X|T], [Y, X|T]) :- 
    move1(X, Y).

bw([], [], []).
bw([X|TX], [Y|TY], [X, Y|T]) :- bw(TX, TY, T).

dig(0).
dig(X) :- dig(X1), X is X1 + 1.

change([w], [w], []).
change([b], [], [b]).
change([w|T], [w|W], B) :- change(T, W, B).
change([b|T], W, [b|B]) :- change(T, W, B).

sol([X|T], R) :- 
    change([X|T], W, B),
   ((X = w, bw(B, W, R)); 
    (X = b, bw(W, B, R))).

more([H | T], R) :- 
    H = R;
    more(T, R).

dsearch1([H | T], H, [H | T]).
dsearch1(P, G, R) :-
    move(P, P1),
    dsearch1(P1, G, R).

dsearch(X, R) :-
    get_time(T0),
    sol(X, P),
    dsearch1([[X, [], []]], [[], [], P], R1),
    reverse(R1, R),
    get_time(T1),
    T is T1 - T0,
    write(T).

disearch1([H | T], H, [H | T], _).
disearch1(P, G, R, C) :-
    C > 0,
    move(P, P1), 
    C1 is C - 1,
    disearch1(P1, G, R, C1).

disearch(X, R) :- 
    get_time(T0),
    sol(X, P),
    dig(N),
    disearch1([[X, [], []]], [[], [], P], R1, N), !,
    reverse(R1, R2), 
    more(R2, R),
    get_time(T1),
    T is T1 - T0,
    write(T).

wsearch1([[H | T] | _], H, [H | T]).
wsearch1([H | I], P, R) :-
   findall1(H, L),
   append(I, L, O),
   wsearch1(O, P, R).

wsearch2(X, R) :-
   sol(X, P),
   wsearch1([[[X, [], []]]], [[], [], P], R1),
   !,
   reverse(R1, R2),
   more(R2, R).

wsearch(X, R) :-
    get_time(T0),
    wsearch2(X, R),
    get_time(T1),
    T is T1 - T0,
    write(T).
    

findall1(P, R) :-
   findall(S, move(P, S), R);
   not(findall(S, move(P, S), R)), R = [].