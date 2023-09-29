SELECT driver_second_name, driver_name, driver_third_name
FROM "Driver" INNER JOIN "Voyage" USING (driver_id)
WHERE point_end = 'Белгород';

SELECT voyage_id, driver_id, point_begin, point_end, date_begin, date_end
FROM "Voyage"
ORDER BY date_begin DESC;

SELECT vehicle_sigh
FROM "Vehicle" INNER JOIN "Types" USING (type_id)
WHERE type_name = 'truck';


