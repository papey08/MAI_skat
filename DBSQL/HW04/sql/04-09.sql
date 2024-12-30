ALTER TABLE students ADD CHECK ( name <>'' );

INSERT INTO students VALUES ( 12346,' ', 0406, 112233 );
INSERT INTO students VALUES ( 12347,' ', 0407, 112234 );

SELECT *, length( name ) FROM students;

TRUNCATE TABLE students;

ALTER TABLE students ADD CHECK (trim(name) <> '');

INSERT INTO students VALUES ( 12346,' ', 0406, 112233 );

SELECT *, length( name ) FROM students;
