SELECT DISTINCT departure_city, arrival_city
FROM routes r
JOIN aircrafts a ON r.aircraft_code = a.aircraft_code
WHERE a.aircraft_code = '773'
  AND departure_city > arrival_city
ORDER BY 1;
