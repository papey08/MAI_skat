[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-718a45dd9cf7e7f842a935f5ebbe5719a5e09af4491e668f4dbf3b35d5cca122.svg)](https://classroom.github.com/online_ide?assignment_repo_id=10795462&assignment_repo_type=AssignmentRepo)
# Лабораторная работа по курсу "Искусственный интеллект"
# Создание своего нейросетевого фреймворка

### Студенты:

| ФИО                       | Роль в проекте                                                               | Оценка |
|---------------------------|------------------------------------------------------------------------------|--------|
| Борисов Ян М8О-308Б-20    | Написал функции оптимизации и функции потерь, писал отчёт                    |        |
| Зубко Дмитрий М8О-308Б-20 | Написал инструменты для работы с датасетами и функции активации, писал отчёт |        |
| Попов Матвей М8О-308Б-20  | Разработал архитектуру проекта, тестрировал нейросеть, писал отчёт           |        |


> *Комментарии проверяющего*

### Задание

Реализовать свой нейросетевой фреймворк для обучения полносвязных нейросетей, который должен включать в себя следующие возможности:

1. Создание многослойной нейросети перечислением слоёв
2. Удобный набор функций для работы с данными и датасетами (map, minibatching, перемешивание и др.)
3. Несколько (не менее 3) алгоритмов оптимизации: SGD, Momentum SGD, Gradient Clipping и др.
4. Описание нескольких передаточных функций и функций потерь для решения задач классификации и регрессии.
5. Обучение нейросети "в несколько строк", при этом с гибкой возможностью конфигурирования
6. 2-3 примера использования нейросети на классических задачах (MNIST, Iris и др.)
7. Документация в виде файла README.md


### Документация

#### Инструменты для работы с датасетами

* `Load(path string) (training.Pairs, error)` — загрузить датасет из файла в csv-формате
* `Shuffle()` — перемешать датасет
* `Split(n float64) (first, second Pairs)` — разбить датасет на 2 в случайном порядке
* `SplitSize(size int) []Pairs` — получить датасеты заданного размера из оригинального
* `SplitN(n int) []Pairs` — получить ***n*** датасетов из оригинального

#### Реализованные алгоритмы оптимизации

* SGD
* Momentum SGD
* Adam

#### Функции потерь

* Binary Cross Entropy
* Cross Entropy
* Mean Squared

#### Функции активации

* Linear
* ReLU
* Sigmoid
* TanH

#### Пример обучения нейросети на датасете ***wines***

```go
package main

import (
	"ai_lab3/internal/activation"
	"ai_lab3/internal/dataset"
	"ai_lab3/internal/network"
	"ai_lab3/internal/training"
	"ai_lab3/internal/training/solver"
	"ai_lab3/internal/util"
	"ai_lab3/internal/weights"
)

func main() {
	data, err := dataset.Load("./wine.data")
	if err != nil {
		panic(err)
	}

	for i := range data {
		util.Standardize(data[i].Input)
	}
	data.Shuffle()

	n := network.NewNetwork(&network.Params{
		Inputs:       len(data[0].Input),
		LayoutConfig: []int{6, 3}, // network with 1 hidden layer of 6 nodes and an output layer of 3 nodes
		Activation:   activation.Tanh,
		Mode:         activation.MultiClass,
		Weight:       weights.NewNormal(1, 0),
		Bias:         true,
	})

	trainer := training.NewTrainer(solver.NewSGD(0.005, 0.5, 0.001), 200)
	_ = trainer.Train(n, data, data, 2000)
}
```

##### Результат

```text
Epochs          Elapsed         Loss CE         Accuracy        
---             ---             ---             ---
2000            617.6247ms      0.2072          0.92

```

##### Для запуска

```text
go run cmd/wines/main.go
```
