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
RETURNING aircraft_code, model, range,
current_timestamp,'INSERT'
)
INSERT INTO aircrafts_log
SELECT * FROM add_row;

SELECT * FROM aircrafts_log;