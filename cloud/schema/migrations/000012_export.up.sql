-- Export sinks (S3, GCS, etc.)
CREATE TABLE export_sinks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    namespace_id VARCHAR(255) NOT NULL REFERENCES cloud_namespaces(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    sink_type VARCHAR(50) NOT NULL, -- 's3', 'gcs'
    config JSONB NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    last_export_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_export_sinks_namespace ON export_sinks(namespace_id);
CREATE INDEX idx_export_sinks_type ON export_sinks(sink_type);

CREATE TRIGGER update_export_sinks_updated_at
    BEFORE UPDATE ON export_sinks
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Export jobs (tracking individual export operations)
CREATE TABLE export_jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sink_id UUID NOT NULL REFERENCES export_sinks(id) ON DELETE CASCADE,
    status VARCHAR(50) NOT NULL DEFAULT 'pending', -- 'pending', 'running', 'completed', 'failed'
    workflows_exported BIGINT NOT NULL DEFAULT 0,
    bytes_exported BIGINT NOT NULL DEFAULT 0,
    error_message TEXT,
    started_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_export_jobs_sink ON export_jobs(sink_id, created_at DESC);
CREATE INDEX idx_export_jobs_status ON export_jobs(status);
