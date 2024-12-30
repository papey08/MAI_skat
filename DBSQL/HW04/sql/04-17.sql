CREATE VIEW aircraft_ranges AS
    SELECT aircraft_code, range FROM aircrafts_data;
SELECT * FROM aircraft_ranges;

CREATE VIEW aircrafts_with_high_ranges AS
    SELECT * FROM aircrafts_data
    WHERE range >= 6000;
SELECT * FROM aircrafts_with_high_ranges;
