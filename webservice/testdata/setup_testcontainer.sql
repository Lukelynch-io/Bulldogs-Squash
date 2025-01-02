-- Connect as the default PostgreSQL user and database
-- Replace "postgres" with your desired username and database if needed

\c testdb;
-- Create the 'testdb' table if it doesn't already exist
CREATE TABLE newtable (
    id INT,
    name VARCHAR(255)
);

-- Insert a record into the 'testdb' table
INSERT INTO newtable (id, name) VALUES (1, 'test');
