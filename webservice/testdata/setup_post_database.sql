DO
$$
BEGIN
    IF NOT EXISTS (
        SELECT FROM pg_database WHERE datname = 'Posts'
    ) THEN
        CREATE DATABASE "Posts";
    END IF;
END
$$;

GRANT ALL PRIVILEGES ON DATABASE "Posts" TO postgres;
\c Posts

CREATE TABLE IF NOT EXISTS "Posts" (
    Id UUID PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    Description TEXT,
    ImageUrl VARCHAR(255)
)
