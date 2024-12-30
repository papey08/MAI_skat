SELECT departure_city, arrival_city, COUNT(*) FROM routes
WHERE departure_city = 'Москва' AND arrival_city = 'Санкт-Петербург'
GROUP BY departure_city, arrival_city;