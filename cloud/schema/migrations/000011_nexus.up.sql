-- Nexus endpoints
CREATE TABLE nexus_endpoints (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    target_namespace_id VARCHAR(255) NOT NULL REFERENCES cloud_namespaces(id) ON DELETE CASCADE,
    handler_name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(organization_id, name)
);

CREATE INDEX idx_nexus_endpoints_org ON nexus_endpoints(organization_id);
CREATE INDEX idx_nexus_endpoints_target ON nexus_endpoints(target_namespace_id);

CREATE TRIGGER update_nexus_endpoints_updated_at
    BEFORE UPDATE ON nexus_endpoints
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Nexus endpoint allowlist (which namespaces can call this endpoint)
CREATE TABLE nexus_endpoint_allowlist (
    endpoint_id UUID NOT NULL REFERENCES nexus_endpoints(id) ON DELETE CASCADE,
    caller_namespace_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (endpoint_id, caller_namespace_id)
);

CREATE INDEX idx_nexus_allowlist_caller ON nexus_endpoint_allowlist(caller_namespace_id);
