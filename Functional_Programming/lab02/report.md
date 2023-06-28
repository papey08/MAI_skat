# Отчёт по лабораторной работе №2

## Обработка списков, оббработка данных

## по курсу "Функциональное программирование"

### студент: Попов Матвей Романович, группа М8О-308Б-20, вариант 20

## Результат проверки

|Преподаватель|Дата|Оценка|
|-------------|----|------|
|Сошников Д.В.|    |      |


## Введение

Списки в F# немного отличаются от того, что обычно называют
списками в императивных языках программирования. Во превых,
списки в F# являются неизменяемыми, что придаёт некоторые 
трудности, зато обеспечивает безопасность. Также в F# любой 
список можно представить голову и хвост, причём в отличие от 
пролога можно взять несколько головных элементов. Работа со 
списками в F# немного напоминает работу с очередями (так как 
обход списка чаще всего делается через отделение головного 
элемента от остальной части) или массивами.

## Часть 1: Обработка списков

### Вариант 10: Проверка того, является ли список геометрической прогрессией.

### С использованием библиотечных функций

```(F#)
let isGeometricProgression1 (lst: float list) =
    match lst with
    | [] | [_] -> true
    | x::y::t -> 
        let q = y / x
        lst |> List.pairwise |> List.forall (fun (a, b) -> b / a = q)
```

### Рекурсивно, без использования библиотечных функций

```(F#)
let rec isGeometricProgression2 (lst: float list) = 
    match lst with
    | [] | [_] | [_;_] -> true
    | x::y::z::t -> 
        if y / x = z / y then
            isGeometricProgression2 ([y]@[z]@t)
        else
            false
```

### Рекурсивно, без использования библиотечных функций, с хвостовой рекурсией

```(F#)
let rec tailRec (acc: float) (lst: float list) =
    match lst with
    | [] | [_] -> true
    | x::y::t -> 
        let q = y / x
        if q = acc then
            tailRec q (y::t)
        else
            false

let isGeometricProgression3 (lst: float list) =
    match lst with
    | [] | [_] -> true
    | x::y::t -> tailRec (y / x) (y::t)
```

Фактически оба рекурсивных алгоритма содержат хвостовую рекурсию, так как в 
обоих случаях вызов рекурсивной функции является последней операцией в функции, 
в одну из функций был добавлен аккумулятор, чтобы они отличались.

## Часть 2: Обработка данных

### Вариант 20

Способ хранения данных: `One.fsx`

Task 3:
- Для каждого студента, напечатайте его среднюю оценку и сдал ли он сессию (все оценки >2)
- Для каждого предмета, напечатайте список двоечников
- Для каждой группы, найдите студента (или студентов) с максимальной суммарной оценкой

### Исходный код

```(F#)
// Part 2

// Load the data
#load "One.fsx"
open One

// Task 3.1
let averageMarks =
    Data.marks
    |> List.groupBy (fun (student, _, _) -> student)
    |> List.map (fun (student, values) ->
        let hasMark2 = values |> List.exists (fun (_, _, mark) -> mark = 2)
        let averageMark = List.averageBy (fun (_, _, mark) -> float mark) values
        student, averageMark, not hasMark2)

// Task 3.2
let studentsWithMarkTwo =
    Data.marks
    |> List.filter (fun (_, _, mark) -> mark = 2)
    |> List.groupBy (fun (_, subject, _) -> subject)
    |> List.map (fun (subject, values) ->
        let students = values |> List.map (fun (student, _, _) -> student)
        subject, students)

let findFullSubjectName (subject: string) =
    match List.tryFind (fun (k, _) -> k = subject) Data.subjs with
    | Some (_, value) -> value
    | _ -> ""

// Task 3.3
let maxSummaryMarks =
    Data.marks
    |> Seq.groupBy (fun (student, _, _) -> (student, Seq.pick (fun (s, g) -> 
        if s = student then Some g else None) Data.studs))
    |> Seq.map (fun ((student, group), marks) -> 
        (student, group, Seq.sumBy (fun (s, sub, mark) -> mark) marks))
    |> Seq.groupBy (fun (_, group, _) -> group)
    |> Seq.map (fun (group, students) -> 
        let maxStudent = Seq.maxBy (fun (_, _, summaryMark) -> summaryMark) students
        (group, maxStudent))
    |> Seq.map (fun (group, (student, _, _)) -> (group, student))
    |> List.ofSeq
    |> List.sortBy (fun (k, _) -> k)


printfn "Task 3.1: Для каждого студента напечатайте его среднюю оценку и сдал ли он сессию (все оценки >2)"
for (student, averageMark, hasMark2) in averageMarks do
    printfn "%s %f %b" student averageMark hasMark2
printf "\n\n\n"

printfn "Task 3.2: Для каждого предмета напечатайте список двоечников"
for (subject, students) in studentsWithMarkTwo do
    let fullSubjectName = findFullSubjectName subject
    printf "%s: " fullSubjectName
    for student in students do
        printf "%s " student
    printfn ""
printf "\n\n\n"

printfn "Task 3.3: Для каждой группы найдите студента (или студентов) с максимальной суммарной оценкой"
for (group, student) in maxSummaryMarks do
    printfn "%d: %s" group student

```

### Результат работы:

```
PS C:\Users\papey08\github-classroom\MAILabs-Edu-2023\fp-lab2-papey08> dotnet fsi Data.fsx
Task 3.1: Для каждого студента напечатайте его среднюю оценку и сдал ли он сессию (все оценки >2)
Петров 3.833333 true
Петровский 3.833333 true
Иванов 3.833333 false
Ивановский 4.000000 false
Запорожцев 4.000000 true
Сидоров 4.333333 true
Сидоркин 3.666667 true
Биткоинов 3.500000 false
Эфиркина 3.833333 false
Сиплюсплюсов 4.166667 true
Программиро 4.500000 true
Джаво 3.500000 false
Клавиатурникова 4.166667 true
Мышин 4.000000 true
Фулл 4.000000 true
Безумников 4.000000 true
Шарпин 3.666667 true
Круглосчиталкин 4.000000 true
Решетников 3.666667 false
Эксель 4.500000 true
Текстописов 3.500000 false
Текстописова 3.666667 true
Густобуквенникова 3.500000 false
Криптовалютников 4.166667 true
Блокчейнис 4.500000 true
Азурин 4.166667 true
Вебсервисов 3.666667 false
Круглотличников 3.500000 true



Task 3.2: Для каждого предмета напечатайте список двоечников
Логическое программирование: Иванов Решетников
Психология: Ивановский Эфиркина
Математический анализ: Биткоинов Вебсервисов
Функциональное программирование: Джаво
Информатика: Текстописов Густобуквенникова
Английский язык: Густобуквенникова



Task 3.3: Для каждой группы найдите студента (или студентов) с максимальной суммарной оценкой
101: Запорожцев
102: Блокчейнис
103: Эксель
104: Программиро
```

## Выводы

Выполняя данную лабораторную работу, я продолжил знакомство с концепцией 
функционального пограммирования в общем и со списками в F#. Таким образом, я 
убедился, что средствами функционального программирования можно эффективно 
работать с базами данных.
