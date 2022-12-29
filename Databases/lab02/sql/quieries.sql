-- Проверить, что рейсы каждого водителя не пересекаются по времени

SELECT driver_second_name, driver_name, date_begin
FROM "Voyage" INNER JOIN "Driver" USING(driver_id)
GROUP BY driver_second_name, driver_name, date_begin
HAVING COUNT(*) > 1
ORDER BY date_begin;


-- Выдать для каждого водителя среднюю длину маршрута

SELECT
    driver_second_name,
    driver_name,
    AVG(distance) AS average_distance
FROM
    "Driver"
        INNER JOIN "Voyage" V on "Driver".driver_id = V.driver_id
        INNER JOIN "Distance" D on V.point_begin = D.point_begin
        AND V.point_end = D.point_end
GROUP BY
    distance,
    driver_second_name,
    driver_name;


-- Рейсы из Москвы продолжительностью более 3-х часов

SELECT * FROM "Voyage"
WHERE date_end - date_begin <= 3 AND point_begin = 'Москва';

SELECT vehicle_sigh, model FROM "Vehicle"
    INNER JOIN "Driver" USING(vehicle_sigh)
    INNER JOIN "Voyage" USING(driver_id)
WHERE date_end < now()::date - 7;


-- Водители, которые за сегодня проехали более 300 км

SELECT driver_second_name, driver_name
FROM "Voyage"
         INNER JOIN "Driver" D on "Voyage".driver_id = D.driver_id
         INNER JOIN "Distance" D2 on "Voyage".point_begin = D2.point_begin AND "Voyage".point_end = D2.point_end
WHERE (distance > 300) AND (date_begin = CURRENT_DATE);
