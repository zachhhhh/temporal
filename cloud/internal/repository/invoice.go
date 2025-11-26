package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Invoice represents an invoice in the database.
type Invoice struct {
	ID                  uuid.UUID
	OrganizationID      uuid.UUID
	InvoiceNumber       string
	PeriodStart         time.Time
	PeriodEnd           time.Time
	LineItems           json.RawMessage
	SubtotalCents       int64
	TaxCents            int64
	CreditsAppliedCents int64
	TotalCents          int64
	Status              string
	StripeInvoiceID     sql.NullString
	PDFURL              sql.NullString
	CreatedAt           time.Time
	PaidAt              sql.NullTime
	DueAt               sql.NullTime
}

// InvoiceLineItem represents a line item on an invoice.
type InvoiceLineItem struct {
	ID             uuid.UUID
	InvoiceID      uuid.UUID
	Description    string
	Quantity       float64
	Unit           string
	UnitPriceCents int64
	AmountCents    int64
	Metadata       json.RawMessage
}

// InvoiceRepository handles invoice data access.
type InvoiceRepository struct {
	db *PostgresDB
}

// NewInvoiceRepository creates a new invoice repository.
func NewInvoiceRepository(db *PostgresDB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

// Create creates a new invoice.
func (r *InvoiceRepository) Create(ctx context.Context, inv *Invoice) error {
	query := `
		INSERT INTO invoices (
			id, organization_id, invoice_number, period_start, period_end,
			line_items, subtotal_cents, tax_cents, credits_applied_cents, total_cents,
			status, stripe_invoice_id, pdf_url, created_at, paid_at, due_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	`
	if inv.ID == uuid.Nil {
		inv.ID = uuid.New()
	}
	if inv.LineItems == nil {
		inv.LineItems = json.RawMessage("[]")
	}
	inv.CreatedAt = time.Now()

	_, err := r.db.DB().ExecContext(ctx, query,
		inv.ID, inv.OrganizationID, inv.InvoiceNumber, inv.PeriodStart, inv.PeriodEnd,
		inv.LineItems, inv.SubtotalCents, inv.TaxCents, inv.CreditsAppliedCents, inv.TotalCents,
		inv.Status, inv.StripeInvoiceID, inv.PDFURL, inv.CreatedAt, inv.PaidAt, inv.DueAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create invoice: %w", err)
	}
	return nil
}

// GetByID retrieves an invoice by ID.
func (r *InvoiceRepository) GetByID(ctx context.Context, id uuid.UUID) (*Invoice, error) {
	query := `
		SELECT id, organization_id, invoice_number, period_start, period_end,
			line_items, subtotal_cents, tax_cents, credits_applied_cents, total_cents,
			status, stripe_invoice_id, pdf_url, created_at, paid_at, due_at
		FROM invoices
		WHERE id = $1
	`
	inv := &Invoice{}
	err := r.db.DB().QueryRowContext(ctx, query, id).Scan(
		&inv.ID, &inv.OrganizationID, &inv.InvoiceNumber, &inv.PeriodStart, &inv.PeriodEnd,
		&inv.LineItems, &inv.SubtotalCents, &inv.TaxCents, &inv.CreditsAppliedCents, &inv.TotalCents,
		&inv.Status, &inv.StripeInvoiceID, &inv.PDFURL, &inv.CreatedAt, &inv.PaidAt, &inv.DueAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get invoice: %w", err)
	}
	return inv, nil
}

// GetByStripeInvoiceID retrieves an invoice by Stripe invoice ID.
func (r *InvoiceRepository) GetByStripeInvoiceID(ctx context.Context, stripeID string) (*Invoice, error) {
	query := `
		SELECT id, organization_id, invoice_number, period_start, period_end,
			line_items, subtotal_cents, tax_cents, credits_applied_cents, total_cents,
			status, stripe_invoice_id, pdf_url, created_at, paid_at, due_at
		FROM invoices
		WHERE stripe_invoice_id = $1
	`
	inv := &Invoice{}
	err := r.db.DB().QueryRowContext(ctx, query, stripeID).Scan(
		&inv.ID, &inv.OrganizationID, &inv.InvoiceNumber, &inv.PeriodStart, &inv.PeriodEnd,
		&inv.LineItems, &inv.SubtotalCents, &inv.TaxCents, &inv.CreditsAppliedCents, &inv.TotalCents,
		&inv.Status, &inv.StripeInvoiceID, &inv.PDFURL, &inv.CreatedAt, &inv.PaidAt, &inv.DueAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get invoice by Stripe ID: %w", err)
	}
	return inv, nil
}

// Update updates an invoice.
func (r *InvoiceRepository) Update(ctx context.Context, inv *Invoice) error {
	query := `
		UPDATE invoices
		SET invoice_number = $2, period_start = $3, period_end = $4,
			line_items = $5, subtotal_cents = $6, tax_cents = $7, credits_applied_cents = $8, total_cents = $9,
			status = $10, stripe_invoice_id = $11, pdf_url = $12, paid_at = $13, due_at = $14
		WHERE id = $1
	`
	_, err := r.db.DB().ExecContext(ctx, query,
		inv.ID, inv.InvoiceNumber, inv.PeriodStart, inv.PeriodEnd,
		inv.LineItems, inv.SubtotalCents, inv.TaxCents, inv.CreditsAppliedCents, inv.TotalCents,
		inv.Status, inv.StripeInvoiceID, inv.PDFURL, inv.PaidAt, inv.DueAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update invoice: %w", err)
	}
	return nil
}

// UpdateStatus updates the invoice status.
func (r *InvoiceRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string, paidAt *time.Time) error {
	query := `UPDATE invoices SET status = $2, paid_at = $3 WHERE id = $1`
	var paidAtVal sql.NullTime
	if paidAt != nil {
		paidAtVal = sql.NullTime{Time: *paidAt, Valid: true}
	}
	_, err := r.db.DB().ExecContext(ctx, query, id, status, paidAtVal)
	if err != nil {
		return fmt.Errorf("failed to update invoice status: %w", err)
	}
	return nil
}

// ListByOrganization lists invoices for an organization.
func (r *InvoiceRepository) ListByOrganization(ctx context.Context, orgID uuid.UUID, limit, offset int) ([]*Invoice, error) {
	query := `
		SELECT id, organization_id, invoice_number, period_start, period_end,
			line_items, subtotal_cents, tax_cents, credits_applied_cents, total_cents,
			status, stripe_invoice_id, pdf_url, created_at, paid_at, due_at
		FROM invoices
		WHERE organization_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.DB().QueryContext(ctx, query, orgID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list invoices: %w", err)
	}
	defer rows.Close()

	var invoices []*Invoice
	for rows.Next() {
		inv := &Invoice{}
		if err := rows.Scan(
			&inv.ID, &inv.OrganizationID, &inv.InvoiceNumber, &inv.PeriodStart, &inv.PeriodEnd,
			&inv.LineItems, &inv.SubtotalCents, &inv.TaxCents, &inv.CreditsAppliedCents, &inv.TotalCents,
			&inv.Status, &inv.StripeInvoiceID, &inv.PDFURL, &inv.CreatedAt, &inv.PaidAt, &inv.DueAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan invoice: %w", err)
		}
		invoices = append(invoices, inv)
	}
	return invoices, nil
}

// ListByStatus lists invoices by status.
func (r *InvoiceRepository) ListByStatus(ctx context.Context, status string, limit, offset int) ([]*Invoice, error) {
	query := `
		SELECT id, organization_id, invoice_number, period_start, period_end,
			line_items, subtotal_cents, tax_cents, credits_applied_cents, total_cents,
			status, stripe_invoice_id, pdf_url, created_at, paid_at, due_at
		FROM invoices
		WHERE status = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.DB().QueryContext(ctx, query, status, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list invoices by status: %w", err)
	}
	defer rows.Close()

	var invoices []*Invoice
	for rows.Next() {
		inv := &Invoice{}
		if err := rows.Scan(
			&inv.ID, &inv.OrganizationID, &inv.InvoiceNumber, &inv.PeriodStart, &inv.PeriodEnd,
			&inv.LineItems, &inv.SubtotalCents, &inv.TaxCents, &inv.CreditsAppliedCents, &inv.TotalCents,
			&inv.Status, &inv.StripeInvoiceID, &inv.PDFURL, &inv.CreatedAt, &inv.PaidAt, &inv.DueAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan invoice: %w", err)
		}
		invoices = append(invoices, inv)
	}
	return invoices, nil
}

// GenerateInvoiceNumber generates a unique invoice number.
func (r *InvoiceRepository) GenerateInvoiceNumber(ctx context.Context, orgID uuid.UUID) (string, error) {
	query := `SELECT COUNT(*) + 1 FROM invoices WHERE organization_id = $1`
	var count int
	if err := r.db.DB().QueryRowContext(ctx, query, orgID).Scan(&count); err != nil {
		return "", fmt.Errorf("failed to generate invoice number: %w", err)
	}
	return fmt.Sprintf("INV-%s-%04d", orgID.String()[:8], count), nil
}
