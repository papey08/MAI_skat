module Data =

    let studs = [
        ("Петров",101);
        ("Петровский",101);
        ("Иванов",103);
        ("Ивановский",102);
        ("Запорожцев",101);
        ("Сидоров",104);
        ("Сидоркин",101);
        ("Биткоинов",101);
        ("Эфиркина",103);
        ("Сиплюсплюсов",102);
        ("Программиро",104);
        ("Джаво",101);
        ("Клавиатурникова",103);
        ("Мышин",101);
        ("Фулл",103);
        ("Безумников",101);
        ("Шарпин",101);
        ("Круглосчиталкин",103);
        ("Решетников",101);
        ("Эксель",103);
        ("Текстописов",101);
        ("Текстописова",101);
        ("Густобуквенникова",101);
        ("Криптовалютников",102);
        ("Блокчейнис",102);
        ("Азурин",102);
        ("Вебсервисов",104);
        ("Круглотличников",103)
    ]

    let subjs = [
        ("LP","Логическое программирование");
        ("MTH","Математический анализ");
        ("FP","Функциональное программирование");
        ("INF","Информатика");
        ("ENG","Английский язык");
        ("PSY","Психология")
    ]

    let marks = [
        ("Петров","LP",4);
        ("Петров","MTH",4);
        ("Петров","FP",4);
        ("Петров","INF",4);
        ("Петров","ENG",4);
        ("Петров","PSY",3);
        ("Петровский","LP",5);
        ("Петровский","MTH",4);
        ("Петровский","FP",3);
        ("Петровский","INF",5);
        ("Петровский","ENG",3);
        ("Петровский","PSY",3);
        ("Иванов","LP",2);
        ("Иванов","MTH",5);
        ("Иванов","FP",3);
        ("Иванов","INF",5);
        ("Иванов","ENG",3);
        ("Иванов","PSY",5);
        ("Ивановский","LP",5);
        ("Ивановский","MTH",4);
        ("Ивановский","FP",4);
        ("Ивановский","INF",5);
        ("Ивановский","ENG",4);
        ("Ивановский","PSY",2);
        ("Запорожцев","LP",4);
        ("Запорожцев","MTH",4);
        ("Запорожцев","FP",4);
        ("Запорожцев","INF",3);
        ("Запорожцев","ENG",5);
        ("Запорожцев","PSY",4);
        ("Сидоров","LP",5);
        ("Сидоров","MTH",4);
        ("Сидоров","FP",4);
        ("Сидоров","INF",4);
        ("Сидоров","ENG",5);
        ("Сидоров","PSY",4);
        ("Сидоркин","LP",5);
        ("Сидоркин","MTH",4);
        ("Сидоркин","FP",3);
        ("Сидоркин","INF",3);
        ("Сидоркин","ENG",4);
        ("Сидоркин","PSY",3);
        ("Биткоинов","LP",3);
        ("Биткоинов","MTH",2);
        ("Биткоинов","FP",4);
        ("Биткоинов","INF",5);
        ("Биткоинов","ENG",4);
        ("Биткоинов","PSY",3);
        ("Эфиркина","LP",5);
        ("Эфиркина","MTH",5);
        ("Эфиркина","FP",5);
        ("Эфиркина","INF",3);
        ("Эфиркина","ENG",3);
        ("Эфиркина","PSY",2);
        ("Сиплюсплюсов","LP",3);
        ("Сиплюсплюсов","MTH",5);
        ("Сиплюсплюсов","FP",5);
        ("Сиплюсплюсов","INF",4);
        ("Сиплюсплюсов","ENG",3);
        ("Сиплюсплюсов","PSY",5);
        ("Программиро","LP",5);
        ("Программиро","MTH",4);
        ("Программиро","FP",5);
        ("Программиро","INF",4);
        ("Программиро","ENG",5);
        ("Программиро","PSY",4);
        ("Джаво","LP",4);
        ("Джаво","MTH",4);
        ("Джаво","FP",2);
        ("Джаво","INF",4);
        ("Джаво","ENG",3);
        ("Джаво","PSY",4);
        ("Клавиатурникова","LP",4);
        ("Клавиатурникова","MTH",4);
        ("Клавиатурникова","FP",5);
        ("Клавиатурникова","INF",4);
        ("Клавиатурникова","ENG",4);
        ("Клавиатурникова","PSY",4);
        ("Мышин","LP",4);
        ("Мышин","MTH",4);
        ("Мышин","FP",5);
        ("Мышин","INF",3);
        ("Мышин","ENG",3);
        ("Мышин","PSY",5);
        ("Фулл","LP",5);
        ("Фулл","MTH",4);
        ("Фулл","FP",3);
        ("Фулл","INF",4);
        ("Фулл","ENG",3);
        ("Фулл","PSY",5);
        ("Безумников","LP",4);
        ("Безумников","MTH",3);
        ("Безумников","FP",4);
        ("Безумников","INF",4);
        ("Безумников","ENG",4);
        ("Безумников","PSY",5);
        ("Шарпин","LP",3);
        ("Шарпин","MTH",4);
        ("Шарпин","FP",3);
        ("Шарпин","INF",4);
        ("Шарпин","ENG",3);
        ("Шарпин","PSY",5);
        ("Круглосчиталкин","LP",3);
        ("Круглосчиталкин","MTH",4);
        ("Круглосчиталкин","FP",5);
        ("Круглосчиталкин","INF",4);
        ("Круглосчиталкин","ENG",4);
        ("Круглосчиталкин","PSY",4);
        ("Решетников","LP",2);
        ("Решетников","MTH",4);
        ("Решетников","FP",3);
        ("Решетников","INF",4);
        ("Решетников","ENG",5);
        ("Решетников","PSY",4);
        ("Эксель","LP",5);
        ("Эксель","MTH",4);
        ("Эксель","FP",4);
        ("Эксель","INF",5);
        ("Эксель","ENG",5);
        ("Эксель","PSY",4);
        ("Текстописов","LP",4);
        ("Текстописов","MTH",3);
        ("Текстописов","FP",4);
        ("Текстописов","INF",2);
        ("Текстописов","ENG",3);
        ("Текстописов","PSY",5);
        ("Текстописова","LP",5);
        ("Текстописова","MTH",3);
        ("Текстописова","FP",3);
        ("Текстописова","INF",3);
        ("Текстописова","ENG",3);
        ("Текстописова","PSY",5);
        ("Густобуквенникова","LP",3);
        ("Густобуквенникова","MTH",5);
        ("Густобуквенникова","FP",4);
        ("Густобуквенникова","INF",2);
        ("Густобуквенникова","ENG",2);
        ("Густобуквенникова","PSY",5);
        ("Криптовалютников","LP",3);
        ("Криптовалютников","MTH",4);
        ("Криптовалютников","FP",5);
        ("Криптовалютников","INF",5);
        ("Криптовалютников","ENG",4);
        ("Криптовалютников","PSY",4);
        ("Блокчейнис","LP",5);
        ("Блокчейнис","MTH",5);
        ("Блокчейнис","FP",3);
        ("Блокчейнис","INF",5);
        ("Блокчейнис","ENG",4);
        ("Блокчейнис","PSY",5);
        ("Азурин","LP",4);
        ("Азурин","MTH",4);
        ("Азурин","FP",5);
        ("Азурин","INF",4);
        ("Азурин","ENG",3);
        ("Азурин","PSY",5);
        ("Вебсервисов","LP",4);
        ("Вебсервисов","MTH",2);
        ("Вебсервисов","FP",3);
        ("Вебсервисов","INF",5);
        ("Вебсервисов","ENG",4);
        ("Вебсервисов","PSY",4);
        ("Круглотличников","LP",3);
        ("Круглотличников","MTH",5);
        ("Круглотличников","FP",4);
        ("Круглотличников","INF",3);
        ("Круглотличников","ENG",3);
        ("Круглотличников","PSY",3)
    ]