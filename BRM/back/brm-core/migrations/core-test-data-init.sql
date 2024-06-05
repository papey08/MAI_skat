INSERT INTO "companies" ("name", "description", "industry", "owner_id", "rating", "creation_date", "is_deleted")
VALUES
    ('ООО «Драйв»', 'Перевозка любого оборудования на любые расстояния (в пределах России)', 3, 1, 4.2, '2022.03.20', false),
    ('ООО «Банкет»', 'Компания по организации банкетов и деловых обедов', 5, 6, 4.9, '2022.03.21', false),
    ('ООО «Облако»', 'Хостинг-провайдер', 1, 9, 2.1, '2022.09.01', false);

ALTER TABLE "employees"
ALTER COLUMN "image_url" TYPE VARCHAR(500);

INSERT INTO "employees" ("company_id", "first_name", "second_name", "email", "job_title", "department", "image_url", "creation_date", "is_deleted")
VALUES
    (1, 'Григорий', 'Панов', 'test_account01@yandex.ru', 'Генеральный директор', '', '', '2022.03.20', false),
    (1, 'Макар', 'Тимофеев', 'test_account02@yandex.ru', 'Менеджер по продажам', 'Продажи', '', '2022.04.20', false),
    (1, 'Артём', 'Кузнецов', 'test_account03@yandex.ru', 'Менеджер по продажам', 'Продажи', '', '2022.04.20', false),
    (1, 'Елизавета', 'Романова', 'test_account04@yandex.ru', 'Менеджер по продажам', 'Продажи', '', '2022.04.20', false),
    (1, 'Алиса', 'Шарова', 'test_account05@yandex.ru', 'Менеджер по продажам', 'Продажи', '', '2022.04.20', false),
    (2, 'Виктор', 'Антонов', 'test_account06@yandex.ru', 'Владелец', '', '', '2022.03.21', false),
    (2, 'Наталья', 'Полякова', 'test_account07@yandex.ru', 'Менеджер по продажам', 'Продажи', '', '2022.03.25', false),
    (2, 'Мария', 'Захарова', 'test_account08@yandex.ru', 'Менеджер по продажам', 'Продажи', '', '2022.03.25', false),
    (3, 'Артемий', 'Белоусов', 'test_account09@yandex.ru', 'Директор', '', '', '2022.09.01', false),
    (3, 'Ярослав', 'Агеев', 'test_account10@yandex.ru', 'Менеджер по продажам', 'Продажи', '', '2022.09.01', false);

INSERT INTO "contacts" ("owner_id", "employee_id", "notes", "creation_date", "is_deleted")
VALUES
    (1, 2, '', '2022.04.20', false),
    (1, 3, '', '2022.04.20', false),
    (1, 4, '', '2022.04.20', false),
    (1, 5, '', '2022.04.20', false),
    (1, 6, '', '2022.03.21', false),
    (1, 7, '', '2022.03.25', false),
    (1, 8, '', '2022.03.25', false),
    (1, 9, '', '2022.09.01', false),
    (2, 1, '', '2022.04.21', false),
    (2, 7, '', '2022.03.26', false),
    (7, 2, '', '2022.03.26', false),
    (9, 10, '', '2022.09.02', false),
    (10, 9, '', '2022.09.02', false),
    (6, 9, '', '2022.09.03', false),
    (9, 6, '', '2022.09.03', false),
    (6, 1, '', '2022.03.22', false);
