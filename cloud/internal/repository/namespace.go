package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Namespace represents a cloud namespace in the database.
type Namespace struct {
	ID                      string
	OrganizationID          uuid.UUID
	Name                    string
	Region                  string
	ClusterID               sql.NullString
	State                   string
	RetentionDays           int
	DeletionProtected       bool
	HAEnabled               bool
	StandbyRegion           sql.NullString
	CodecEndpoint           sql.NullString
	CodecPassToken          bool
	CodecIncludeCredentials bool
	GRPCEndpoint            sql.NullString
	WebEndpoint             sql.NullString
	MetricsEndpoint         sql.NullString
	Tags                    json.RawMessage
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

// NamespaceCertificate represents a namespace certificate.
type NamespaceCertificate struct {
	ID             uuid.UUID
	NamespaceID    string
	CertificatePEM string
	Fingerprint    string
	Issuer         sql.NullString
	Subject        sql.NullString
	ExpiresAt      time.Time
	CreatedAt      time.Time
}

// NamespaceCertificateFilter represents a certificate filter.
type NamespaceCertificateFilter struct {
	ID                     uuid.UUID
	NamespaceID            string
	CommonName             sql.NullString
	Organization           sql.NullString
	OrganizationalUnit     sql.NullString
	SubjectAlternativeName sql.NullString
	CreatedAt              time.Time
}

// NamespaceSearchAttribute represents a search attribute.
type NamespaceSearchAttribute struct {
	NamespaceID string
	Name        string
	Type        string
	CreatedAt   time.Time
}

// NamespaceRepository handles namespace data access.
type NamespaceRepository struct {
	db *PostgresDB
}

// NewNamespaceRepository creates a new namespace repository.
func NewNamespaceRepository(db *PostgresDB) *NamespaceRepository {
	return &NamespaceRepository{db: db}
}

// Create creates a new namespace.
func (r *NamespaceRepository) Create(ctx context.Context, ns *Namespace) error {
	query := `
		INSERT INTO cloud_namespaces (
			id, organization_id, name, region, cluster_id, state,
			retention_days, deletion_protected, ha_enabled, standby_region,
			codec_endpoint, codec_pass_token, codec_include_credentials,
			grpc_endpoint, web_endpoint, metrics_endpoint, tags,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
		)
	`
	if ns.Tags == nil {
		ns.Tags = json.RawMessage("{}")
	}
	now := time.Now()
	ns.CreatedAt = now
	ns.UpdatedAt = now

	_, err := r.db.DB().ExecContext(ctx, query,
		ns.ID, ns.OrganizationID, ns.Name, ns.Region, ns.ClusterID, ns.State,
		ns.RetentionDays, ns.DeletionProtected, ns.HAEnabled, ns.StandbyRegion,
		ns.CodecEndpoint, ns.CodecPassToken, ns.CodecIncludeCredentials,
		ns.GRPCEndpoint, ns.WebEndpoint, ns.MetricsEndpoint, ns.Tags,
		ns.CreatedAt, ns.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create namespace: %w", err)
	}
	return nil
}

// GetByID retrieves a namespace by ID.
func (r *NamespaceRepository) GetByID(ctx context.Context, id string) (*Namespace, error) {
	query := `
		SELECT id, organization_id, name, region, cluster_id, state,
			retention_days, deletion_protected, ha_enabled, standby_region,
			codec_endpoint, codec_pass_token, codec_include_credentials,
			grpc_endpoint, web_endpoint, metrics_endpoint, tags,
			created_at, updated_at
		FROM cloud_namespaces
		WHERE id = $1
	`
	ns := &Namespace{}
	err := r.db.DB().QueryRowContext(ctx, query, id).Scan(
		&ns.ID, &ns.OrganizationID, &ns.Name, &ns.Region, &ns.ClusterID, &ns.State,
		&ns.RetentionDays, &ns.DeletionProtected, &ns.HAEnabled, &ns.StandbyRegion,
		&ns.CodecEndpoint, &ns.CodecPassToken, &ns.CodecIncludeCredentials,
		&ns.GRPCEndpoint, &ns.WebEndpoint, &ns.MetricsEndpoint, &ns.Tags,
		&ns.CreatedAt, &ns.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get namespace: %w", err)
	}
	return ns, nil
}

// GetByOrgAndName retrieves a namespace by organization and name.
func (r *NamespaceRepository) GetByOrgAndName(ctx context.Context, orgID uuid.UUID, name string) (*Namespace, error) {
	query := `
		SELECT id, organization_id, name, region, cluster_id, state,
			retention_days, deletion_protected, ha_enabled, standby_region,
			codec_endpoint, codec_pass_token, codec_include_credentials,
			grpc_endpoint, web_endpoint, metrics_endpoint, tags,
			created_at, updated_at
		FROM cloud_namespaces
		WHERE organization_id = $1 AND name = $2
	`
	ns := &Namespace{}
	err := r.db.DB().QueryRowContext(ctx, query, orgID, name).Scan(
		&ns.ID, &ns.OrganizationID, &ns.Name, &ns.Region, &ns.ClusterID, &ns.State,
		&ns.RetentionDays, &ns.DeletionProtected, &ns.HAEnabled, &ns.StandbyRegion,
		&ns.CodecEndpoint, &ns.CodecPassToken, &ns.CodecIncludeCredentials,
		&ns.GRPCEndpoint, &ns.WebEndpoint, &ns.MetricsEndpoint, &ns.Tags,
		&ns.CreatedAt, &ns.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get namespace by org and name: %w", err)
	}
	return ns, nil
}

// Update updates a namespace.
func (r *NamespaceRepository) Update(ctx context.Context, ns *Namespace) error {
	query := `
		UPDATE cloud_namespaces
		SET name = $2, region = $3, cluster_id = $4, state = $5,
			retention_days = $6, deletion_protected = $7, ha_enabled = $8, standby_region = $9,
			codec_endpoint = $10, codec_pass_token = $11, codec_include_credentials = $12,
			grpc_endpoint = $13, web_endpoint = $14, metrics_endpoint = $15, tags = $16,
			updated_at = $17
		WHERE id = $1
	`
	ns.UpdatedAt = time.Now()
	_, err := r.db.DB().ExecContext(ctx, query,
		ns.ID, ns.Name, ns.Region, ns.ClusterID, ns.State,
		ns.RetentionDays, ns.DeletionProtected, ns.HAEnabled, ns.StandbyRegion,
		ns.CodecEndpoint, ns.CodecPassToken, ns.CodecIncludeCredentials,
		ns.GRPCEndpoint, ns.WebEndpoint, ns.MetricsEndpoint, ns.Tags,
		ns.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update namespace: %w", err)
	}
	return nil
}

// UpdateState updates the namespace state.
func (r *NamespaceRepository) UpdateState(ctx context.Context, id string, state string) error {
	query := `UPDATE cloud_namespaces SET state = $2, updated_at = $3 WHERE id = $1`
	_, err := r.db.DB().ExecContext(ctx, query, id, state, time.Now())
	if err != nil {
		return fmt.Errorf("failed to update namespace state: %w", err)
	}
	return nil
}

// Delete deletes a namespace.
func (r *NamespaceRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM cloud_namespaces WHERE id = $1`
	_, err := r.db.DB().ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete namespace: %w", err)
	}
	return nil
}

// ListByOrganization lists namespaces for an organization.
func (r *NamespaceRepository) ListByOrganization(ctx context.Context, orgID uuid.UUID, limit, offset int) ([]*Namespace, error) {
	query := `
		SELECT id, organization_id, name, region, cluster_id, state,
			retention_days, deletion_protected, ha_enabled, standby_region,
			codec_endpoint, codec_pass_token, codec_include_credentials,
			grpc_endpoint, web_endpoint, metrics_endpoint, tags,
			created_at, updated_at
		FROM cloud_namespaces
		WHERE organization_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.DB().QueryContext(ctx, query, orgID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list namespaces: %w", err)
	}
	defer rows.Close()

	var namespaces []*Namespace
	for rows.Next() {
		ns := &Namespace{}
		if err := rows.Scan(
			&ns.ID, &ns.OrganizationID, &ns.Name, &ns.Region, &ns.ClusterID, &ns.State,
			&ns.RetentionDays, &ns.DeletionProtected, &ns.HAEnabled, &ns.StandbyRegion,
			&ns.CodecEndpoint, &ns.CodecPassToken, &ns.CodecIncludeCredentials,
			&ns.GRPCEndpoint, &ns.WebEndpoint, &ns.MetricsEndpoint, &ns.Tags,
			&ns.CreatedAt, &ns.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan namespace: %w", err)
		}
		namespaces = append(namespaces, ns)
	}
	return namespaces, nil
}

// AddCertificate adds a certificate to a namespace.
func (r *NamespaceRepository) AddCertificate(ctx context.Context, cert *NamespaceCertificate) error {
	query := `
		INSERT INTO namespace_certificates (id, namespace_id, certificate_pem, fingerprint, issuer, subject, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	if cert.ID == uuid.Nil {
		cert.ID = uuid.New()
	}
	cert.CreatedAt = time.Now()

	_, err := r.db.DB().ExecContext(ctx, query,
		cert.ID, cert.NamespaceID, cert.CertificatePEM, cert.Fingerprint,
		cert.Issuer, cert.Subject, cert.ExpiresAt, cert.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to add certificate: %w", err)
	}
	return nil
}

// ListCertificates lists certificates for a namespace.
func (r *NamespaceRepository) ListCertificates(ctx context.Context, namespaceID string) ([]*NamespaceCertificate, error) {
	query := `
		SELECT id, namespace_id, certificate_pem, fingerprint, issuer, subject, expires_at, created_at
		FROM namespace_certificates
		WHERE namespace_id = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.DB().QueryContext(ctx, query, namespaceID)
	if err != nil {
		return nil, fmt.Errorf("failed to list certificates: %w", err)
	}
	defer rows.Close()

	var certs []*NamespaceCertificate
	for rows.Next() {
		cert := &NamespaceCertificate{}
		if err := rows.Scan(
			&cert.ID, &cert.NamespaceID, &cert.CertificatePEM, &cert.Fingerprint,
			&cert.Issuer, &cert.Subject, &cert.ExpiresAt, &cert.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan certificate: %w", err)
		}
		certs = append(certs, cert)
	}
	return certs, nil
}

// AddSearchAttribute adds a search attribute to a namespace.
func (r *NamespaceRepository) AddSearchAttribute(ctx context.Context, attr *NamespaceSearchAttribute) error {
	query := `
		INSERT INTO namespace_search_attributes (namespace_id, name, type, created_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (namespace_id, name) DO UPDATE SET type = $3
	`
	attr.CreatedAt = time.Now()
	_, err := r.db.DB().ExecContext(ctx, query, attr.NamespaceID, attr.Name, attr.Type, attr.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to add search attribute: %w", err)
	}
	return nil
}

// RemoveSearchAttribute removes a search attribute from a namespace.
func (r *NamespaceRepository) RemoveSearchAttribute(ctx context.Context, namespaceID, name string) error {
	query := `DELETE FROM namespace_search_attributes WHERE namespace_id = $1 AND name = $2`
	_, err := r.db.DB().ExecContext(ctx, query, namespaceID, name)
	if err != nil {
		return fmt.Errorf("failed to remove search attribute: %w", err)
	}
	return nil
}

// ListSearchAttributes lists search attributes for a namespace.
func (r *NamespaceRepository) ListSearchAttributes(ctx context.Context, namespaceID string) ([]*NamespaceSearchAttribute, error) {
	query := `
		SELECT namespace_id, name, type, created_at
		FROM namespace_search_attributes
		WHERE namespace_id = $1
		ORDER BY name
	`
	rows, err := r.db.DB().QueryContext(ctx, query, namespaceID)
	if err != nil {
		return nil, fmt.Errorf("failed to list search attributes: %w", err)
	}
	defer rows.Close()

	var attrs []*NamespaceSearchAttribute
	for rows.Next() {
		attr := &NamespaceSearchAttribute{}
		if err := rows.Scan(&attr.NamespaceID, &attr.Name, &attr.Type, &attr.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan search attribute: %w", err)
		}
		attrs = append(attrs, attr)
	}
	return attrs, nil
}
