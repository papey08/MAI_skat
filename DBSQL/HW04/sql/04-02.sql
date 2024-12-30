ALTER TABLE progress
    ADD COLUMN test_form text,
    ADD CHECK (
    ( test_form ='экзамен' AND mark IN ( 3, 4, 5 ))
    OR
    ( test_form ='зачет' AND mark IN ( 0, 1 ))
);

INSERT INTO progress (record_book, subject, acad_year, term, mark, test_form)
VALUES (12345, 'МАТАН', '2020', 2, 3, 'экзамен');

INSERT INTO progress (record_book, subject, acad_year, term, mark, test_form)
VALUES (54321, 'БД', '2024', 1, 4, 'зачет');

INSERT INTO progress (record_book, subject, acad_year, term, mark, test_form)
VALUES (54321, 'ИИ', '2024', 1, 1, 'зачет');

ALTER TABLE progress
DROP CONSTRAINT progress_mark_check;

ALTER TABLE progress
    ADD CHECK (
    ( test_form ='экзамен' AND mark IN ( 3, 4, 5 ))
    OR
    ( test_form ='зачет' AND mark IN ( 0, 1 ))
);

INSERT INTO progress (record_book, subject, acad_year, term, mark, test_form)
VALUES (54321, 'ИИ', '2024', 1, 1, 'зачет');
