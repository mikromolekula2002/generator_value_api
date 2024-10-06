CREATE TABLE IF NOT EXISTS generated_values (
    id SERIAL PRIMARY KEY,
    request_id VARCHAR(50) NOT NULL UNIQUE,
    random_value TEXT NOT NULL UNIQUE,
    value_type VARCHAR(50) NOT NULL,
    length INT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users_requests (
    id SERIAL PRIMARY KEY,
    request_id VARCHAR(50) NOT NULL,
     random_value TEXT NOT NULL,
    user_agent TEXT NOT NULL,
    url TEXT NOT NULL,
    method VARCHAR(10) NOT NULL,
    request_count INT DEFAULT 1,
    created_at TIMESTAMPTZ DEFAULT NOW()
);