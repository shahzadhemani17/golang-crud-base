CREATE TABLE IF NOT EXISTS "authors"(
    id serial PRIMARY KEY,
    name varchar(255) not null,
    age integer not null,
    email varchar unique not null
);