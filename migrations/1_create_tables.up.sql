CREATE TABLE users (
    id UUID NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT,
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    deleted_email TEXT, 
    CONSTRAINT email_unique UNIQUE (email)
);