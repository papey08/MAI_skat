import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.neural_network import MLPClassifier
from sklearn.tree import DecisionTreeClassifier
from sklearn.metrics import accuracy_score, precision_score, recall_score, f1_score
import joblib
from sklearn.tree import export_graphviz
import graphviz
from IPython.display import Image
from sklearn.model_selection import GridSearchCV

param_grid = {
    'max_depth': [3, 4, 5, 10, 15],
    'min_samples_leaf': [1, 2, 3, 4],
    'max_leaf_nodes': [10, 20, 35, 50]
}

df = pd.read_csv('leisure2.csv')
#df = df.drop_duplicates()

feature_names = ['area', 'duration', 'budget', 'time', 'type']

X = df[['area', 'duration', 'budget', 'time', 'type']].values
y = df['place'].values

X_train, X_test, y_train, y_test = train_test_split(X, y, test_size = 0.1, random_state = 1) #random_state = 1

#model = DecisionTreeClassifier()

model = DecisionTreeClassifier(max_depth=10, min_samples_leaf=2, max_leaf_nodes=25)
model.fit(X_train, y_train)

joblib.dump(model, 'derevo.pkl')
print("Модель сохранена в файл derevo.pkl")

y_pred = model.predict(X_test)
print("accuracy:", accuracy_score(y_test, y_pred))



# Добавление параметра class_names
dot_file = export_graphviz(
    model,
    feature_names=feature_names,
    class_names=model.classes_.astype(str),  # Указываем названия классов
    filled=True,  # Заливка цветом для улучшения восприятия
    rounded=True,  # Скругленные узлы
    special_characters=True  # Специальные символы, если используются
)

#dot_file = export_graphviz(model, feature_names=feature_names)
graph = graphviz.Source(dot_file)
graph.render(filename='tree', format='png', cleanup=True)


area = int(input("Введите, где бы Вы хотели отдохнуть: улица - 1, помещение - 2 "))
duration = int(input("Введите время отдыха: 1 час = 1, 3 часа = 3, 6 часов = 6"))
budget = int(input("Введите ваш бюджет: до 1000 рублей = 1000, больше 1000 = 10000"))
time = int(input("Введите время суток: утро = 1, день = 2, вечер = 3, ночь = 4"))
type_l = int(input("Введите тип отдыха: активный = 1, пассивный = 2"))
answer = model.predict([[area, duration, budget, time, type_l]])
if answer == 1:
    print("Прогулка в парке или на набережной.")
elif answer == 2:
    print("Велосипед, самокат или лыжи, если лежит снег.")
elif answer == 3:
    print("Экскурсия или городской квест.")
elif answer == 4:
    print("Пикник или шашлыки.")
elif answer == 5:
    print("Кинотеатр.")
elif answer == 6:
    print("Музей, выставка или галерея.")
elif answer == 7:
    print("Тир или боулинг.")
elif answer == 8:
    print("Тренажёрный зал или фитнес-клуб.")
elif answer == 9:
    print("Бар или кафе.")
    
