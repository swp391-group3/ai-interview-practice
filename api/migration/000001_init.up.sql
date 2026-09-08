CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE role AS ENUM (
    'participant',
    'jury',
    'admin'
);

CREATE TABLE IF NOT EXISTS accounts (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),

    email text UNIQUE NOT NULL,
    full_name text NOT NULL,
    password_hash varchar(128) NOT NULL,

    role role NOT NULL DEFAULT 'participant'::role,
    is_locked bool NOT NULL DEFAULT false,

    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
