ALTER TABLE class DROP CONSTRAINT class_teacher_fk;

ALTER TABLE user DROP CONSTRAINT student_class_fk;

ALTER TABLE user DROP COLUMN class_id;

DROP TABLE IF EXISTS class;