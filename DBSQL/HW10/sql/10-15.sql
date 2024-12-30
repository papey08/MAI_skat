SELECT * FROM personnel_org_chart;
SELECT * FROM create_paths;


BEGIN;
SELECT * FROM delete_subtree(6);
SELECT * FROM personnel_org_chart;
SELECT * FROM create_paths;
ROLLBACK;

BEGIN;
SELECT * FROM delete_subtree((SELECT emp_nbr FROM personnel WHERE emp_name = 'Захар'));
SELECT * FROM personnel_org_chart;
SELECT * FROM create_paths;
ROLLBACK;
