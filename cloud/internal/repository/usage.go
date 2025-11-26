package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// UsageRecord represents a usage record in the database.
type UsageRecord struct {
	ID                    uuid.UUID
	OrganizationID        uuid.UUID
	NamespaceID           string
	PeriodStart           time.Time
	PeriodEnd             time.Time
	ActionCount           int64
	WorkflowStarted       int64
	WorkflowReset         int64
	TimerStarted          int64
	SignalSent            int64
	QueryReceived         int64
	UpdateReceived        int64
	ActivityStarted       int64
	ActivityHeartbeat     int64
	LocalActivityBatch    int64
	ChildWorkflowStarted  int64
	ScheduleExecution     int64
	NexusOperation        int64
	SearchAttributeUpsert int64
	SideEffectRecorded    int64
	WorkflowExported      int64
	ActiveStorageGBH      decimal.Decimal
	RetainedStorageGBH    decimal.Decimal
	CreatedAt             time.Time
}

// UsageAggregate represents an aggregated usage record.
type UsageAggregate struct {
	ID                 uuid.UUID
	OrganizationID     uuid.UUID
	NamespaceID        sql.NullString
	PeriodType         string
	PeriodStart        time.Time
	PeriodEnd          time.Time
	TotalActions       int64
	ActiveStorageGBH   decimal.Decimal
	RetainedStorageGBH decimal.Decimal
	ActionBreakdown    []byte
	CreatedAt          time.Time
}

// UsageRepository handles usage data access.
type UsageRepository struct {
	db *PostgresDB
}

// NewUsageRepository creates a new usage repository.
func NewUsageRepository(db *PostgresDB) *UsageRepository {
	return &UsageRepository{db: db}
}

// Create creates a new usage record.
func (r *UsageRepository) Create(ctx context.Context, record *UsageRecord) error {
	query := `
		INSERT INTO usage_records (
			id, organization_id, namespace_id, period_start, period_end,
			action_count, workflow_started, workflow_reset, timer_started,
			signal_sent, query_received, update_received, activity_started,
			activity_heartbeat, local_activity_batch, child_workflow_started,
			schedule_execution, nexus_operation, search_attribute_upsert,
			side_effect_recorded, workflow_exported, active_storage_gbh,
			retained_storage_gbh, created_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24
		)
		ON CONFLICT (organization_id, namespace_id, period_start) DO UPDATE SET
			action_count = usage_records.action_count + EXCLUDED.action_count,
			workflow_started = usage_records.workflow_started + EXCLUDED.workflow_started,
			workflow_reset = usage_records.workflow_reset + EXCLUDED.workflow_reset,
			timer_started = usage_records.timer_started + EXCLUDED.timer_started,
			signal_sent = usage_records.signal_sent + EXCLUDED.signal_sent,
			query_received = usage_records.query_received + EXCLUDED.query_received,
			update_received = usage_records.update_received + EXCLUDED.update_received,
			activity_started = usage_records.activity_started + EXCLUDED.activity_started,
			activity_heartbeat = usage_records.activity_heartbeat + EXCLUDED.activity_heartbeat,
			local_activity_batch = usage_records.local_activity_batch + EXCLUDED.local_activity_batch,
			child_workflow_started = usage_records.child_workflow_started + EXCLUDED.child_workflow_started,
			schedule_execution = usage_records.schedule_execution + EXCLUDED.schedule_execution,
			nexus_operation = usage_records.nexus_operation + EXCLUDED.nexus_operation,
			search_attribute_upsert = usage_records.search_attribute_upsert + EXCLUDED.search_attribute_upsert,
			side_effect_recorded = usage_records.side_effect_recorded + EXCLUDED.side_effect_recorded,
			workflow_exported = usage_records.workflow_exported + EXCLUDED.workflow_exported,
			active_storage_gbh = usage_records.active_storage_gbh + EXCLUDED.active_storage_gbh,
			retained_storage_gbh = usage_records.retained_storage_gbh + EXCLUDED.retained_storage_gbh
	`
	if record.ID == uuid.Nil {
		record.ID = uuid.New()
	}
	record.CreatedAt = time.Now()

	_, err := r.db.DB().ExecContext(ctx, query,
		record.ID, record.OrganizationID, record.NamespaceID, record.PeriodStart, record.PeriodEnd,
		record.ActionCount, record.WorkflowStarted, record.WorkflowReset, record.TimerStarted,
		record.SignalSent, record.QueryReceived, record.UpdateReceived, record.ActivityStarted,
		record.ActivityHeartbeat, record.LocalActivityBatch, record.ChildWorkflowStarted,
		record.ScheduleExecution, record.NexusOperation, record.SearchAttributeUpsert,
		record.SideEffectRecorded, record.WorkflowExported, record.ActiveStorageGBH,
		record.RetainedStorageGBH, record.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create usage record: %w", err)
	}
	return nil
}

