package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	"go.temporal.io/cloud/internal/repository"
	"go.temporal.io/server/common/log"
)

// NamespaceService handles namespace business logic.
type NamespaceService struct {
	repos  *repository.Repositories
	logger log.Logger
}

// NewNamespaceService creates a new namespace service.
func NewNamespaceService(repos *repository.Repositories, logger log.Logger) *NamespaceService {
	return &NamespaceService{repos: repos, logger: logger}
}

// CreateNamespaceInput is the input for creating a namespace.
type CreateNamespaceInput struct {
	OrganizationID    uuid.UUID
	Name              string
	Region            string
	RetentionDays     int
	DeletionProtected bool
	HAEnabled         bool
	StandbyRegion     string
	Tags              map[string]string
}

// CreateNamespace creates a new namespace.
func (s *NamespaceService) CreateNamespace(ctx context.Context, input *CreateNamespaceInput) (*repository.Namespace, string, error) {
	// Validate name
	if !isValidNamespaceName(input.Name) {
		return nil, "", fmt.Errorf("invalid namespace name: must be 2-63 lowercase alphanumeric characters, hyphens, or underscores")
	}

	// Validate region
	if !isValidRegion(input.Region) {
		return nil, "", fmt.Errorf("invalid region: %s", input.Region)
	}

	// Check subscription limits
	sub, err := s.repos.Subscriptions.GetByOrganizationID(ctx, input.OrganizationID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get subscription: %w", err)
	}
	if sub == nil {
		return nil, "", fmt.Errorf("no subscription found")
	}

	// Check namespace count limit
	namespaces, err := s.repos.Namespaces.ListByOrganization(ctx, input.OrganizationID, 1000, 0)
	if err != nil {
		return nil, "", fmt.Errorf("failed to list namespaces: %w", err)
	}
	maxNamespaces := getMaxNamespaces(sub.Plan)
	if len(namespaces) >= maxNamespaces {
		return nil, "", fmt.Errorf("namespace limit reached for plan %s", sub.Plan)
	}

	// Check if name is taken
	existing, err := s.repos.Namespaces.GetByOrgAndName(ctx, input.OrganizationID, input.Name)
	if err != nil {
		return nil, "", fmt.Errorf("failed to check namespace name: %w", err)
	}
	if existing != nil {
		return nil, "", fmt.Errorf("namespace name already exists")
	}

	// Set defaults
	retentionDays := input.RetentionDays
	if retentionDays == 0 {
		retentionDays = 7
	}
	maxRetention := getMaxRetention(sub.Plan)
	if retentionDays > maxRetention {
		return nil, "", fmt.Errorf("retention period exceeds plan limit of %d days", maxRetention)
	}

	// Generate namespace ID
	nsID := fmt.Sprintf("%s.%s", input.Name, input.OrganizationID.String()[:8])

	// Convert tags to JSON
	tagsJSON, _ := json.Marshal(input.Tags)

	ns := &repository.Namespace{
		ID:                nsID,
		OrganizationID:    input.OrganizationID,
		Name:              input.Name,
		Region:            input.Region,
		State:             "pending",
		RetentionDays:     retentionDays,
		DeletionProtected: input.DeletionProtected,
		HAEnabled:         input.HAEnabled,
		Tags:              tagsJSON,
	}

	if input.StandbyRegion != "" {
		ns.StandbyRegion = sql.NullString{String: input.StandbyRegion, Valid: true}
	}

	if err := s.repos.Namespaces.Create(ctx, ns); err != nil {
		return nil, "", fmt.Errorf("failed to create namespace: %w", err)
	}

	// Generate operation ID for async provisioning
	operationID := uuid.New().String()

	// TODO: Start provisioning workflow
	// This would trigger a Temporal workflow to:
	// 1. Select target cluster
	// 2. Generate mTLS certificates
	// 3. Call Temporal RegisterNamespace API
	// 4. Configure retention, limits, search attrs
	// 5. Create DNS record
	// 6. Update namespace state to 'active'

	return ns, operationID, nil
}

// GetNamespace retrieves a namespace by ID.
func (s *NamespaceService) GetNamespace(ctx context.Context, id string) (*repository.Namespace, error) {
	return s.repos.Namespaces.GetByID(ctx, id)
}

// UpdateNamespaceInput is the input for updating a namespace.
type UpdateNamespaceInput struct {
	ID                string
	RetentionDays     *int
	DeletionProtected *bool
	CodecEndpoint     *string
	CodecPassToken    *bool
	Tags              map[string]string
}

