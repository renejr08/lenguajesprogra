%Ejercicio 1
sumlist([], 0).
sumlist([H|T], S) :- sumlist(T, S1), S is S1 + H.

%Ejercicio 2
subconj([], _).
subconj([X|XS], S) :- member(X, S), subconj(XS, S).

%Ejercicio 3
aplanar([], []).
aplanar([H|T], L2) :- is_list(H), aplanar(H, H1), aplanar(T, T1), append(H1, T1, L2).
aplanar([H|T], [H|T1]) :- \+ is_list(H), aplanar(T, T1).

%Ejercicio 4
partir([], _, [], []).
partir([H|T], Umbral, [H|Menores], Mayores) :- H =< Umbral, partir(T, Umbral, Menores, Mayores).
partir([H|T], Umbral, Menores, [H|Mayores]) :- H > Umbral, partir(T, Umbral, Menores, Mayores).

%Ejercicio 5
contiene_subcadena(_, [], []).
contiene_subcadena(Subcadena, [X|XS], [X|Filtradas]) :- sub_atom(X, _, _, _, Subcadena), contiene_subcadena(Subcadena, XS, Filtradas).
contiene_subcadena(Subcadena, [_|XS], Filtradas) :- contiene_subcadena(Subcadena, XS, Filtradas).

%Ejercicio 6
sub_cadenas(Subcadena, ListaCadenas, Filtradas) :- contiene_subcadena(Subcadena, ListaCadenas, Filtradas).
