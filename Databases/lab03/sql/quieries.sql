-- Выборка

SELECT * FROM "Distance"
WHERE distance > 1000;


-- Переименование

SELECT driver_second_name AS Фамилия, driver_name AS Имя, driver_third_name AS Отчество
FROM "Driver";


-- Проекция

SELECT * FROM "Driver"
WHERE driver_class = '9Б';


-- Объединение

SELECT * FROM "Driver"
UNION SELECT * FROM "DriverA"
ORDER BY driver_id;


-- Пересечение

SELECT * FROM (
    SELECT * FROM "Driver"
    UNION SELECT * FROM "DriverA"
    ORDER BY driver_id) as "D*DA*"
WHERE vehicle_sigh = 'КГ045Е43';


-- Разность

SELECT V.vehicle_sigh
FROM "Vehicle" AS V
WHERE V.vehicle_sigh NOT IN (
    SELECT DISTINCT D.vehicle_sigh
    FROM "DriverA" AS D
);


-- Агрегирование

SELECT vehicle_sigh, COUNT(driver_id)
FROM (
    SELECT * FROM "Driver"
    UNION SELECT * FROM "DriverA"
) as "D*DA*"
GROUP BY vehicle_sigh;


-- Внутреннее соединение

SELECT driver_id, driver_second_name, driver_name, model
FROM (
    SELECT * FROM "Driver"
    UNION SELECT * FROM "DriverA"
) as "D*DA*" INNER JOIN "Vehicle" USING(vehicle_sigh)
ORDER BY driver_id;


-- Внешнее соединение

SELECT driver_id, driver_second_name, driver_name, model
FROM (
    SELECT * FROM "Driver"
    UNION SELECT * FROM "DriverA"
) as "D*DA*" INNER JOIN "Vehicle" USING(vehicle_sigh)
ORDER BY driver_id;


-- Декартово произведение

SELECT vehicle_sigh, type_name
FROM "Vehicle" CROSS JOIN "Types";