// UpdateNamespace updates a namespace.
func (s *NamespaceService) UpdateNamespace(ctx context.Context, input *UpdateNamespaceInput) (*repository.Namespace, string, error) {
	ns, err := s.repos.Namespaces.GetByID(ctx, input.ID)
	if err != nil {
		return nil, "", err
	}
	if ns == nil {
		return nil, "", fmt.Errorf("namespace not found")
	}

	if input.RetentionDays != nil {
		// Can only increase retention
		if *input.RetentionDays < ns.RetentionDays {
			return nil, "", fmt.Errorf("retention period can only be increased")
		}
		ns.RetentionDays = *input.RetentionDays
	}

	if input.DeletionProtected != nil {
		ns.DeletionProtected = *input.DeletionProtected
	}

	if input.CodecEndpoint != nil {
		ns.CodecEndpoint = sql.NullString{String: *input.CodecEndpoint, Valid: *input.CodecEndpoint != ""}
	}

	if input.CodecPassToken != nil {
		ns.CodecPassToken = *input.CodecPassToken
	}

	if input.Tags != nil {
		tagsJSON, _ := json.Marshal(input.Tags)
		ns.Tags = tagsJSON
	}

	ns.State = "updating"
	if err := s.repos.Namespaces.Update(ctx, ns); err != nil {
		return nil, "", fmt.Errorf("failed to update namespace: %w", err)
	}

	operationID := uuid.New().String()
	// TODO: Start update workflow

	return ns, operationID, nil
}

// DeleteNamespace deletes a namespace.
func (s *NamespaceService) DeleteNamespace(ctx context.Context, id string) (string, error) {
	ns, err := s.repos.Namespaces.GetByID(ctx, id)
	if err != nil {
		return "", err
	}
	if ns == nil {
		return "", fmt.Errorf("namespace not found")
	}

	if ns.DeletionProtected {
		return "", fmt.Errorf("namespace is deletion protected")
	}

	ns.State = "deleting"
	if err := s.repos.Namespaces.Update(ctx, ns); err != nil {
		return "", fmt.Errorf("failed to update namespace state: %w", err)
	}

	operationID := uuid.New().String()
	// TODO: Start deletion workflow

	return operationID, nil
}

// ListNamespaces lists namespaces for an organization.
func (s *NamespaceService) ListNamespaces(ctx context.Context, orgID uuid.UUID, limit, offset int) ([]*repository.Namespace, error) {
	return s.repos.Namespaces.ListByOrganization(ctx, orgID, limit, offset)
}

// AddSearchAttributes adds search attributes to a namespace.
func (s *NamespaceService) AddSearchAttributes(ctx context.Context, namespaceID string, attrs map[string]string) (string, error) {
	for name, attrType := range attrs {
		attr := &repository.NamespaceSearchAttribute{
			NamespaceID: namespaceID,
			Name:        name,
			Type:        attrType,
			CreatedAt:   time.Now(),
		}
		if err := s.repos.Namespaces.AddSearchAttribute(ctx, attr); err != nil {
			return "", fmt.Errorf("failed to add search attribute %s: %w", name, err)
		}
	}

	operationID := uuid.New().String()
	// TODO: Start workflow to add search attributes to Temporal cluster

	return operationID, nil
}

// Helper functions

func isValidNamespaceName(name string) bool {
	if len(name) < 2 || len(name) > 63 {
		return false
	}
	matched, _ := regexp.MatchString(`^[a-z0-9][a-z0-9_-]*[a-z0-9]$`, name)
	return matched
}

func isValidRegion(region string) bool {
	validRegions := map[string]bool{
		"us-east-1": true, "us-east-2": true, "us-west-2": true,
		"eu-west-1": true, "eu-west-2": true, "eu-central-1": true,
		"ap-south-1": true, "ap-southeast-1": true, "ap-northeast-1": true,
	}
	return validRegions[region]
}

func getMaxNamespaces(plan string) int {
	switch plan {
	case "free":
		return 1
	case "essentials":
		return 5
	case "business":
		return 20
	case "enterprise", "mission_critical":
		return 1000
	default:
		return 1
	}
}

func getMaxRetention(plan string) int {
	switch plan {
	case "free":
		return 7
	case "essentials":
		return 30
	case "business":
		return 90
	case "enterprise", "mission_critical":
		return 365
	default:
		return 7
	}
}
