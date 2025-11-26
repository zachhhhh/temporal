-- Connectivity rules (IP allowlist, PrivateLink, etc.)
CREATE TABLE connectivity_rules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL, -- 'ip_allowlist', 'private_link', 'vpc_peering'
    config JSONB NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_connectivity_rules_org ON connectivity_rules(organization_id);
CREATE INDEX idx_connectivity_rules_type ON connectivity_rules(type);

CREATE TRIGGER update_connectivity_rules_updated_at
    BEFORE UPDATE ON connectivity_rules
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Namespace connectivity bindings
CREATE TABLE namespace_connectivity_bindings (
    namespace_id VARCHAR(255) NOT NULL REFERENCES cloud_namespaces(id) ON DELETE CASCADE,
    connectivity_rule_id UUID NOT NULL REFERENCES connectivity_rules(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (namespace_id, connectivity_rule_id)
);

CREATE INDEX idx_ns_conn_bindings_rule ON namespace_connectivity_bindings(connectivity_rule_id);
