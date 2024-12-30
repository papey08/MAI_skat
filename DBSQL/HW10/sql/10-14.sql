SELECT * FROM up_tree_traversal(6);

SELECT * FROM up_tree_traversal(4);

SELECT * FROM up_tree_traversal2(6) AS (emp int, boss int);

SELECT * FROM up_tree_traversal((SELECT emp_nbr FROM personnel WHERE emp_name = 'Иван'));

SELECT * FROM up_tree_traversal((SELECT emp_nbr FROM personnel WHERE emp_name = 'Захар'));
