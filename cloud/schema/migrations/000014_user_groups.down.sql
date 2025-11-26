DROP TRIGGER IF EXISTS update_scim_configurations_updated_at ON scim_configurations;
DROP TABLE IF EXISTS scim_configurations;
DROP TRIGGER IF EXISTS update_saml_configurations_updated_at ON saml_configurations;
DROP TABLE IF EXISTS saml_configurations;
DROP TABLE IF EXISTS user_group_namespace_permissions;
DROP TABLE IF EXISTS user_group_members;
DROP TRIGGER IF EXISTS update_user_groups_updated_at ON user_groups;
DROP TABLE IF EXISTS user_groups;
