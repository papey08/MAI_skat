CREATE TABLE ads (
    "id" SERIAL PRIMARY KEY,
    "company_id" INTEGER NOT NULL,
    "title" VARCHAR(200) NOT NULL,
    "text" VARCHAR(1000) NOT NULL,
    "industry" INTEGER NOT NULL,
    "price" INTEGER NOT NULL,
    "image_url" VARCHAR(200) NOT NULL,
    "creation_date" DATE NOT NULL,
    "created_by" INTEGER NOT NULL,
    "responsible" INTEGER NOT NULL,
    "is_deleted" BOOLEAN
);

CREATE TABLE responses (
    "id" SERIAL PRIMARY KEY,
    "company_id" INTEGER,
    "employee_id" INTEGER,
    "ad_id" INTEGER,
    "creation_date" DATE
);

CREATE TABLE industries (
    "id" SERIAL,
    "name" VARCHAR(100) PRIMARY KEY
);

INSERT INTO "industries" (name)
VALUES
    ('Информационные технологии'),
    ('Юридические услуги'),
    ('Транспорт'),
    ('Промышленность'),
    ('Еда и напитки'),
    ('Одежда и обувь'),
    ('Развлечения'),
    ('Туризм'),
    ('Медицина'),
    ('Рестораны и бары'),
    ('Гостиницы');
