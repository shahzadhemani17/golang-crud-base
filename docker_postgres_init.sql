CREATE USER testuser WITH PASSWORD 'password' CREATEDB;
CREATE DATABASE postgres
    WITH
    OWNER = testuser
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;