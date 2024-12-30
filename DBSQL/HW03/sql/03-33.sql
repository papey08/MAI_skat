DROP TABLE IF EXISTS pilots;

CREATE TABLE pilots
(
    pilot_name text,
    schedule   integer[],
    meal       text[][]
);

INSERT INTO pilots
VALUES ('ИВАН',
        '{1, 3, 5, 6, 7}'::integer[],
        '{ { "сосиска", "макароны", "кофе" },
        { "котлета", "каша", "кофе" },
        { "сосиска", "каша", "кофе" },
        { "котлета", "каша", "чай" } }'::text[][]);

SELECT meal[2][2] FROM pilots;

UPDATE pilots
SET meal[2][2] = 'суп';

SELECT meal[2][2] FROM pilots;
