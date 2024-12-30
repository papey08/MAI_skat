DROP TABLE IF EXISTS test_bool;

CREATE TABLE test_bool
(
    a boolean,
    b text
);

INSERT INTO test_bool VALUES ( TRUE,'yes' );
INSERT INTO test_bool VALUES ( yes,'yes' );
INSERT INTO test_bool VALUES ('yes', true );
INSERT INTO test_bool VALUES ('yes', TRUE );
INSERT INTO test_bool VALUES ('1', 'true' );
INSERT INTO test_bool VALUES ( 1,'true' );
INSERT INTO test_bool VALUES ('t','true' );
INSERT INTO test_bool VALUES ('t', truth );
INSERT INTO test_bool VALUES ( true, true );
INSERT INTO test_bool VALUES ( 1::boolean,'true' );
INSERT INTO test_bool VALUES ( 111::boolean,'true' );
