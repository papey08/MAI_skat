BEGIN;
SELECT * FROM delete_and_promote_subtree(5);
SELECT * FROM personnel_org_chart;
SELECT * FROM create_paths;
ROLLBACK;

BEGIN;
SELECT * FROM delete_and_promote_subtree((SELECT emp_nbr FROM personnel WHERE emp_name = 'Анна'));
SELECT * FROM personnel_org_chart;
SELECT * FROM create_paths;
ROLLBACK;
