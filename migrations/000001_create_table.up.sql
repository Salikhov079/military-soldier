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



INSERT INTO use_bullets (quantity_weapon, quantity_big_weapon, soldier_id, date) VALUES
(5, 2, '720f5034-cfb5-4cc4-bb45-f1f4a307f858', '2024-07-01'),
(3, 1, '720f5034-cfb5-4cc4-bb45-f1f4a307f858', '2024-07-01'),
(4, 2, '720f5034-cfb5-4cc4-bb45-f1f4a307f858', '2024-07-01'),
(2, 0, '720f5034-cfb5-4cc4-bb45-f1f4a307f858', '2024-07-02'),
(7, 3, '720f5034-cfb5-4cc4-bb45-f1f4a307f858', '2024-07-02'),
(6, 1, '720f5034-cfb5-4cc4-bb45-f1f4a307f858', '2024-07-03'),
(8, 2, '720f5034-cfb5-4cc4-bb45-f1f4a307f858', '2024-07-03'),
(5, 1, '720f5034-cfb5-4cc4-bb45-f1f4a307f858', '2024-07-03'),
(9, 4, '720f5034-cfb5-4cc4-bb45-f1f4a307f858', '2024-07-04'),
(3, 1, '720f5034-cfb5-4cc4-bb45-f1f4a307f858', '2024-07-04');


INSERT INTO soldiers (id, name, email, date_of_birth, phone_number, group_id, join_date, end_date)
VALUES 
    (gen_random_uuid(), 'John Do', 'jon@example.com', '1990-01-01', '1234567890', '6bf3ee89-4bc2-453d-97c6-69a1697ac556', '2022-01-01', '2024-01-01'),
    (gen_random_uuid(), 'Jane Smth', 'jane@xample.com', '1992-02-02', '2345678901', '6bf3ee89-4bc2-453d-97c6-69a1697ac556', '2022-01-01', '2024-01-01'),
    (gen_random_uuid(), 'Alice Jhnson', 'alce@example.com', '1993-03-03', '3456789012', '6bf3ee89-4bc2-453d-97c6-69a1697ac556', '2022-03-01', '2024-01-01'),
    (gen_random_uuid(), 'Bob Bron', 'bob@exmple.com', '1994-04-04', '4567890123', '6bf3ee89-4bc2-453d-97c6-69a1697ac556', '2022-01-01', '2024-01-01'),
    (gen_random_uuid(), 'Charlie avis', 'carlie@example.com', '1995-05-05', '5678901234', '6bf3ee89-4bc2-453d-97c6-69a1697ac556', '2022-01-01', '2024-01-01'),
    (gen_random_uuid(), 'Diana Evns', 'dianaexample.com', '1996-06-06', '6789012345', '6bf3ee89-4bc2-453d-97c6-69a1697ac556', '2022-01-01', '2024-01-01'),
    (gen_random_uuid(), 'Edward reen', 'edard@example.com', '1997-07-07', '7890123456', '6bf3ee89-4bc2-453d-97c6-69a1697ac556', '2022-01-01', '2024-01-01'),
    (gen_random_uuid(), 'Fiona Haris', 'fina@example.com', '1998-08-08', '8901234567', '6bf3ee89-4bc2-453d-97c6-69a1697ac556', '2022-01-01', '2024-01-01'),
    (gen_random_uuid(), 'George Kng', 'georgeexample.com', '1999-09-09', '9012345678', '6bf3ee89-4bc2-453d-97c6-69a1697ac556', '2022-01-01', '2024-01-01'),
    (gen_random_uuid(), 'Hannah ee', 'hanna@example.com', '2000-10-10', '0123456789', '6bf3ee89-4bc2-453d-97c6-69a1697ac556', '2022-01-01', '2024-01-01');
