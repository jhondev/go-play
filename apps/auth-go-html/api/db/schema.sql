CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    name VARCHAR(200),
    password_hash BINARY(60),
    telephone VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- CREATE TABLE user_identities (
--     id CHAR(36) PRIMARY KEY,
--     user_id CHAR(36) NOT NULL,
--     provider_name VARCHAR(50) NOT NULL,
--     provider_id VARCHAR(200) NOT NULL,
--     auth_token TEXT,
--     refresh_token TEXT,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
-- );
-- FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE