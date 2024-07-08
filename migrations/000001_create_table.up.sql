CREATE TABLE soldiers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    date_of_birth TIMESTAMP NOT NULL,
    phone_number VARCHAR(13) NOT NULL,
    group_id UUID NOT NULL,
    join_date TIMESTAMP NOT NULL DEFAULT now(),
    end_date TIMESTAMP NOT NULL DEFAULT now() + INTERVAL '12 month',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE departments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    commanders_id UUID UNIQUE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    department_id UUID REFERENCES departments(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TYPE position_commander AS ENUM ('o''nbosh', 'yuzbosh', 'mingbosh', 'amir', 'qo''mondon');

CREATE TABLE commanders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    date_of_birth TIMESTAMP NOT NULL,
    phone_number VARCHAR(13) NOT NULL,
    position position_commander NOT NULL DEFAULT 'o''nbosh',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE use_bullets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    quantity_weapon INT NOT NULL,
    quantity_big_weapon INT NOT NULL,
    soldier_id UUID REFERENCES soldiers(id),
    date DATE NOT NULL
);

CREATE TABLE use_fuels (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    diesel INT NOT NULL,
    petrol INT NOT NULL,
    soldier_id UUID REFERENCES soldiers(id),
    date DATE NOT NULL
);