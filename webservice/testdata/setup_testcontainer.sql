-- Replace "postgres" with your desired username and database if needed

-- \c testdb
-- Create the 'testdb' table if it doesn't already exist
CREATE TABLE hello (
    id INT,
    name VARCHAR(255)
);

-- Insert a record into the 'testdb' table
INSERT INTO hello (id, name) VALUES (1, 'test');
