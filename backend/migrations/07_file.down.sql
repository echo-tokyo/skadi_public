ALTER TABLE task_file DROP CONSTRAINT task_file_task_fk;

ALTER TABLE task_file DROP CONSTRAINT task_file_file_fk;

ALTER TABLE solution_file DROP CONSTRAINT solution_file_solution_fk;

ALTER TABLE solution_file DROP CONSTRAINT solution_file_file_fk;

DROP TABLE IF EXISTS file;

DROP TABLE IF EXISTS task_file;

DROP TABLE IF EXISTS solution_file;