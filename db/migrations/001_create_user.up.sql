CREATE TABLE users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    username varchar(200) NOT NULL UNIQUE,
    password varchar(100) NOT NULL
);