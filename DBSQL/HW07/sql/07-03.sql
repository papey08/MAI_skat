SELECT count( * )
FROM ticket_flights
WHERE fare_conditions ='Comfort';

SELECT count( * )
FROM ticket_flights
WHERE fare_conditions ='Business';

SELECT count( * )
FROM ticket_flights
WHERE fare_conditions ='Economy';

CREATE INDEX fare_conditions_index
ON ticket_flights (fare_conditions);

SELECT count( * )
FROM ticket_flights
WHERE fare_conditions ='Comfort';

SELECT count( * )
FROM ticket_flights
WHERE fare_conditions ='Business';

SELECT count( * )
FROM ticket_flights
WHERE fare_conditions ='Economy';

SELECT count(*)
FROM ticket_flights;
