ALTER TABLE profile DROP CONSTRAINT profile_user_fk;

ALTER TABLE profile DROP CONSTRAINT profile_contact_fk;

ALTER TABLE profile DROP CONSTRAINT profile_parent_contact_fk;

DROP TABLE IF EXISTS contact;

DROP TABLE IF EXISTS profile;