CREATE TABLE companies (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "description" VARCHAR(1000),
    "industry" INTEGER NOT NULL,
    "owner_id" INTEGER NOT NULL,
    "rating" FLOAT NOT NULL,
    "creation_date" DATE NOT NULL,
    "is_deleted" BOOLEAN NOT NULL
);

CREATE TABLE employees (
    "id" SERIAL PRIMARY KEY,
    "company_id" INTEGER NOT NULL,
    "first_name" VARCHAR(100) NOT NULL,
    "second_name" VARCHAR(100),
    "email" VARCHAR(100) UNIQUE NOT NULL,
    "job_title" VARCHAR(100),
    "department" VARCHAR(100),
    "image_url" VARCHAR(200),
    "creation_date" DATE NOT NULL,
    "is_deleted" BOOLEAN NOT NULL
);

CREATE TABLE contacts (
    "id" SERIAL PRIMARY KEY,
    "owner_id" INTEGER NOT NULL,
    "employee_id" INTEGER NOT NULL,
    "notes" VARCHAR(500),
    "creation_date" DATE NOT NULL,
    "is_deleted" BOOLEAN NOT NULL,
    UNIQUE ("owner_id", "employee_id")
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
