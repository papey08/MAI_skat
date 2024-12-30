# Лабораторная работа №3, Фундаментальные концепции ИИ
# Оптимизация гиперпараметра
# Попов Матвей, М8О-114СВ-24

# Для запуска docker-контейнера с postgresql:
# docker run --name postgres-optuna -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:15.5

# Для удаления контейнера:
# docker stop postgres-optuna
# docker rm postgres-optuna

import optuna
from optuna.visualization import (
    plot_optimization_history,
    plot_param_importances,
    plot_slice,
    plot_parallel_coordinate,
)
import sklearn.datasets
import sklearn.ensemble
import sklearn.model_selection
from sklearn.metrics import accuracy_score

def objective_classification(trial):
    """
    Задача: классификация
    Датасет: wine
    Метод: Random Forest
    """
    wine = sklearn.datasets.load_wine()
    X_train, X_test, y_train, y_test = sklearn.model_selection.train_test_split(
        wine.data, wine.target, test_size=0.25, random_state=42
    )

    n_estimators = trial.suggest_int("n_estimators", 10, 200)
    max_depth = trial.suggest_int("max_depth", 2, 32, log=True)
    min_samples_split = trial.suggest_int("min_samples_split", 2, 20)

    clf = sklearn.ensemble.RandomForestClassifier(
        n_estimators=n_estimators,
        max_depth=max_depth,
        min_samples_split=min_samples_split,
        random_state=42,
    )
    clf.fit(X_train, y_train)

    y_pred = clf.predict(X_test)
    accuracy = accuracy_score(y_test, y_pred)
    return accuracy

storage_url = "postgresql://postgres:postgres@localhost:5432/postgres"
study_name = "wine_classification"

study = optuna.create_study(
    study_name=study_name,
    storage=storage_url,
    direction="maximize",
    load_if_exists=False
)

study.optimize(objective_classification, n_trials=50)

print("best params:", study.best_params)
print("best value:", study.best_value)

plot_optimization_history(study).show()
plot_param_importances(study).show()
plot_slice(study).show()
plot_parallel_coordinate(study).show()
