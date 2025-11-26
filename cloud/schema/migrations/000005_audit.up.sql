-- Audit result enum
CREATE TYPE audit_result AS ENUM ('success', 'failure', 'denied');

-- Audit events table (partitioned by created_at for efficient querying)
CREATE TABLE audit_events (
    id UUID NOT NULL,
    organization_id UUID NOT NULL,
    -- Actor
    actor_type VARCHAR(50) NOT NULL,
    actor_id VARCHAR(255) NOT NULL,
    actor_email VARCHAR(255),
    actor_name VARCHAR(255),
    -- Action
    action VARCHAR(100) NOT NULL,
    result audit_result NOT NULL DEFAULT 'success',
    -- Resource
    resource_type VARCHAR(100) NOT NULL,
    resource_id VARCHAR(255),
    resource_name VARCHAR(255),
    -- Request metadata
    request_id VARCHAR(255),
    ip_address INET,
    user_agent TEXT,
    method VARCHAR(20),
    path TEXT,
    -- Details
    details JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id, created_at)
) PARTITION BY RANGE (created_at);

-- Create partitions for the next 12 months
CREATE TABLE audit_events_2024_01 PARTITION OF audit_events
    FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');
CREATE TABLE audit_events_2024_02 PARTITION OF audit_events
    FOR VALUES FROM ('2024-02-01') TO ('2024-03-01');
CREATE TABLE audit_events_2024_03 PARTITION OF audit_events
    FOR VALUES FROM ('2024-03-01') TO ('2024-04-01');
CREATE TABLE audit_events_2024_04 PARTITION OF audit_events
    FOR VALUES FROM ('2024-04-01') TO ('2024-05-01');
CREATE TABLE audit_events_2024_05 PARTITION OF audit_events
    FOR VALUES FROM ('2024-05-01') TO ('2024-06-01');
CREATE TABLE audit_events_2024_06 PARTITION OF audit_events
    FOR VALUES FROM ('2024-06-01') TO ('2024-07-01');
CREATE TABLE audit_events_2024_07 PARTITION OF audit_events
    FOR VALUES FROM ('2024-07-01') TO ('2024-08-01');
CREATE TABLE audit_events_2024_08 PARTITION OF audit_events
    FOR VALUES FROM ('2024-08-01') TO ('2024-09-01');
CREATE TABLE audit_events_2024_09 PARTITION OF audit_events
    FOR VALUES FROM ('2024-09-01') TO ('2024-10-01');
CREATE TABLE audit_events_2024_10 PARTITION OF audit_events
    FOR VALUES FROM ('2024-10-01') TO ('2024-11-01');
CREATE TABLE audit_events_2024_11 PARTITION OF audit_events
    FOR VALUES FROM ('2024-11-01') TO ('2024-12-01');
CREATE TABLE audit_events_2024_12 PARTITION OF audit_events
    FOR VALUES FROM ('2024-12-01') TO ('2025-01-01');
CREATE TABLE audit_events_2025_01 PARTITION OF audit_events
    FOR VALUES FROM ('2025-01-01') TO ('2025-02-01');
CREATE TABLE audit_events_2025_02 PARTITION OF audit_events
    FOR VALUES FROM ('2025-02-01') TO ('2025-03-01');
CREATE TABLE audit_events_2025_03 PARTITION OF audit_events
    FOR VALUES FROM ('2025-03-01') TO ('2025-04-01');
CREATE TABLE audit_events_2025_04 PARTITION OF audit_events
    FOR VALUES FROM ('2025-04-01') TO ('2025-05-01');
CREATE TABLE audit_events_2025_05 PARTITION OF audit_events
    FOR VALUES FROM ('2025-05-01') TO ('2025-06-01');
CREATE TABLE audit_events_2025_06 PARTITION OF audit_events
    FOR VALUES FROM ('2025-06-01') TO ('2025-07-01');
CREATE TABLE audit_events_2025_07 PARTITION OF audit_events
    FOR VALUES FROM ('2025-07-01') TO ('2025-08-01');
CREATE TABLE audit_events_2025_08 PARTITION OF audit_events
    FOR VALUES FROM ('2025-08-01') TO ('2025-09-01');
CREATE TABLE audit_events_2025_09 PARTITION OF audit_events
    FOR VALUES FROM ('2025-09-01') TO ('2025-10-01');
CREATE TABLE audit_events_2025_10 PARTITION OF audit_events
    FOR VALUES FROM ('2025-10-01') TO ('2025-11-01');
CREATE TABLE audit_events_2025_11 PARTITION OF audit_events
    FOR VALUES FROM ('2025-11-01') TO ('2025-12-01');
CREATE TABLE audit_events_2025_12 PARTITION OF audit_events
    FOR VALUES FROM ('2025-12-01') TO ('2026-01-01');

CREATE INDEX idx_audit_org_time ON audit_events(organization_id, created_at DESC);
CREATE INDEX idx_audit_actor ON audit_events(actor_id, created_at DESC);
CREATE INDEX idx_audit_action ON audit_events(action, created_at DESC);
CREATE INDEX idx_audit_resource ON audit_events(resource_type, resource_id, created_at DESC);
