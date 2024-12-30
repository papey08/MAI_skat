SELECT city
FROM airports
WHERE city <>'Москва'
EXCEPT
SELECT arrival_city
FROM routes
WHERE departure_city ='Москва'
ORDER BY city;
