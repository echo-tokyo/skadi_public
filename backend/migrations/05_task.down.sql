ALTER TABLE task DROP CONSTRAINT task_teacher_fk;

ALTER TABLE solution DROP CONSTRAINT solution_task_fk;

ALTER TABLE solution DROP CONSTRAINT solution_status_fk;

ALTER TABLE solution DROP CONSTRAINT solution_student_fk;

DROP TABLE IF EXISTS task;

DROP TABLE IF EXISTS status;

DROP TABLE IF EXISTS solution;