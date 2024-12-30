DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS progress;

CREATE TABLE students
(
    record_book NUMERIC(5) NOT NULL,
    name        text       NOT NULL,
    doc_ser     NUMERIC(4),
    doc_num     NUMERIC(6)
);

CREATE TABLE progress
(
    record_book NUMERIC(5) NOT NULL,
    subject     text       NOT NULL,
    acad_year   text       NOT NULL,
    term        numeric(1) NOT NULL CHECK (term = 1 OR term = 2),
    mark        numeric(1) DEFAULT 5 CHECK ( mark >= 3 AND mark <= 5 )
);
