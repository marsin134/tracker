CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_name VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    refresh_token_hash TEXT,
    refresh_token_expiry_time TIMESTAMP WITH TIME ZONE
);

CREATE UNIQUE INDEX idx_users_user_name ON users(user_name);