CREATE TABLE "Distance" (
    "distance" INTEGER NOT NULL,
    "point_begin" VARCHAR(50) NOT NULL,
    "point_end" VARCHAR(50) NOT NULL
);

INSERT INTO "Distance" (distance, point_begin, point_end)
VALUES
    (673, 'Москва', 'Белгород'),
    (673, 'Москва', 'Белгород'),
    (331, 'Иваново', 'Москва'),
    (331, 'Москва', 'Иваново'),
    (1156, 'Москва', 'Донбасс'),
    (1156, 'Донбасс', 'Москва'),
    (8410, 'Москва', 'Хабаровск'),
    (8410, 'Хабаровск', 'Москва');
