// Package workflows provides Temporal workflows for cloud operations.
package workflows

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// ProvisionNamespaceInput is the input for the provision namespace workflow.
type ProvisionNamespaceInput struct {
	NamespaceID    string
	OrganizationID string
	Name           string
	Region         string
	RetentionDays  int
	HAEnabled      bool
	StandbyRegion  string
}

// ProvisionNamespaceOutput is the output of the provision namespace workflow.
type ProvisionNamespaceOutput struct {
	NamespaceID     string
	GRPCEndpoint    string
	WebEndpoint     string
	MetricsEndpoint string
	ClusterID       string
}

// ProvisionNamespaceWorkflow provisions a new namespace.
func ProvisionNamespaceWorkflow(ctx workflow.Context, input ProvisionNamespaceInput) (*ProvisionNamespaceOutput, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting namespace provisioning", "namespace_id", input.NamespaceID)

	// Activity options with retries
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    5,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Step 1: Select target cluster
	var clusterID string
	var a *Activities
	err := workflow.ExecuteActivity(ctx, a.SelectClusterActivity, SelectClusterInput{
		Region:    input.Region,
		HAEnabled: input.HAEnabled,
	}).Get(ctx, &clusterID)
	if err != nil {
		return nil, fmt.Errorf("failed to select cluster: %w", err)
	}

	// Step 2: Generate mTLS certificates
	var certOutput GenerateCertificatesOutput
	err = workflow.ExecuteActivity(ctx, a.GenerateCertificatesActivity, GenerateCertificatesInput{
		NamespaceID:    input.NamespaceID,
		OrganizationID: input.OrganizationID,
	}).Get(ctx, &certOutput)
	if err != nil {
		return nil, fmt.Errorf("failed to generate certificates: %w", err)
	}

	// Step 3: Register namespace in Temporal cluster
	var registerOutput RegisterNamespaceOutput
	err = workflow.ExecuteActivity(ctx, a.RegisterNamespaceActivity, RegisterNamespaceInput{
		ClusterID:     clusterID,
		NamespaceID:   input.NamespaceID,
		Name:          input.Name,
		RetentionDays: input.RetentionDays,
	}).Get(ctx, &registerOutput)
	if err != nil {
		return nil, fmt.Errorf("failed to register namespace: %w", err)
	}

	// Step 4: Create DNS record
	var dnsOutput CreateDNSRecordOutput
	err = workflow.ExecuteActivity(ctx, a.CreateDNSRecordActivity, CreateDNSRecordInput{
		NamespaceID: input.NamespaceID,
		Region:      input.Region,
		ClusterID:   clusterID,
	}).Get(ctx, &dnsOutput)
	if err != nil {
		return nil, fmt.Errorf("failed to create DNS record: %w", err)
	}

	// Step 5: Update namespace state to active
	err = workflow.ExecuteActivity(ctx, a.UpdateNamespaceStateActivity, UpdateNamespaceStateInput{
		NamespaceID:     input.NamespaceID,
		State:           "active",
		ClusterID:       clusterID,
		GRPCEndpoint:    dnsOutput.GRPCEndpoint,
		WebEndpoint:     dnsOutput.WebEndpoint,
		MetricsEndpoint: dnsOutput.MetricsEndpoint,
	}).Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update namespace state: %w", err)
	}

	// Step 6: If HA enabled, setup standby region
	if input.HAEnabled && input.StandbyRegion != "" {
		err = workflow.ExecuteActivity(ctx, a.SetupStandbyRegionActivity, SetupStandbyInput{
			NamespaceID:   input.NamespaceID,
			PrimaryRegion: input.Region,
			StandbyRegion: input.StandbyRegion,
		}).Get(ctx, nil)
		if err != nil {
			logger.Warn("Failed to setup standby region", "error", err)
			// Non-fatal, continue
		}
	}

	logger.Info("Namespace provisioning completed", "namespace_id", input.NamespaceID)

	return &ProvisionNamespaceOutput{
		NamespaceID:     input.NamespaceID,
		GRPCEndpoint:    dnsOutput.GRPCEndpoint,
		WebEndpoint:     dnsOutput.WebEndpoint,
		MetricsEndpoint: dnsOutput.MetricsEndpoint,
		ClusterID:       clusterID,
	}, nil
}

// DeleteNamespaceInput is the input for the delete namespace workflow.
type DeleteNamespaceInput struct {
	NamespaceID    string
	OrganizationID string
}

