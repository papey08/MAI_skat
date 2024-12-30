EXPLAIN WITH cities AS (
    SELECT DISTINCT city FROM airports
)
SELECT count(*)
FROM cities a1
JOIN cities a2 ON a1.city <> a2.city;
