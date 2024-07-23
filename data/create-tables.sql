CREATE DATABASE `Todos-Service`;

USE `Todos-Service`;

DROP TABLE IF EXISTS Todo;

CREATE TABLE Todo (
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(500) NOT NULL,
    description VARCHAR(255) NOT NULL,
    completed BOOLEAN NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO
    Todo (title, description, completed)
VALUES
    ('Walk the dog', 'No more than 2 km', false),
    (
        'Clean my bedroom',
        'Use orange flavoring',
        false
    ),
    (
        'Learn React',
        'Watch some video from Youtube',
        true
    );