-- 000001_create_transactions.up.sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    amount DOUBLE PRECISION,
    merchant TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);