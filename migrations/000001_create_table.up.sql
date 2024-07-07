CREATE TABLE soldiers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    date_of_birth TIMESTAMP NOT NULL
    phone_number VARCHAR(13) NOT NULL,
    group_id UUID NOT NULL,
    join_date TIMESTAMP NOT NULL DEFAULT now(),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE departments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL
);

CREATE TABLE groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    department_id UUID REFERENCES departments(id),
    commanders_id UUID UNIQUE
);

CREATE TYPE position AS ENUM ('o''nbosh', 'yuzbosh', 'mingbosh', 'amir', 'qo''mondon');

CREATE TABLE commanders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    date_of_birth TIMESTAMP NOT NULL
    phone_number VARCHAR(13) NOT NULL,
    position position,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);
