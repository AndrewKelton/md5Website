-- SQLite
CREATE TABLE users (
    username TEXT NOT NULL,
    pass TEXT NOT NULL
);

INSERT INTO users(username, pass) VALUES ('Andrew', '1234')

SELECT * FROM users