DO
$$
BEGIN
    IF NOT EXISTS (
        SELECT FROM pg_database WHERE datname = 'Auth'
    ) THEN
        CREATE DATABASE "Auth";
    END IF;
END
$$;

GRANT ALL PRIVILEGES ON DATABASE "Auth" TO postgres;
\c Auth
-- Create the 'testdb' table if it doesn't already exist
CREATE TABLE IF NOT EXISTS "Users" (
    userId UUID NOT NULL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    passwordHash BYTEA NOT NULL,
    role VARCHAR(255)
);

