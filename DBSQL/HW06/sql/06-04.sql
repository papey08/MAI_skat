CREATE TABLE seats_copy (LIKE seats INCLUDING ALL);

INSERT INTO seats_copy
SELECT * FROM seats;
i
INSERT INTO seats_copy (aircraft_code, seat_no, fare_conditions)
VALUES (319, '2A', 'Business')
ON CONFLICT (aircraft_code, seat_no)
DO NOTHING;

INSERT INTO seats_copy (aircraft_code, seat_no, fare_conditions)
VALUES (319, '2A', 'Business')
ON CONFLICT ON CONSTRAINT seats_copy_pkey
DO NOTHING;
