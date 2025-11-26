// Package service provides business logic for cloud services.
package service

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"go.temporal.io/cloud/internal/repository"
	"go.temporal.io/server/common/log"
)

// OrganizationService handles organization business logic.
type OrganizationService struct {
	repos  *repository.Repositories
	logger log.Logger
}

// NewOrganizationService creates a new organization service.
func NewOrganizationService(repos *repository.Repositories, logger log.Logger) *OrganizationService {
	return &OrganizationService{repos: repos, logger: logger}
}

// CreateOrganizationInput is the input for creating an organization.
type CreateOrganizationInput struct {
	Name    string
	Slug    string
	OwnerID uuid.UUID
}

// CreateOrganization creates a new organization.
func (s *OrganizationService) CreateOrganization(ctx context.Context, input *CreateOrganizationInput) (*repository.Organization, error) {
	// Validate name
	if len(input.Name) < 2 || len(input.Name) > 255 {
		return nil, fmt.Errorf("organization name must be between 2 and 255 characters")
	}

	// Generate slug if not provided
	slug := input.Slug
	if slug == "" {
		slug = generateSlug(input.Name)
	}

	// Validate slug
	if !isValidSlug(slug) {
		return nil, fmt.Errorf("invalid slug: must be 2-63 lowercase alphanumeric characters or hyphens")
	}

	// Check if slug is taken
	existing, err := s.repos.Organizations.GetBySlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("failed to check slug availability: %w", err)
	}
	if existing != nil {
		return nil, fmt.Errorf("organization slug already exists")
	}

	// Create organization
	org := &repository.Organization{
		Name:     input.Name,
		Slug:     slug,
		Settings: json.RawMessage("{}"),
	}

	if err := s.repos.Organizations.Create(ctx, org); err != nil {
		return nil, fmt.Errorf("failed to create organization: %w", err)
	}

	// Add owner as member
	member := &repository.OrganizationMember{
		OrganizationID: org.ID,
		UserID:         input.OwnerID,
		Role:           "owner",
	}
	if err := s.repos.Organizations.AddMember(ctx, member); err != nil {
		return nil, fmt.Errorf("failed to add owner: %w", err)
	}

	// Create default subscription (free tier)
	actions, activeGB, retainedGB := repository.GetPlanLimits("free")
	sub := &repository.Subscription{
		OrganizationID:    org.ID,
		Plan:              "free",
		Status:            "active",
		ActionsIncluded:   actions,
		ActiveStorageGB:   activeGB,
		RetainedStorageGB: retainedGB,
	}
	if err := s.repos.Subscriptions.Create(ctx, sub); err != nil {
		return nil, fmt.Errorf("failed to create subscription: %w", err)
	}

	return org, nil
}

// GetOrganization retrieves an organization by ID.
func (s *OrganizationService) GetOrganization(ctx context.Context, id uuid.UUID) (*repository.Organization, error) {
	return s.repos.Organizations.GetByID(ctx, id)
}

// GetOrganizationBySlug retrieves an organization by slug.
func (s *OrganizationService) GetOrganizationBySlug(ctx context.Context, slug string) (*repository.Organization, error) {
	return s.repos.Organizations.GetBySlug(ctx, slug)
}

// UpdateOrganizationInput is the input for updating an organization.
type UpdateOrganizationInput struct {
	ID       uuid.UUID
	Name     *string
	Settings json.RawMessage
}

// UpdateOrganization updates an organization.
func (s *OrganizationService) UpdateOrganization(ctx context.Context, input *UpdateOrganizationInput) (*repository.Organization, error) {
	org, err := s.repos.Organizations.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	if org == nil {
		return nil, fmt.Errorf("organization not found")
	}

	if input.Name != nil {
		org.Name = *input.Name
	}
	if input.Settings != nil {
		org.Settings = input.Settings
	}

	if err := s.repos.Organizations.Update(ctx, org); err != nil {
		return nil, fmt.Errorf("failed to update organization: %w", err)
	}

	return org, nil
}

// DeleteOrganization deletes an organization.
func (s *OrganizationService) DeleteOrganization(ctx context.Context, id uuid.UUID) error {
	// Check for active namespaces
	namespaces, err := s.repos.Namespaces.ListByOrganization(ctx, id, 1, 0)
	if err != nil {
		return fmt.Errorf("failed to check namespaces: %w", err)
	}
	if len(namespaces) > 0 {
		return fmt.Errorf("cannot delete organization with active namespaces")
	}

	return s.repos.Organizations.Delete(ctx, id)
}

// ListOrganizations lists organizations for a user.
func (s *OrganizationService) ListOrganizations(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*repository.Organization, error) {
	return s.repos.Organizations.ListByUserID(ctx, userID, limit, offset)
}

// InviteUserInput is the input for inviting a user.
type InviteUserInput struct {
	OrganizationID uuid.UUID
	Email          string
	Role           string
	InvitedBy      uuid.UUID
}

// InviteUser invites a user to an organization.
func (s *OrganizationService) InviteUser(ctx context.Context, input *InviteUserInput) (*repository.UserInvitation, error) {
	// Validate role
	validRoles := map[string]bool{"owner": true, "admin": true, "developer": true, "read_only": true}
	if !validRoles[input.Role] {
		return nil, fmt.Errorf("invalid role: %s", input.Role)
	}

	// Generate invitation token
	token := uuid.New().String()

	inv := &repository.UserInvitation{
		OrganizationID: input.OrganizationID,
		Email:          input.Email,
		Role:           input.Role,
		InvitedBy:      input.InvitedBy,
		Token:          token,
	}

	if err := s.repos.Users.CreateInvitation(ctx, inv); err != nil {
		return nil, fmt.Errorf("failed to create invitation: %w", err)
	}

	return inv, nil
}

// ListMembers lists members of an organization.
func (s *OrganizationService) ListMembers(ctx context.Context, orgID uuid.UUID, limit, offset int) ([]*repository.OrganizationMember, error) {
	return s.repos.Organizations.ListMembers(ctx, orgID, limit, offset)
}

// UpdateMemberRole updates a member's role.
func (s *OrganizationService) UpdateMemberRole(ctx context.Context, orgID, userID uuid.UUID, role string) error {
	member := &repository.OrganizationMember{
		OrganizationID: orgID,
		UserID:         userID,
		Role:           role,
	}
	return s.repos.Organizations.AddMember(ctx, member)
}

// RemoveMember removes a member from an organization.
func (s *OrganizationService) RemoveMember(ctx context.Context, orgID, userID uuid.UUID) error {
	return s.repos.Organizations.RemoveMember(ctx, orgID, userID)
}

// Helper functions

func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	if len(slug) > 63 {
		slug = slug[:63]
	}
	return slug
}

func isValidSlug(slug string) bool {
	if len(slug) < 2 || len(slug) > 63 {
		return false
	}
	matched, _ := regexp.MatchString(`^[a-z0-9][a-z0-9-]*[a-z0-9]$`, slug)
	return matched
}
