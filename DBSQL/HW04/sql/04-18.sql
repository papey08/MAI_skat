ALTER TABLE airports_data
ADD COLUMN location jsonb;

UPDATE airports_data
SET location =
    '{"city": "Москва",
      "country": "Россия",
      "coordinates": {
        "lattitude": 37.4,
        "longitude": 56.0
      }}
    '
WHERE airport_code = 'SVO';

SELECT location FROM airports_data
WHERE airport_code = 'SVO';

SELECT location->'city' FROM airports_data
WHERE airport_code = 'SVO';

SELECT location #> '{coordinates, longitude}' FROM airports_data
WHERE airport_code = 'SVO';
