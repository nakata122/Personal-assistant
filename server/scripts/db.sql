create database if not exists assistant;

use assistant;

CREATE TYPE user_role AS ENUM  ('admin', 'user', 'guest');

create table if not exists users (
    user_id serial PRIMARY KEY,
    email varchar(255) NOT NULL UNIQUE,
    google_id varchar(255) NOT NULL,
    role user_role DEFAULT 'user',
    name varchar(255),
    picture varchar(255)
);

create table if not exists emails (
    email_id serial PRIMARY KEY,
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    title varchar(255),
    summary varchar(2000),
    profilepic varchar(255),
    score REAL,
    tags TEXT[]
);