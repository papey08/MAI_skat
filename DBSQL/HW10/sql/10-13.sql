SELECT * FROM tree_test();

BEGIN;
UPDATE org_chart
SET boss_emp_nbr = 3
WHERE emp_nbr = 2;

UPDATE org_chart
SET boss_emp_nbr = 2
WHERE emp_nbr = 3;

SELECT * FROM tree_test();
ROLLBACK;

BEGIN;
UPDATE org_chart
SET boss_emp_nbr = 5
WHERE emp_nbr = 4;

UPDATE org_chart
SET boss_emp_nbr = 6
WHERE emp_nbr = 5;

UPDATE org_chart
SET boss_emp_nbr = 4
WHERE emp_nbr = 6;

SELECT * FROM tree_test();
ROLLBACK;

