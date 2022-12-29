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
