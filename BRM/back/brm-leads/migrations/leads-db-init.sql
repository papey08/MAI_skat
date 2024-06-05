CREATE TABLE leads (
    "id" SERIAL PRIMARY KEY,
    "ad_id" INTEGER NOT NULL,
    "title" VARCHAR(200) NOT NULL,
    "description" VARCHAR(1000),
    "price" INTEGER NOT NULL,
    "status" INTEGER NOT NULL,
    "responsible" INTEGER NOT NULL,
    "company_id" INTEGER NOT NULL,
    "client_company" INTEGER NOT NULL,
    "client_employee" INTEGER NOT NULL,
    "creation_date" DATE NOT NULL,
    "is_deleted" BOOLEAN NOT NULL
);

CREATE TABLE statuses (
    "id" SERIAL,
    "name" VARCHAR(100) PRIMARY KEY
);

INSERT INTO statuses (name)
VALUES
    ('Новая сделка'),
    ('Установка контакта'),
    ('Обсуждение деталей'),
    ('Заключительные детали'),
    ('Завершено'),
    ('Отклонено');
