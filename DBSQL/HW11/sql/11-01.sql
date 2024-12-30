CREATE TABLE books (
    book_id INTEGER PRIMARY KEY,
    book_description TEXT
);

COPY books FROM '/books3.txt';

ALTER TABLE books ADD COLUMN ts_description tsvector;

UPDATE books
SET ts_description = to_tsvector('russian', book_description);

EXPLAIN ANALYSE SELECT
    CASE
        WHEN ts_description @@ to_tsquery('python') THEN 'python'
        WHEN ts_description @@ to_tsquery('java') THEN 'java'
        WHEN ts_description @@ to_tsquery('pascal') THEN 'pascal'
        WHEN ts_description @@ to_tsquery('php') THEN 'php'
        WHEN ts_description @@ to_tsquery('sql') THEN 'sql'
    END AS language,
    COUNT(*) AS amount
FROM books
WHERE ts_description @@ to_tsquery('python | java | pascal | php | sql')
GROUP BY language
ORDER BY amount DESC;

CREATE INDEX books_idx ON books
USING GIN (ts_description);

EXPLAIN ANALYSE SELECT
    CASE
        WHEN ts_description @@ to_tsquery('python') THEN 'python'
        WHEN ts_description @@ to_tsquery('java') THEN 'java'
        WHEN ts_description @@ to_tsquery('pascal') THEN 'pascal'
        WHEN ts_description @@ to_tsquery('php') THEN 'php'
        WHEN ts_description @@ to_tsquery('sql') THEN 'sql'
    END AS language,
    COUNT(*) AS amount
FROM books
WHERE ts_description @@ to_tsquery('python | java | pascal | php | sql')
GROUP BY language
ORDER BY amount DESC;
