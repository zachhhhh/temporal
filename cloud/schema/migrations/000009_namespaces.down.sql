DROP TABLE IF EXISTS namespace_search_attributes;
DROP TABLE IF EXISTS namespace_certificate_filters;
DROP TABLE IF EXISTS namespace_certificates;
DROP TRIGGER IF EXISTS update_namespaces_updated_at ON cloud_namespaces;
DROP TABLE IF EXISTS cloud_namespaces;
DROP TYPE IF EXISTS namespace_state;
