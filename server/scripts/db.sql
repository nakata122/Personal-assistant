create database if not exists assistant;

use assistant;

CREATE TYPE user_role AS ENUM  ('admin', 'user', 'guest');

create table if not exists users (
    user_id serial PRIMARY KEY,
    google_id varchar(255) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    role user_role DEFAULT 'user',
    name varchar(255),
    picture varchar(255)
);