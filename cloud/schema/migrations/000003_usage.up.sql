-- Usage records table (hourly granularity)
CREATE TABLE usage_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    namespace_id VARCHAR(255) NOT NULL,
    period_start TIMESTAMPTZ NOT NULL,
    period_end TIMESTAMPTZ NOT NULL,
    -- Action counts
    action_count BIGINT NOT NULL DEFAULT 0,
    workflow_started BIGINT NOT NULL DEFAULT 0,
    workflow_reset BIGINT NOT NULL DEFAULT 0,
    timer_started BIGINT NOT NULL DEFAULT 0,
    signal_sent BIGINT NOT NULL DEFAULT 0,
    query_received BIGINT NOT NULL DEFAULT 0,
    update_received BIGINT NOT NULL DEFAULT 0,
    activity_started BIGINT NOT NULL DEFAULT 0,
    activity_heartbeat BIGINT NOT NULL DEFAULT 0,
    local_activity_batch BIGINT NOT NULL DEFAULT 0,
    child_workflow_started BIGINT NOT NULL DEFAULT 0,
    schedule_execution BIGINT NOT NULL DEFAULT 0,
    nexus_operation BIGINT NOT NULL DEFAULT 0,
    search_attribute_upsert BIGINT NOT NULL DEFAULT 0,
    side_effect_recorded BIGINT NOT NULL DEFAULT 0,
    workflow_exported BIGINT NOT NULL DEFAULT 0,
    -- Storage (GB-hours)
    active_storage_gbh DECIMAL(20,6) NOT NULL DEFAULT 0,
    retained_storage_gbh DECIMAL(20,6) NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(organization_id, namespace_id, period_start)
);

CREATE INDEX idx_usage_org_period ON usage_records(organization_id, period_start);
CREATE INDEX idx_usage_namespace_period ON usage_records(namespace_id, period_start);
CREATE INDEX idx_usage_period_start ON usage_records(period_start);

-- Aggregated usage (daily/monthly rollups)
CREATE TABLE usage_aggregates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    namespace_id VARCHAR(255),
    period_type VARCHAR(20) NOT NULL, -- 'daily', 'monthly'
    period_start TIMESTAMPTZ NOT NULL,
    period_end TIMESTAMPTZ NOT NULL,
    total_actions BIGINT NOT NULL DEFAULT 0,
    active_storage_gbh DECIMAL(20,6) NOT NULL DEFAULT 0,
    retained_storage_gbh DECIMAL(20,6) NOT NULL DEFAULT 0,
    action_breakdown JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(organization_id, namespace_id, period_type, period_start)
);

CREATE INDEX idx_usage_agg_org_period ON usage_aggregates(organization_id, period_type, period_start);
