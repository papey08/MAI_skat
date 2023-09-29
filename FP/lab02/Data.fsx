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
