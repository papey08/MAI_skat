CREATE TABLE "Driver" (
    "driver_id" SERIAL PRIMARY KEY NOT NULL,
    "driver_second_name" VARCHAR(30) NOT NULL,
    "driver_name" VARCHAR(30) NOT NULL,
    "driver_third_name" VARCHAR(30) NOT NULL,
    "driver_class" VARCHAR(30) NOT NULL,
    "vehicle_sigh" VARCHAR(9) NOT NULL
);

CREATE TABLE "Types" (
    "type_id" SERIAL PRIMARY KEY NOT NULL,
    "type_name" VARCHAR(30) NOT NULL,
    "class" VARCHAR(30) NOT NULL,
    "capacity" INTEGER NOT NULL,
    "price" FLOAT NOT NULL
);

CREATE TABLE "Vehicle" (
    "vehicle_sigh" VARCHAR(9) PRIMARY KEY NOT NULL,
    "model" VARCHAR(30) NOT NULL,
    "type_id" INTEGER NOT NULL,
    "price_coeff" FLOAT NOT NULL
);

CREATE TABLE "Voyage" (
    "voyage_id" SERIAL PRIMARY KEY NOT NULL,
    "driver_id" INTEGER NOT NULL,
    "point_begin" VARCHAR(50) NOT NULL,
    "point_end" VARCHAR(50) NOT NULL,
    "date_begin" DATE NOT NULL,
    "date_end" DATE NOT NULL
);

ALTER TABLE "Voyage"
    ADD FOREIGN KEY ("driver_id")
        REFERENCES "Driver"("driver_id");

ALTER TABLE "Driver"
    ADD FOREIGN KEY ("vehicle_sigh")
        REFERENCES "Vehicle"("vehicle_sigh");

ALTER TABLE "Vehicle"
    ADD FOREIGN KEY ("type_id")
        REFERENCES "Types"("type_id");


INSERT INTO "Types" ("type_name", "class", "capacity", "price")
VALUES
    ('truck', 'B', 4, 75.0),
    ('van', 'B+', 2, 60.0),
    ('car', 'B+', 0, 0.0);

INSERT INTO "Vehicle" (vehicle_sigh, model, type_id, price_coeff)
VALUES
    ('ДР985У95', 'VAZ', 1, 0.9),
    ('СО029Л31', 'VAZ', 1, 0.85),
    ('ЛП069Г23', 'GAZEL', 2, 0.6),
    ('КГ045Е43', 'VAZ', 1, 0.8),
    ('ФА666Н02', 'VAZ', 1, 0.8),
    ('СВ033Г67', 'GAZEL', 2, 0.7);

INSERT INTO "Driver" (driver_second_name, driver_name, driver_third_name, driver_class, vehicle_sigh)
VALUES
    ('Гослинг', 'Раян', 'Алексеевич', '9Б', 'ДР985У95'),
    ('Суляева', 'Алина', 'Игоревна', '5А', 'СО029Л31'),
    ('Дорджиев', 'Тимур', 'Батырович', '46447', 'ЛП069Г23'),
    ('Баталин', 'Дмитрий', 'Андреевич', '11Б', 'КГ045Е43'),
    ('Старцев', 'Иван', 'Романович', '208Б', 'КГ045Е43'),
    ('Шашков', 'Дмитрий', 'Дмитриевич', '9Б', 'ЛП069Г23'),
    ('Зайцев', 'Кирилл', 'Владимирович', 'БАЗА', 'ФА666Н02'),
    ('Васютинский', 'Вадим', 'Александрович', 'КРИНЖ', 'ФА666Н02'),
    ('Ядров', 'Артем', 'Леонидович', 'БАЗА', 'ДР985У95'),
    ('Котов', 'Дмитрий', 'Валерьевич', 'ТНКФ', 'СО029Л31'),
    ('Богуж', 'Владислав', 'Андреевич', 'ХЗ', 'СВ033Г67'),
    ('Яценко', 'Александр', 'Владимирович', 'ХЗ', 'СВ033Г67');

INSERT INTO "Voyage" (driver_id, point_begin, point_end, date_begin, date_end)
VALUES
    (2, 'Москва', 'Белгород', '29.04.2022', '29.04.2022'),
    (2, 'Белгород', 'Москва', '04.05.2022', '04.05.2022'),
    (4, 'Москва', 'Иваново', '28.04.2022', '28.04.2022'),
    (4, 'Иваново', 'Москва', '04.05.2022', '04.05.2022'),
    (7, 'Москва', 'Донбасс', '24.02.2022', '25.02.2022'),
    (6, 'Москва', 'Хабаровск', '28.09.2022', '09.10.2022');


CREATE TABLE "Distance" (
    "distance" INTEGER NOT NULL,
    "point_begin" VARCHAR(50) NOT NULL,
    "point_end" VARCHAR(50) NOT NULL
);

INSERT INTO "Distance" (distance, point_begin, point_end)
VALUES
    (673, 'Москва', 'Белгород'),
    (673, 'Москва', 'Белгород'),
    (331, 'Иваново', 'Москва'),
    (331, 'Москва', 'Иваново'),
    (1156, 'Москва', 'Донбасс'),
    (1156, 'Донбасс', 'Москва'),
    (8410, 'Москва', 'Хабаровск'),
    (8410, 'Хабаровск', 'Москва');


CREATE TABLE "DriverA" (
    "driver_id" SERIAL PRIMARY KEY NOT NULL,
    "driver_second_name" VARCHAR(30) NOT NULL,
    "driver_name" VARCHAR(30) NOT NULL,
    "driver_third_name" VARCHAR(30) NOT NULL,
    "driver_class" VARCHAR(30) NOT NULL,
    "vehicle_sigh" VARCHAR(9) NOT NULL
);

ALTER TABLE "DriverA"
    ADD FOREIGN KEY("vehicle_sigh")
        REFERENCES "Vehicle"("vehicle_sigh");

INSERT INTO "DriverA" (driver_second_name, driver_name, driver_third_name, driver_class, vehicle_sigh)
VALUES
    ('Иванов', 'Сергей', 'Валерьевич', 'A', 'ДР985У95'),
    ('Шандрюк', 'Пётр', 'Николаевич', 'А', 'СО029Л31'),
    ('Битюков', 'Юра', 'Иванович', 'А', 'ЛП069Г23'),
    ('Беляков', 'Юрий', 'Александрович', 'А', 'КГ045Е43'),
    ('Беляев', 'Никита', 'Александрович', 'А', 'КГ045Е43'),
    ('Бортаковский', 'Александр', 'Сергеевич', 'А', 'ЛП069Г23');