// DeleteNamespaceWorkflow deletes a namespace.
func DeleteNamespaceWorkflow(ctx workflow.Context, input DeleteNamespaceInput) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting namespace deletion", "namespace_id", input.NamespaceID)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    3,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Step 1: Deprecate namespace in Temporal cluster
	var a *Activities
	err := workflow.ExecuteActivity(ctx, a.DeprecateNamespaceActivity, input.NamespaceID).Get(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to deprecate namespace: %w", err)
	}

	// Step 2: Wait for workflows to drain (optional, with timeout)
	_ = workflow.Sleep(ctx, 5*time.Minute)

	// Step 3: Remove DNS record
	err = workflow.ExecuteActivity(ctx, a.RemoveDNSRecordActivity, input.NamespaceID).Get(ctx, nil)
	if err != nil {
		logger.Warn("Failed to remove DNS record", "error", err)
	}

	// Step 4: Archive namespace data to S3
	err = workflow.ExecuteActivity(ctx, a.ArchiveNamespaceActivity, input.NamespaceID).Get(ctx, nil)
	if err != nil {
		logger.Warn("Failed to archive namespace", "error", err)
	}

	// Step 5: Update namespace state to deleted
	err = workflow.ExecuteActivity(ctx, a.UpdateNamespaceStateActivity, UpdateNamespaceStateInput{
		NamespaceID: input.NamespaceID,
		State:       "deleted",
	}).Get(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to update namespace state: %w", err)
	}

	logger.Info("Namespace deletion completed", "namespace_id", input.NamespaceID)
	return nil
}

// FailoverNamespaceInput is the input for the failover workflow.
type FailoverNamespaceInput struct {
	NamespaceID  string
	TargetRegion string
}

// FailoverNamespaceWorkflow performs a namespace failover.
func FailoverNamespaceWorkflow(ctx workflow.Context, input FailoverNamespaceInput) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting namespace failover", "namespace_id", input.NamespaceID, "target", input.TargetRegion)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    30 * time.Second,
			MaximumAttempts:    3,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Step 1: Verify standby is ready
	var a *Activities
	err := workflow.ExecuteActivity(ctx, a.VerifyStandbyReadyActivity, input).Get(ctx, nil)
	if err != nil {
		return fmt.Errorf("standby not ready: %w", err)
	}

	// Step 2: Fence primary (stop accepting writes)
	err = workflow.ExecuteActivity(ctx, a.FencePrimaryActivity, input.NamespaceID).Get(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to fence primary: %w", err)
	}

	// Step 3: Promote standby
	err = workflow.ExecuteActivity(ctx, a.PromoteStandbyActivity, input).Get(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to promote standby: %w", err)
	}

	// Step 4: Update DNS to point to new primary
	err = workflow.ExecuteActivity(ctx, a.UpdateDNSForFailoverActivity, input).Get(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to update DNS: %w", err)
	}

	// Step 5: Verify traffic switched
	err = workflow.ExecuteActivity(ctx, a.VerifyTrafficSwitchedActivity, input).Get(ctx, nil)
	if err != nil {
		logger.Warn("Traffic verification failed", "error", err)
	}

	logger.Info("Namespace failover completed", "namespace_id", input.NamespaceID)
	return nil
}

// Activity input/output types

type SelectClusterInput struct {
	Region    string
	HAEnabled bool
}

type GenerateCertificatesInput struct {
	NamespaceID    string
	OrganizationID string
}

type GenerateCertificatesOutput struct {
	CertificatePEM string
	Fingerprint    string
}

type RegisterNamespaceInput struct {
	ClusterID     string
	NamespaceID   string
	Name          string
	RetentionDays int
}

type RegisterNamespaceOutput struct {
	Success bool
}

type CreateDNSRecordInput struct {
	NamespaceID string
	Region      string
	ClusterID   string
}

type CreateDNSRecordOutput struct {
	GRPCEndpoint    string
	WebEndpoint     string
	MetricsEndpoint string
}

type UpdateNamespaceStateInput struct {
	NamespaceID     string
	State           string
	ClusterID       string
	GRPCEndpoint    string
	WebEndpoint     string
	MetricsEndpoint string
}

type SetupStandbyInput struct {
	NamespaceID   string
	PrimaryRegion string
	StandbyRegion string
}
