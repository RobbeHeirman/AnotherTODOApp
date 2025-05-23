CREATE TABLE IF NOT EXISTS users (
                       id SERIAL PRIMARY KEY,
                       email TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
    );


CREATE INDEX IF NOT EXISTS  idx_email ON users (email)

