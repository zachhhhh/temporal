-- User groups
CREATE TABLE user_groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    -- SCIM integration
    scim_id VARCHAR(255),
    scim_external_id VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(organization_id, name)
);

CREATE INDEX idx_user_groups_org ON user_groups(organization_id);
CREATE INDEX idx_user_groups_scim ON user_groups(scim_id);

CREATE TRIGGER update_user_groups_updated_at
    BEFORE UPDATE ON user_groups
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- User group members
CREATE TABLE user_group_members (
    group_id UUID NOT NULL REFERENCES user_groups(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (group_id, user_id)
);

CREATE INDEX idx_user_group_members_user ON user_group_members(user_id);

-- User group namespace permissions
CREATE TABLE user_group_namespace_permissions (
    group_id UUID NOT NULL REFERENCES user_groups(id) ON DELETE CASCADE,
    namespace_id VARCHAR(255) NOT NULL,
    permission VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (group_id, namespace_id)
);

CREATE INDEX idx_user_group_ns_perms_namespace ON user_group_namespace_permissions(namespace_id);

-- SAML configuration
CREATE TABLE saml_configurations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE UNIQUE,
    enabled BOOLEAN NOT NULL DEFAULT FALSE,
    idp_entity_id VARCHAR(255) NOT NULL,
    idp_sso_url TEXT NOT NULL,
    idp_certificate TEXT NOT NULL,
    sp_entity_id VARCHAR(255) NOT NULL,
    sp_acs_url TEXT NOT NULL,
    attribute_mapping JSONB DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_saml_config_org ON saml_configurations(organization_id);

CREATE TRIGGER update_saml_configurations_updated_at
    BEFORE UPDATE ON saml_configurations
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- SCIM configuration
CREATE TABLE scim_configurations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE UNIQUE,
    enabled BOOLEAN NOT NULL DEFAULT FALSE,
    token_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_scim_config_org ON scim_configurations(organization_id);

CREATE TRIGGER update_scim_configurations_updated_at
    BEFORE UPDATE ON scim_configurations
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
