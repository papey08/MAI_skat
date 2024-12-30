EXPLAIN ANALYZE
SELECT a.aircraft_code AS a_code,
       a.model,
       (SELECT count(r.aircraft_code)
        FROM routes r
        WHERE r.aircraft_code = a.aircraft_code) AS num_routes
FROM aircrafts a
GROUP BY 1, 2
ORDER BY 3 DESC;


EXPLAIN ANALYZE
SELECT a.aircraft_code AS a_code,
       a.model,
       count(r.aircraft_code) AS num_routes
FROM aircrafts a
         LEFT OUTER JOIN routes r ON r.aircraft_code = a.aircraft_code
GROUP BY 1, 2
ORDER BY 3 DESC;
