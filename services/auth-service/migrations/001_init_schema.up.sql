-- =====================================================
-- EXTENSION (OPTIONAL - PostgreSQL)
-- =====================================================
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- =====================================================
-- DIVISIONS
-- =====================================================
CREATE TABLE divisions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP NULL
);

-- =====================================================
-- ROLES
-- =====================================================
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP NULL
);

-- =====================================================
-- MENUS
-- =====================================================
CREATE TABLE menus (
    id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    path VARCHAR(150),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP NULL
);

-- =====================================================
-- PERMISSIONS
-- =====================================================
CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP NULL
);

-- =====================================================
-- USERS
-- =====================================================
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    division_id INTEGER,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP NULL,

    CONSTRAINT fk_users_division
        FOREIGN KEY (division_id)
        REFERENCES divisions(id)
        ON DELETE SET NULL
);

-- =====================================================
-- USER ROLES (Many-to-Many)
-- =====================================================
CREATE TABLE user_roles (
    user_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, role_id),

    CONSTRAINT fk_user_roles_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_user_roles_role
        FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE CASCADE
);

-- =====================================================
-- DIVISION ROLES
-- =====================================================
CREATE TABLE division_roles (
    division_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    PRIMARY KEY (division_id, role_id),

    CONSTRAINT fk_division_roles_division
        FOREIGN KEY (division_id)
        REFERENCES divisions(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_division_roles_role
        FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE CASCADE
);

-- =====================================================
-- ROLE MENUS
-- =====================================================
CREATE TABLE role_menus (
    role_id INTEGER NOT NULL,
    menu_id INTEGER NOT NULL,
    PRIMARY KEY (role_id, menu_id),

    CONSTRAINT fk_role_menus_role
        FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_role_menus_menu
        FOREIGN KEY (menu_id)
        REFERENCES menus(id)
        ON DELETE CASCADE
);

-- =====================================================
-- ROLE PERMISSIONS
-- =====================================================
CREATE TABLE role_permissions (
    role_id INTEGER NOT NULL,
    permission_id INTEGER NOT NULL,
    PRIMARY KEY (role_id, permission_id),

    CONSTRAINT fk_role_permissions_role
        FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_role_permissions_permission
        FOREIGN KEY (permission_id)
        REFERENCES permissions(id)
        ON DELETE CASCADE
);

-- =====================================================
-- ACTIVITY LOGS
-- =====================================================
CREATE TABLE activity_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    activity VARCHAR(255) NOT NULL,
    ip_address VARCHAR(50),
    created_at TIMESTAMP DEFAULT now(),

    CONSTRAINT fk_activity_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_activity_user ON activity_logs(user_id);

-- =====================================================
-- AUDIT LOGS
-- =====================================================
CREATE TABLE audit_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    action VARCHAR(50) NOT NULL,
    entity VARCHAR(50) NOT NULL,
    entity_id INTEGER,
    ip_address VARCHAR(50),
    created_at TIMESTAMP DEFAULT now(),

    CONSTRAINT fk_audit_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_audit_user ON audit_logs(user_id);

-- =====================================================
-- REFRESH TOKENS
-- =====================================================
CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    token TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    revoked BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT now(),

    CONSTRAINT fk_refresh_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE

-- =====================================================
-- INDEXES (PERFORMANCE)
-- =====================================================
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_division ON users(division_id);
CREATE INDEX idx_role_menus_role ON role_menus(role_id);
CREATE INDEX idx_role_permissions_role ON role_permissions(role_id);
CREATE INDEX idx_refresh_token ON refresh_tokens(token);
CREATE INDEX idx_refresh_user ON refresh_tokens(user_id);
