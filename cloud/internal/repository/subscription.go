package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Subscription represents a subscription in the database.
type Subscription struct {
	ID                   uuid.UUID
	OrganizationID       uuid.UUID
	Plan                 string
	Status               string
	ActionsIncluded      int64
	ActiveStorageGB      decimal.Decimal
	RetainedStorageGB    decimal.Decimal
	StripeCustomerID     sql.NullString
	StripeSubscriptionID sql.NullString
	CurrentPeriodStart   sql.NullTime
	CurrentPeriodEnd     sql.NullTime
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

// SubscriptionRepository handles subscription data access.
type SubscriptionRepository struct {
	db *PostgresDB
}

// NewSubscriptionRepository creates a new subscription repository.
func NewSubscriptionRepository(db *PostgresDB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

// Create creates a new subscription.
func (r *SubscriptionRepository) Create(ctx context.Context, sub *Subscription) error {
	query := `
		INSERT INTO subscriptions (
			id, organization_id, plan, status, actions_included,
			active_storage_gb, retained_storage_gb, stripe_customer_id, stripe_subscription_id,
			current_period_start, current_period_end, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`
	if sub.ID == uuid.Nil {
		sub.ID = uuid.New()
	}
	now := time.Now()
	sub.CreatedAt = now
	sub.UpdatedAt = now

	_, err := r.db.DB().ExecContext(ctx, query,
		sub.ID, sub.OrganizationID, sub.Plan, sub.Status, sub.ActionsIncluded,
		sub.ActiveStorageGB, sub.RetainedStorageGB, sub.StripeCustomerID, sub.StripeSubscriptionID,
		sub.CurrentPeriodStart, sub.CurrentPeriodEnd, sub.CreatedAt, sub.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create subscription: %w", err)
	}
	return nil
}

// GetByOrganizationID retrieves a subscription by organization ID.
func (r *SubscriptionRepository) GetByOrganizationID(ctx context.Context, orgID uuid.UUID) (*Subscription, error) {
	query := `
		SELECT id, organization_id, plan, status, actions_included,
			active_storage_gb, retained_storage_gb, stripe_customer_id, stripe_subscription_id,
			current_period_start, current_period_end, created_at, updated_at
		FROM subscriptions
		WHERE organization_id = $1
	`
	sub := &Subscription{}
	err := r.db.DB().QueryRowContext(ctx, query, orgID).Scan(
		&sub.ID, &sub.OrganizationID, &sub.Plan, &sub.Status, &sub.ActionsIncluded,
		&sub.ActiveStorageGB, &sub.RetainedStorageGB, &sub.StripeCustomerID, &sub.StripeSubscriptionID,
		&sub.CurrentPeriodStart, &sub.CurrentPeriodEnd, &sub.CreatedAt, &sub.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription: %w", err)
	}
	return sub, nil
}

// GetByStripeCustomerID retrieves a subscription by Stripe customer ID.
func (r *SubscriptionRepository) GetByStripeCustomerID(ctx context.Context, customerID string) (*Subscription, error) {
	query := `
		SELECT id, organization_id, plan, status, actions_included,
			active_storage_gb, retained_storage_gb, stripe_customer_id, stripe_subscription_id,
			current_period_start, current_period_end, created_at, updated_at
		FROM subscriptions
		WHERE stripe_customer_id = $1
	`
	sub := &Subscription{}
	err := r.db.DB().QueryRowContext(ctx, query, customerID).Scan(
		&sub.ID, &sub.OrganizationID, &sub.Plan, &sub.Status, &sub.ActionsIncluded,
		&sub.ActiveStorageGB, &sub.RetainedStorageGB, &sub.StripeCustomerID, &sub.StripeSubscriptionID,
		&sub.CurrentPeriodStart, &sub.CurrentPeriodEnd, &sub.CreatedAt, &sub.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription by Stripe customer: %w", err)
	}
	return sub, nil
}

// Update updates a subscription.
func (r *SubscriptionRepository) Update(ctx context.Context, sub *Subscription) error {
	query := `
		UPDATE subscriptions
		SET plan = $2, status = $3, actions_included = $4,
			active_storage_gb = $5, retained_storage_gb = $6,
			stripe_customer_id = $7, stripe_subscription_id = $8,
			current_period_start = $9, current_period_end = $10, updated_at = $11
		WHERE id = $1
	`
	sub.UpdatedAt = time.Now()
	_, err := r.db.DB().ExecContext(ctx, query,
		sub.ID, sub.Plan, sub.Status, sub.ActionsIncluded,
		sub.ActiveStorageGB, sub.RetainedStorageGB,
		sub.StripeCustomerID, sub.StripeSubscriptionID,
		sub.CurrentPeriodStart, sub.CurrentPeriodEnd, sub.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update subscription: %w", err)
	}
	return nil
}

// UpdateStatus updates the subscription status.
func (r *SubscriptionRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	query := `UPDATE subscriptions SET status = $2, updated_at = $3 WHERE id = $1`
	_, err := r.db.DB().ExecContext(ctx, query, id, status, time.Now())
	if err != nil {
		return fmt.Errorf("failed to update subscription status: %w", err)
	}
	return nil
}

// Delete deletes a subscription.
func (r *SubscriptionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM subscriptions WHERE id = $1`
	_, err := r.db.DB().ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete subscription: %w", err)
	}
	return nil
}

// GetPlanLimits returns the limits for a plan tier.
func GetPlanLimits(plan string) (actionsIncluded int64, activeStorageGB, retainedStorageGB decimal.Decimal) {
	switch plan {
	case "free":
		return 100000, decimal.NewFromFloat(0.1), decimal.NewFromFloat(4)
	case "essentials":
		return 1000000, decimal.NewFromFloat(1), decimal.NewFromFloat(40)
	case "business":
		return 2500000, decimal.NewFromFloat(2.5), decimal.NewFromFloat(100)
	case "enterprise":
		return 10000000, decimal.NewFromFloat(10), decimal.NewFromFloat(400)
	case "mission_critical":
		return 100000000, decimal.NewFromFloat(100), decimal.NewFromFloat(4000)
	default:
		return 100000, decimal.NewFromFloat(0.1), decimal.NewFromFloat(4)
	}
}
