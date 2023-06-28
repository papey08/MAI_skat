# Отчет по лабораторной работе
## по курсу "Функциональное программирование на Прологе"

### Студенты: 

| ФИО           | Группа      | Роль в проекте                     | Оценка       |
|---------------|-------------|------------------------------------|--------------|
| Борисов Ян    | М8О-308Б-20 | Парсер, обработка ошибок, отчёт    |              |
| Зубко Дмитрий | М8О-308Б-20 | Синтаксическое дерево, функционал  |              |
| Попов Матвей  | М8О-308Б-20 | Архитектура языка, отчёт           |              |


> *Комментарии проверяющих (обратите внимание, что более подробные комментарии возможны непосредственно в репозитории по тексту программы)*

## Пример программ на языке программирования "Чич"

### hello.ch

Для запуска:

```
$ go run cmd/main.go examples/hello.ch
```

Программа:

```
println("Hello World")

```

Результат:

```
Hello World
```

### calc.ch

Для запуска:

```
$ go run cmd/main.go examples/calc.ch
```

Программа:

```
a = 5
b = 6

c = a + b
println(c)

d = a - b
println(d)

e = a * b
println(e)

f = e / b
println(f)

g = e % c
println(g)

h = (a + b) * c - d % g
println(h)

```

Результат:

```
11
-1
30
5
8
122
```

### cond.ch

Для запуска:

```
$ go run cmd/main.go examples/cond.ch
```

Программа:

```
year = 2023
if year == 2023 {
	println("2023")
}

year = 2020
if year >= 2023 {
    println("more or equal")
} else {
    println("less")
}

year = 2025
if year <= 2023 {
    println("less of equal")
} else {
    println("more")
}

year = 2020
if year > 2023 {
    println("more")
} else {
    println("less or equal")
}

year = 2025
if year < 2023 {
    println("less")
} else {
    println("more or equal")
}

year = 2020
if year > 2000 && year < 2100 {
    println("21st century")
}

year = 2020
if (year == 2016 || year == 2020 || year == 2024) && (year != 2100) {
    println("leap year")
}

```

Результат:

```
2023
less
more
less or equal
more or equal
21st century
leap year
```

### list.ch

Для запуска:

```
$ go run cmd/main.go examples/list.ch
```

Программа:

```
lst = [0, 1, 2, 3, 4]

println(lst)
println(len(lst))

```

Результат:

```
[0, 1, 2, 3, 4]
5
```

### func.ch

Для запуска:

```
$ go run cmd/main.go examples/func.ch
```

Программа:

```
min = chich(a, b) {
    if a < b {
        return a
    }
    b
}

max = chich(a, b) {
    if a > b {
        return a
    }
    b
}

comp = chich(a, b, f) {
    f(a, b)
}

println(comp(1, 2, min))
println(comp(1, 2, max))

```

Результат:

```
1
2
```

### loop.ch

Для запуска:

```
$ go run cmd/main.go examples/loop.ch
```

Программа:

```
loop = chich(times, function) {
	if times > 0 {
		function()
		loop(times-1, function)
	}
}

loop(5, chich() { println("Hello World") })

```

Результат:

```
Hello World
Hello World
Hello World
Hello World
Hello World
```

### rec.ch

Для запуска:

```
$ go run cmd/main.go examples/rec.ch
```

Программа:

```
fib = chich(n) {
	if n < 2 {
		return n
	}
	fib(n-1) + fib(n-2)
}

println(fib(4))
println(fib(5))
println(fib(6))

_tail_fib = chich(n, acc1, acc2) {
    if n < 2 {
        return acc1
    }
    _tail_fib(n-1, acc1 + acc2, acc1)
}

tail_fib = chich(n) {
    _tail_fib(n, 1, 0)
}

println(tail_fib(4))
println(tail_fib(5))
println(tail_fib(6))

```

Результат:

```
3
5
8
3
5
8
```

### error.ch

Для запуска:

```
$ go run cmd/main.go examples/error.ch
```

Программа:

```
a = 5
b = 6

println(a + b + c)

```

Результат:

```
CHICH! error in file examples/error.ch at line 4:
    println(a + b + c)
                    ^
undefined variable c
```

## Синтаксическое дерево

Опишите, как представляется в программе синтаксическое дерево

## Интерпретатор

У нас получился транслируемый язык программирования. При подаче в него файла с 
программой на языке "Чич" программа сначала транслируется в байт-код, затем 
этот байт-код выполняется виртуальной машиной языка "Чич". Всё это происходит в 
рамках единственного выполнения программы, то есть в плане выполнения программ 
язык программирования "Чич" похож на интерпретируемые языки программирования, 
хотя в теории возможно сохранение файлов с байт-кодом и последующее их 
выполнение без транслирования.

## Какие фишки вы реализовали

* [ ] Именованные переменные
* [x] Рекурсия
* [ ] Ленивые вычисления
* [x] Функции
* [ ] Замыкания

## Getting Started with "Chich"

В языке программирования "Чич" реализованы два типа данных: целое число и 
строка. 

Над целыми числами возможны все основные операции: 

* сложение `+`
* вычитание `-`
* умножение `*`
* целочисленное деление `/`
* взятие остатка `%`

Синтаксис инициализации переменной и присваивания ей значения соответствует 
такому в языках Python и Ruby, блоки кода задаются фигурными скобками, как во 
всех си-подобных языках. В языке "Чич" также есть условные операторы.

Основной элемент любого функционального языка программирования — это функция. 
Язык программирования "Чич" — не исключение! Функции в этом языке создаются с 
помощью ключевого слова `chich`, могут присваиваться как переменные и 
передаваться в качестве аргументов другим функциям.
