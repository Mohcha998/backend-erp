-- =====================================================
-- DROP CHILD TABLES (RELATION)
-- =====================================================

DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS audit_logs;
DROP TABLE IF EXISTS activity_logs;

DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS role_menus;
DROP TABLE IF EXISTS division_roles;
DROP TABLE IF EXISTS user_roles;

-- =====================================================
-- DROP MASTER TABLES
-- =====================================================

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS menus;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS divisions;
