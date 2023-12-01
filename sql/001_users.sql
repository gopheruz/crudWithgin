CREATE TABLE users(id UUID NOT NULL ,
                   first_name TEXT NOT NULL ,
                   last_name TEXT NOT NULL,
                   email TEXT NOT NULL,
                   created_at TIMESTAMP NOT NULL ,
                   deleted_at TIMESTAMP )