// GetByOrganizationAndPeriod retrieves usage records for an organization in a period.
func (r *UsageRepository) GetByOrganizationAndPeriod(ctx context.Context, orgID uuid.UUID, start, end time.Time) ([]*UsageRecord, error) {
	query := `
		SELECT id, organization_id, namespace_id, period_start, period_end,
			action_count, workflow_started, workflow_reset, timer_started,
			signal_sent, query_received, update_received, activity_started,
			activity_heartbeat, local_activity_batch, child_workflow_started,
			schedule_execution, nexus_operation, search_attribute_upsert,
			side_effect_recorded, workflow_exported, active_storage_gbh,
			retained_storage_gbh, created_at
		FROM usage_records
		WHERE organization_id = $1 AND period_start >= $2 AND period_end <= $3
		ORDER BY period_start
	`
	rows, err := r.db.DB().QueryContext(ctx, query, orgID, start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to get usage records: %w", err)
	}
	defer rows.Close()

	var records []*UsageRecord
	for rows.Next() {
		record := &UsageRecord{}
		if err := rows.Scan(
			&record.ID, &record.OrganizationID, &record.NamespaceID, &record.PeriodStart, &record.PeriodEnd,
			&record.ActionCount, &record.WorkflowStarted, &record.WorkflowReset, &record.TimerStarted,
			&record.SignalSent, &record.QueryReceived, &record.UpdateReceived, &record.ActivityStarted,
			&record.ActivityHeartbeat, &record.LocalActivityBatch, &record.ChildWorkflowStarted,
			&record.ScheduleExecution, &record.NexusOperation, &record.SearchAttributeUpsert,
			&record.SideEffectRecorded, &record.WorkflowExported, &record.ActiveStorageGBH,
			&record.RetainedStorageGBH, &record.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan usage record: %w", err)
		}
		records = append(records, record)
	}
	return records, nil
}

// GetSummaryByOrganization retrieves aggregated usage for an organization in a period.
func (r *UsageRepository) GetSummaryByOrganization(ctx context.Context, orgID uuid.UUID, start, end time.Time) (*UsageRecord, error) {
	query := `
		SELECT 
			COALESCE(SUM(action_count), 0) as action_count,
			COALESCE(SUM(workflow_started), 0) as workflow_started,
			COALESCE(SUM(workflow_reset), 0) as workflow_reset,
			COALESCE(SUM(timer_started), 0) as timer_started,
			COALESCE(SUM(signal_sent), 0) as signal_sent,
			COALESCE(SUM(query_received), 0) as query_received,
			COALESCE(SUM(update_received), 0) as update_received,
			COALESCE(SUM(activity_started), 0) as activity_started,
			COALESCE(SUM(activity_heartbeat), 0) as activity_heartbeat,
			COALESCE(SUM(local_activity_batch), 0) as local_activity_batch,
			COALESCE(SUM(child_workflow_started), 0) as child_workflow_started,
			COALESCE(SUM(schedule_execution), 0) as schedule_execution,
			COALESCE(SUM(nexus_operation), 0) as nexus_operation,
			COALESCE(SUM(search_attribute_upsert), 0) as search_attribute_upsert,
			COALESCE(SUM(side_effect_recorded), 0) as side_effect_recorded,
			COALESCE(SUM(workflow_exported), 0) as workflow_exported,
			COALESCE(SUM(active_storage_gbh), 0) as active_storage_gbh,
			COALESCE(SUM(retained_storage_gbh), 0) as retained_storage_gbh
		FROM usage_records
		WHERE organization_id = $1 AND period_start >= $2 AND period_end <= $3
	`
	record := &UsageRecord{
		OrganizationID: orgID,
		PeriodStart:    start,
		PeriodEnd:      end,
	}
	err := r.db.DB().QueryRowContext(ctx, query, orgID, start, end).Scan(
		&record.ActionCount, &record.WorkflowStarted, &record.WorkflowReset, &record.TimerStarted,
		&record.SignalSent, &record.QueryReceived, &record.UpdateReceived, &record.ActivityStarted,
		&record.ActivityHeartbeat, &record.LocalActivityBatch, &record.ChildWorkflowStarted,
		&record.ScheduleExecution, &record.NexusOperation, &record.SearchAttributeUpsert,
		&record.SideEffectRecorded, &record.WorkflowExported, &record.ActiveStorageGBH,
		&record.RetainedStorageGBH,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get usage summary: %w", err)
	}
	return record, nil
}

// GetSummaryByNamespace retrieves aggregated usage for a namespace in a period.
func (r *UsageRepository) GetSummaryByNamespace(ctx context.Context, namespaceID string, start, end time.Time) (*UsageRecord, error) {
	query := `
		SELECT 
			COALESCE(SUM(action_count), 0) as action_count,
			COALESCE(SUM(active_storage_gbh), 0) as active_storage_gbh,
			COALESCE(SUM(retained_storage_gbh), 0) as retained_storage_gbh
		FROM usage_records
		WHERE namespace_id = $1 AND period_start >= $2 AND period_end <= $3
	`
	record := &UsageRecord{
		NamespaceID: namespaceID,
		PeriodStart: start,
		PeriodEnd:   end,
	}
	err := r.db.DB().QueryRowContext(ctx, query, namespaceID, start, end).Scan(
		&record.ActionCount, &record.ActiveStorageGBH, &record.RetainedStorageGBH,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get namespace usage summary: %w", err)
	}
	return record, nil
}

// CreateAggregate creates a usage aggregate.
func (r *UsageRepository) CreateAggregate(ctx context.Context, agg *UsageAggregate) error {
	query := `
		INSERT INTO usage_aggregates (
			id, organization_id, namespace_id, period_type, period_start, period_end,
			total_actions, active_storage_gbh, retained_storage_gbh, action_breakdown, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (organization_id, namespace_id, period_type, period_start) DO UPDATE SET
			total_actions = EXCLUDED.total_actions,
			active_storage_gbh = EXCLUDED.active_storage_gbh,
			retained_storage_gbh = EXCLUDED.retained_storage_gbh,
			action_breakdown = EXCLUDED.action_breakdown
	`
	if agg.ID == uuid.Nil {
		agg.ID = uuid.New()
	}
	agg.CreatedAt = time.Now()

	_, err := r.db.DB().ExecContext(ctx, query,
		agg.ID, agg.OrganizationID, agg.NamespaceID, agg.PeriodType, agg.PeriodStart, agg.PeriodEnd,
		agg.TotalActions, agg.ActiveStorageGBH, agg.RetainedStorageGBH, agg.ActionBreakdown, agg.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create usage aggregate: %w", err)
	}
	return nil
}
