// Package repository provides data access layer implementations.
package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"go.temporal.io/cloud/internal/config"
)

// PostgresDB wraps a PostgreSQL database connection.
type PostgresDB struct {
	db *sql.DB
}

// NewPostgresDB creates a new PostgreSQL database connection.
func NewPostgresDB(cfg config.DatabaseConfig) (*PostgresDB, error) {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresDB{db: db}, nil
}

// Close closes the database connection.
func (p *PostgresDB) Close() error {
	return p.db.Close()
}

// Ping pings the database.
func (p *PostgresDB) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return p.db.PingContext(ctx)
}

// DB returns the underlying database connection.
func (p *PostgresDB) DB() *sql.DB {
	return p.db
}

// Repositories holds all repository instances.
type Repositories struct {
	Organizations *OrganizationRepository
	Namespaces    *NamespaceRepository
	Users         *UserRepository
	Subscriptions *SubscriptionRepository
	Usage         *UsageRepository
	Invoices      *InvoiceRepository
	APIKeys       *APIKeyRepository
	Audit         *AuditRepository
}

// NewRepositories creates all repository instances.
func NewRepositories(db *PostgresDB) *Repositories {
	return &Repositories{
		Organizations: NewOrganizationRepository(db),
		Namespaces:    NewNamespaceRepository(db),
		Users:         NewUserRepository(db),
		Subscriptions: NewSubscriptionRepository(db),
		Usage:         NewUsageRepository(db),
		Invoices:      NewInvoiceRepository(db),
		APIKeys:       NewAPIKeyRepository(db),
		Audit:         NewAuditRepository(db),
	}
}
