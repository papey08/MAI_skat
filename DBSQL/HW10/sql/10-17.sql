--- добавляем сотрудника в иерархию
INSERT INTO personnel (emp_nbr, emp_name, address, birth_date)
VALUES (9, 'Владислав', 'просп. Реляционных СУБД', '2000-01-01');

INSERT INTO org_chart (job_title, emp_nbr, boss_emp_nbr, salary)
VALUES ('Бэкендер', 9, 8, 300);

SELECT * FROM create_paths;
