CREATE OR REPLACE FUNCTION get_addresses()
RETURNS VOID AS $$
    DECLARE cursor CURSOR FOR
        SELECT emp_name, address FROM personnel;
    record RECORD;
    BEGIN
       OPEN cursor;
       LOOP
           FETCH cursor INTO record;
           EXIT WHEN NOT FOUND;
           RAISE NOTICE 'name: %, address: %', record.emp_name, record.address;
       END LOOP;
       CLOSE cursor;
    END;
    $$ LANGUAGE plpgsql;

SELECT get_addresses();
