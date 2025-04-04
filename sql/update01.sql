ALTER TABLE user ADD COLUMN uid varchar(32) not null comment 'uid';

ALTER TABLE question
    MODIFY COLUMN user_id VARCHAR(32) NOT NULL;

ALTER TABLE question_submit
    MODIFY COLUMN user_id VARCHAR(32) NOT NULL;