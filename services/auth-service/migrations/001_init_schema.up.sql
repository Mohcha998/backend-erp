-- ======================================
-- DIVISIONS
-- ======================================
CREATE TABLE IF NOT EXISTS divisions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- ======================================
-- MENUS
-- ======================================
CREATE TABLE IF NOT EXISTS menus (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    path VARCHAR(100),
    created_at TIMESTAMP DEFAULT now()
);

-- ======================================
-- PERMISSIONS
-- ======================================
CREATE TABLE IF NOT EXISTS permissions (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    code VARCHAR(50) NOT NULL UNIQUE,
    description TEXT
);

-- ======================================
-- ROLES
-- ======================================
CREATE TABLE IF NOT EXISTS roles (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT now()
);

-- ======================================
-- USERS
-- ======================================
CREATE TABLE IF NOT EXISTS users (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    division_id INTEGER,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_users_division
        FOREIGN KEY (division_id)
        REFERENCES divisions(id)
        ON DELETE SET NULL
);

-- ======================================
-- DIVISION MENUS (MANY TO MANY)
-- ======================================
CREATE TABLE IF NOT EXISTS division_menus (
    division_id INTEGER NOT NULL,
    menu_id INTEGER NOT NULL,
    PRIMARY KEY (division_id, menu_id),
    CONSTRAINT fk_division_menus_division
        FOREIGN KEY (division_id)
        REFERENCES divisions(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_division_menus_menu
        FOREIGN KEY (menu_id)
        REFERENCES menus(id)
        ON DELETE CASCADE
);

-- ======================================
-- ROLE MENUS (MANY TO MANY)
-- ======================================
CREATE TABLE IF NOT EXISTS role_menus (
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

-- ======================================
-- ROLE PERMISSIONS (MANY TO MANY)
-- ======================================
CREATE TABLE IF NOT EXISTS role_permissions (
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

-- ======================================
-- USER ROLES (MANY TO MANY)
-- ======================================
CREATE TABLE IF NOT EXISTS user_roles (
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
