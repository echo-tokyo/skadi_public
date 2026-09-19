DROP TABLE IF EXISTS file;

DROP TABLE IF EXISTS task_file;

DROP TABLE IF EXISTS solution_file;

CREATE TABLE IF NOT EXISTS file (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    path VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL DEFAULT "file1",
    size BIGINT NULL,
    mime_type VARCHAR(50) NULL DEFAULT "text/plain"
);

CREATE TABLE IF NOT EXISTS task_file (
    task_id BIGINT NOT NULL,
    file_id BIGINT NOT NULL,
    PRIMARY KEY (task_id, file_id)
);

CREATE TABLE IF NOT EXISTS solution_file (
    solution_id BIGINT NOT NULL,
    file_id BIGINT NOT NULL,
    PRIMARY KEY (solution_id, file_id)
);

ALTER TABLE task_file
ADD CONSTRAINT task_file_task_fk FOREIGN KEY (task_id) REFERENCES task (id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE task_file
ADD CONSTRAINT task_file_file_fk FOREIGN KEY (file_id) REFERENCES file (id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE solution_file
ADD CONSTRAINT solution_file_solution_fk FOREIGN KEY (solution_id) REFERENCES solution (id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE solution_file
ADD CONSTRAINT solution_file_file_fk FOREIGN KEY (file_id) REFERENCES file (id) ON UPDATE CASCADE ON DELETE CASCADE;