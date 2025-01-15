import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.tree import DecisionTreeClassifier
from sklearn.metrics import accuracy_score, make_scorer
import joblib
from sklearn.tree import export_graphviz
import graphviz
from IPython.display import Image
from sklearn.model_selection import GridSearchCV
import numpy as np


def min_classes(y):
    '''
    Функция для нахождения классов в меньшинстве
    '''
    
    class_counts = pd.Series(y).value_counts()

    print("Количество экземпляров для каждого класса:")
    print(class_counts)

    total_samples = len(y)

    #Выбирает те, у которых количество экземпляров меньше, чем среднее количество экземпляров на класс.
    minority_classes = class_counts[class_counts < total_samples / len(class_counts)]

    print("\nКлассы в меньшинстве:")
    print(minority_classes)
    print(' ')


def custom_score(y_true, y_pred, required_classes=9):
    '''
    Данная функция проверяет число классов у созданной модели.
    Если их число не равно 9, то она возвращает точность, равную 0.
    '''
    unique_classes = np.unique(y_pred)  
    if len(unique_classes) == required_classes:  
        return accuracy_score(y_true, y_pred)  
    else:
        return 0  

    
def load_data(file_path):
    '''
    Загрузка файла .csv и его краткий анализ
    '''
    df = pd.read_csv(file_path)
    #df = df.drop_duplicates()
    y = df['place'].values
    min_classes(y)
    
    return df

def preprocess_data(df):
    '''
    Обработка файлов
    '''
    feature_names = ['area', 'duration', 'budget', 'time', 'type']
    X = df[feature_names].values
    y = df['place'].values
    return X, y, feature_names

def save_model(model):
    '''
    Сохранение модели в файл
    '''
    model = model.best_estimator_
    joblib.dump(model, 'derevo.pkl')
    print("Модель сохранена в файл derevo.pkl")

def train_model(X_train, y_train, X_test):
    '''
    Поиск подходящих гиперпараметров и обучение модели.
    '''
    param_grid = {
        'max_depth': [5, 10, 15, 20],
        'min_samples_leaf': [1, 2, 3, 4, 5],
        'max_leaf_nodes': [5, 10, 20, 25]}
    
    custom_scorer = make_scorer(custom_score, greater_is_better=True, required_classes=9)
    class_weights = {1: 2, 2: 2, 3: 4, 4: 3, 5: 2, 6 : 3, 7 : 3, 8 : 3, 9 : 2}
    
    model = GridSearchCV(DecisionTreeClassifier(class_weight=class_weights), param_grid = param_grid, cv=7, scoring=custom_scorer, refit = True)
    #model = GridSearchCV(DecisionTreeClassifier(), param_grid, cv=7, scoring='accuracy')
    model.fit(X_train, y_train)
   
    best_model = model.best_estimator_
    num_classes_model_can_recognize = len(best_model.classes_)

    print(f"Модель может распознавать {num_classes_model_can_recognize} классов.")
    print(f"Классы, которые модель может распознавать: {best_model.classes_}")
    
    print("Лучшие гиперпараметры:", model.best_params_)
    
    return model

def divide_data(X,y):
    '''
    Функция разбиения датасета на тренинговый и тестовый датасеты.
    '''
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size = 0.1, random_state = 1) #random_state = 1
    return X_train, X_test, y_train, y_test

def visualize(model, feature_names):
    '''
    Визуализация модели при помощи graphviz.
    '''
    model = model.best_estimator_
    dot_file = export_graphviz(
        model,
        feature_names=feature_names,
        class_names=model.classes_.astype(str),  # Названия классов
        filled=True,  # Заливка цветом для улучшения восприятия
        rounded=True,  # Скругленные узлы
        special_characters=True  # Специальные символы, если используются
    )

    graph = graphviz.Source(dot_file)
    graph.render(filename='tree', format='png', cleanup=True)

def testing_model(model, X_test, y_test):
    model = model.best_estimator_  
    y_pred = model.predict(X_test)
    print("accuracy:", accuracy_score(y_test, y_pred))
    

def get_user_input():
    '''
    Ввод данных пользователем.
    '''
    area = 0
    duration = 0
    budget = 0
    time = 0
    type_1 = 0
    err = 0
    try:
        area = int(input("Введите, где бы Вы хотели отдохнуть: улица - 1, помещение - 2\n"))
        if area not in [1, 2]:  
                raise ValueError("Некорректное значение для 'area'. Введите 1 для улицы или 2 для помещения.\n")
        duration = int(input("Введите время отдыха: 1 час = 1, 3 часа = 3, 6 часов = 6\n"))
        if duration not in [1, 3, 6]:  
                raise ValueError("Некорректное значение для 'duration'. Введите 1 для 1 часа, 3 для 3 часов, 6 для 6 часов.\n")
        budget = int(input("Введите ваш бюджет: до 1000 рублей = 1000, больше 1000 = 10000\n"))
        if budget not in [1000, 10000]:  
                raise ValueError("Некорректное значение для 'budget'. Введите 1000 до 1000 или 10000 если больше 1000.\n")
        time = int(input("Введите время суток: утро = 1, день = 2, вечер = 3, ночь = 4\n"))
        if time not in [1, 2, 3, 4]:
                raise ValueError("Некорректное значение для 'time'. 1 = утро, 2 = день, 3 = вечер, 4 = ночь.\n")
        type_1 = int(input("Введите тип отдыха: активный = 1, пассивный = 2\n"))
        if type_1 not in [1, 2]:
                raise ValueError("Некорректное значение для 'type_1'. 1 = активный, 2 = пассивный.\n")
    except ValueError as e:
        print(f"Ошибка: {e}. Попробуйте снова.")
        err = 1
           
    return [area, duration, budget, time, type_1], err

def predict_activity(model, user_input):
    '''
    Выбор наилучшего времяпроведения на основе ответов, поступивших от пользователя.
    '''
    answer = model.predict([user_input])[0]
    activities = {
        1: "Прогулка в парке или на набережной.",
        2: "Велосипед, самокат или лыжи, если лежит снег.",
        3: "Экскурсия или городской квест.",
        4: "Пикник или шашлыки.",
        5: "Кинотеатр.",
        6: "Музей, выставка или галерея.",
        7: "Тир или боулинг.",
        8: "Тренажёрный зал или фитнес-клуб.",
        9: "Бар или кафе."
    }
    return activities.get(answer, "Неизвестная активность.")

def main():

    df = load_data('leisure2.csv')
    X, y, feature_names = preprocess_data(df)
    
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.1, random_state=1)

    model = train_model(X_train, y_train, X_test)
    save_model(model)

    testing_model(model, X_test, y_test)
    
    visualize(model, feature_names)

    user_input, err = get_user_input()
    if err == 1:
        print('Неверные данные от пользователя, попробуйте запустить программу заново.')
    else:
        result = predict_activity(model, user_input)
        print(result)
    

if __name__ == "__main__":
    main()
    
