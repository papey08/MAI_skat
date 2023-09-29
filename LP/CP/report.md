# Отчет по курсовому проекту
## по курсу "Логическое программирование"

### студент: Попов Матвей Романович

## Результат проверки

Вариант задания:

 - [x] стандартный, без NLP (на 3)
 - [ ] стандартный, с NLP (на 3-4)
 - [ ] продвинутый (на 3-5)
 
| Преподаватель     | Дата         |  Оценка       |
|-------------------|--------------|---------------|
| Сошников Д.В. |              |               |
| Левинская М.А.|              |               |

> *Комментарии проверяющих*

## Введение

В рамках выполнения курсового проекта я рассчитываю получить навыки работы с генеалогическими деревьями в формате GEDCOM и работы с соответствующим программным обеспечением, навыки создания так называемых парсеров (программ, преобразующих текст из одного формата в другой), а также улучшить навыки работы с предикатами на языке Prolog.

## Задание

1. Создать родословное дерево своего рода на несколько поколений (3-4) назад в стандартном формате GEDCOM.  
2. Преобразовать файл в формате GEDCOM в набор утверждений на языке Prolog с использованием предиката `parents(потомок, отец, мать)`.
3. Реализовать предикат проверки/поиска двоюродных братьев/сестёр

## Получение родословного дерева

Родословное дерево в формате GEDCOM было получено с помощью приложения [Gramps](https://gramps-project.org/blog/), для заполнения дерева использовались имена персонажей популярной мобильной видеоигры Brawl Stars.

## Конвертация родословного дерева

Для написания парсера был выбран язык программирования Ruby, так как в Ruby существует множество удобных инструментов для чтения/записи файлов, помимо этого Ruby обладает понятным синтаксисом и, как следствие, высокой читаемостью кода.
 
Исходный код файла `parser.ru`:
```
if File.exist?("familytree.ged")
    file = File.open("familytree.ged", "r")
    writefile = File.new("familytree.pl", "w+")
    idbase = []
    namebase = []
    namestr = ""
    lines = file.readlines
    lines.each do |x|
        if x[2] == '@' && x[3] == 'I'
            i = 4
            r = ""
            while x[i] != '@'
                r += x[i]
                i += 1
            end
            id = r.to_i
            idbase << id
        end
        if x[2] == 'N' && x.length > 6 && x[0] != '2'
            t = 0
            for i in (7...x.length) do
                if x[i + 1] == '/'
                    t = i
                    break
                end
                namestr << x[i]
            end
            if t != 0
                for j in (i + 3 ...x.length - 1) do
                    namestr << x[j]
                end
            end
            namebase << namestr
            namestr = ""
        end
        if x[2] == '@' && x[3] == 'F'
            break
        end
    end
    namebase.delete_at(0)
    base = {}
    i = 0
    idbase.each do |x|
        base[x] = namebase[i]
        i += 1
    end
    father = ""
    mother = ""
    child = ""
    namebase.clear
    idbase.clear
    lines.each do |x|
        if x[2] == 'H' && x[3] == 'U'
            i = 9
            r = ""
            while x[i] != '@'
                r += x[i]
                i += 1
            end
            id = r.to_i
            f = id
            father = base[f]
        end
        if x[2] == 'W' && x[3] == 'I' && x[4] == 'F' && x[5] == 'E'
            i = 9
            r = ""
            while x[i] != '@'
                r += x[i]
                i += 1
            end
            id = r.to_i
            m = id
            mother = base[m]
        end
        if x[2] == 'C' && x[3] == 'H' && x[4] == 'I' && x[5] == 'L'
            i = 9
            r = ""
            while x[i] != '@'
                r += x[i]
                i += 1
            end
            id = r.to_i
            ch = id
            child = base[ch]
            writefile.puts "parents('#{child}', '#{father}', '#{mother}')."
        end
    end
    lines.clear
    base.clear
    file.close
    writefile.close
else
    puts("File familytree.ged not found")
end
```

Первым делом производится проверка на наличие GEDCOM-файла familytree.ged в каталоге, далее в тексте этого файла будет производиться поиск элементов дерева, которые будут добавлены в хэш-таблицу (имени человека сопоставляется уникальный id), затем производится поиск связей различных элементов дерева, на основании которых файл familytree.pl заполнится предикатами.

## Предикат поиска родственника

Исходный код файла `familytree.pl`:
```
parents('Stu', 'Bull', 'Colt').
parents('Buzz', 'Bull', 'Colt').
parents('Darryl', 'Bull', 'Colt').
parents('Mortis', 'Shelly', 'Emz').
parents('Sprout', 'Shelly', 'Emz').
parents('Edgar', 'Shelly', 'Emz').
parents('Colette', '8-Bit', 'Barley').
parents('Spike', '8-Bit', 'Barley').
parents('Penny', 'Darryl', 'Sprout').
parents('Rosa', 'Darryl', 'Sprout').
parents('Griff', 'Edgar', 'Colette').
```
Исходный код файла `main.pl`:
```
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
    parents(X, F, M),
    parents(F, _, Y); parents(M, _, Y).

relative(grandfather, X, Y) :-
    parents(X, F, M),
    parents(F, Y, _); parents(M, Y, _).

relative(child, X, Y) :-
    relative(father, Y, X); relative(mother, Y, X).

relative(grandchild, X, Y) :-
    relative(grandfather, Y, X); relative(grandmother, Y, X).

```
`findcousins(X).` — пердикат, который выводит список двоюродных братьев и сестёр элемента x, сначала находятся оба родителя элемента x, затем через родителей родителей x находим братьев и сестёр элемента x (то есть его тёти и дяди), затем выводим их детей, которые и будут двоюродными братьями и сёстрами элемента x.

Пример использования предиката:
```
2 ?- findcousins('Griff').
Penny
Rosa
true .

3 ?- findcousins('Penny'). 
Griff
true .
```

## Определение степени родства

В данной программе реализованы предикаты определения отца, матери, ребёнка, брата, сестры, дедушки, бабушки, внуков. Все эти предикаты реализованы через стандартный пердикат `parents`.

Примеры использования предикатов:
```
2 ?- relative(child, 'Colette', X).
X = 'Griff'.

3 ?- relative(X, 'Buzz', 'Colt'). 
X = mother .

4 ?- relative(father, 'Buzz', X).      
X = 'Bull'.

5 ?- relative(brother/sister, 'Rosa', X).
X = 'Penny' .

6 ?- relative(grandmother, X, 'Colt').  
X = 'Penny' ;
X = 'Rosa'.

7 ?- relative(grandfather, 'Griff', X).
X = 'Shelly' ;
X = '8-Bit'.

8 ?- relative(grandchild, 'Emz', X).
X = 'Griff' ;
X = 'Penny' ;
X = 'Rosa'.
```

## Выводы

Таким образом, проделав лабораторную работу, я познакомился с новым типом файлов GEDCOM, научился читать и понимать содержимое этих файлов, а также применил свои навыки в программировании на языке Ruby для того, чтобы перевести содержимое файла GEDCOM в набор предикатов на языке Prolog. Помимо этого основной частью моей работы являлось закрепление навыков обращения с предикатами языка Prolog, в этом мне помогло решение прикладной задачи, связанной с различными родственными связями. Тем самым я познакомился с множеством новых концепций в логическом (и не только) программировании.