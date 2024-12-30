CREATE TEMP TABLE aircrafts_log AS
SELECT * FROM aircrafts WITH NO DATA;

ALTER TABLE aircrafts_log
ADD COLUMN when_add timestamp;

ALTER TABLE aircrafts_log
ADD COLUMN operation text;

ALTER TABLE aircrafts_log
ADD COLUMN "current_timestamp" timestamp;

CREATE TEMP TABLE aircrafts_tmp
( LIKE aircrafts INCLUDING CONSTRAINTS INCLUDING INDEXES );

WITH add_row AS
( INSERT INTO aircrafts_tmp
SELECT * FROM aircrafts
RETURNING *
)
INSERT INTO aircrafts_log
SELECT add_row.aircraft_code, add_row.model, add_row.range,
current_timestamp,'INSERT',current_timestamp
FROM add_row;

WITH add_row AS
( INSERT INTO aircrafts_tmp
VALUES ('SU9', 'Sukhoi SuperJet-100', 3000 )
ON CONFLICT DO NOTHING
RETURNING *
)
INSERT INTO aircrafts_log
SELECT add_row.aircraft_code, add_row.model, add_row.range,
current_timestamp,'INSERT',current_timestamp
FROM add_row;

WITH update_row AS
( UPDATE aircrafts_tmp
SET range = range * 1.2
WHERE model ~ '^Bom'
RETURNING *
)
INSERT INTO aircrafts_log
SELECT ur.aircraft_code, ur.model, ur.range,
current_timestamp,'UPDATE',current_timestamp
FROM update_row ur;

SELECT * FROM aircrafts_log;
