-- Namespace state enum
CREATE TYPE namespace_state AS ENUM (
    'pending', 'provisioning', 'active', 'updating', 
    'deleting', 'deleted', 'suspended', 'failed'
);

-- Cloud namespaces table
CREATE TABLE cloud_namespaces (
    id VARCHAR(255) PRIMARY KEY,
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    region VARCHAR(50) NOT NULL,
    cluster_id VARCHAR(255),
    state namespace_state NOT NULL DEFAULT 'pending',
    -- Configuration
    retention_days INT NOT NULL DEFAULT 7,
    deletion_protected BOOLEAN NOT NULL DEFAULT FALSE,
    -- HA configuration
    ha_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    standby_region VARCHAR(50),
    -- Codec server
    codec_endpoint TEXT,
    codec_pass_token BOOLEAN NOT NULL DEFAULT FALSE,
    codec_include_credentials BOOLEAN NOT NULL DEFAULT FALSE,
    -- Endpoints
    grpc_endpoint VARCHAR(255),
    web_endpoint VARCHAR(255),
    metrics_endpoint VARCHAR(255),
    -- Metadata
    tags JSONB DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(organization_id, name)
);

CREATE INDEX idx_namespaces_org ON cloud_namespaces(organization_id);
CREATE INDEX idx_namespaces_region ON cloud_namespaces(region);
CREATE INDEX idx_namespaces_state ON cloud_namespaces(state);
CREATE INDEX idx_namespaces_cluster ON cloud_namespaces(cluster_id);

CREATE TRIGGER update_namespaces_updated_at
    BEFORE UPDATE ON cloud_namespaces
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Namespace certificates (for mTLS)
CREATE TABLE namespace_certificates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    namespace_id VARCHAR(255) NOT NULL REFERENCES cloud_namespaces(id) ON DELETE CASCADE,
    certificate_pem TEXT NOT NULL,
    fingerprint VARCHAR(255) NOT NULL,
    issuer VARCHAR(255),
    subject VARCHAR(255),
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_ns_certs_namespace ON namespace_certificates(namespace_id);
CREATE INDEX idx_ns_certs_fingerprint ON namespace_certificates(fingerprint);

-- Namespace certificate filters
CREATE TABLE namespace_certificate_filters (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    namespace_id VARCHAR(255) NOT NULL REFERENCES cloud_namespaces(id) ON DELETE CASCADE,
    common_name VARCHAR(255),
    organization VARCHAR(255),
    organizational_unit VARCHAR(255),
    subject_alternative_name VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_ns_cert_filters_namespace ON namespace_certificate_filters(namespace_id);

-- Namespace search attributes
CREATE TABLE namespace_search_attributes (
    namespace_id VARCHAR(255) NOT NULL REFERENCES cloud_namespaces(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (namespace_id, name)
);
