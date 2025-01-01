-- Connect as the default PostgreSQL user and database
-- Replace "postgres" with your desired username and database if needed

-- Create a user named 'docker'
CREATE USER docker;

-- Create a database named 'docker'
-- CREATE DATABASE docker;

-- Grant all privileges on the 'docker' database to the 'docker' user
GRANT ALL PRIVILEGES ON DATABASE docker TO docker;

-- Connect to the 'docker' database to execute further commands
\c docker

-- Create the 'testdb' table if it doesn't already exist
CREATE TABLE IF NOT EXISTS testdb (
    id INT,
    name VARCHAR(255)
);

-- Insert a record into the 'testdb' table
INSERT INTO testdb (id, name) VALUES (1, 'test');
