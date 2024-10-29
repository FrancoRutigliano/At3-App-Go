CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,           -- Nombre
    last_name VARCHAR(100) NOT NULL,            -- Apellido
    email VARCHAR(255) NOT NULL UNIQUE,         -- Email
    password VARCHAR(255) NOT NULL,             -- Contraseña
    phone_number VARCHAR(50) NOT NULL,          -- Teléfono
    tax_id VARCHAR(100) NOT NULL,               -- CUIT o equivalente (número de identificación fiscal)
    wallet_address VARCHAR(255) NOT NULL,       -- Dirección de billetera
    identity_document_url VARCHAR(255) NOT NULL,         -- URL del documento de identidad (DNI o equivalente)
    is_uiff BOOLEAN DEFAULT FALSE,              -- UIFF
    is_exposed BOOLEAN DEFAULT FALSE,           -- Expuesta
    role INT NOT NULL DEFAULT 1,                -- Rol: 1 (no validado), 2 (validado sin doc.), 3 (validado completo), 4 (admin)
    created_at BIGINT NOT NULL,                 -- Timestamp UNIX
    updated_at BIGINT NOT NULL                  -- Timestamp UNIX
);

CREATE TABLE companies (
    id UUID PRIMARY KEY,
    business_name VARCHAR(255) NOT NULL,        -- Razón social
    email VARCHAR(255) NOT NULL UNIQUE,         -- Email de la empresa
    legal_representative_name VARCHAR(255) NOT NULL,    -- Nombre del representante legal
    legal_representative_id VARCHAR(100) NOT NULL,      -- DNI del representante legal
    password VARCHAR(255) NOT NULL,             -- Contraseña
    phone_number VARCHAR(50) NOT NULL,          -- Teléfono
    tax_id VARCHAR(100) NOT NULL,               -- CUIT o equivalente (número de identificación fiscal)
    address VARCHAR(255) NOT NULL,              -- Domicilio fiscal de la empresa
    company_certificate_url VARCHAR(255) NOT NULL,       -- URL del certificado de la empresa
    role INT NOT NULL DEFAULT 1,                -- Rol: 1 (no validado), 2 (validado sin doc.), 3 (validado completo), 4 (admin)
    created_at BIGINT NOT NULL,                 -- Timestamp UNIX
    updated_at BIGINT NOT NULL                  -- Timestamp UNIX
);

