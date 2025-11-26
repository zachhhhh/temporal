// Package api provides gRPC API handlers for the Cloud API.
package api

import (
	"net/http"

	"connectrpc.com/connect"
	"go.temporal.io/cloud/internal/service"
)

// OrganizationHandler handles organization API requests.
type OrganizationHandler struct {
	service *service.OrganizationService
}

// NewOrganizationHandler creates a new organization handler.
func NewOrganizationHandler(svc *service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{service: svc}
}

// Path returns the base path for the handler.
func (h *OrganizationHandler) Path() string {
	return "/temporal.cloud.api.v1.OrganizationService/"
}

// Handler returns the HTTP handler with interceptors.
func (h *OrganizationHandler) Handler(opts ...connect.HandlerOption) http.Handler {
	mux := http.NewServeMux()
	// TODO: Register Connect handlers when proto is generated
	return mux
}

// NamespaceHandler handles namespace API requests.
type NamespaceHandler struct {
	service *service.NamespaceService
}

// NewNamespaceHandler creates a new namespace handler.
func NewNamespaceHandler(svc *service.NamespaceService) *NamespaceHandler {
	return &NamespaceHandler{service: svc}
}

// Path returns the base path for the handler.
func (h *NamespaceHandler) Path() string {
	return "/temporal.cloud.api.v1.NamespaceService/"
}

// Handler returns the HTTP handler with interceptors.
func (h *NamespaceHandler) Handler(opts ...connect.HandlerOption) http.Handler {
	mux := http.NewServeMux()
	return mux
}

// BillingHandler handles billing API requests.
type BillingHandler struct {
	service *service.BillingService
}

// NewBillingHandler creates a new billing handler.
func NewBillingHandler(svc *service.BillingService) *BillingHandler {
	return &BillingHandler{service: svc}
}

// Path returns the base path for the handler.
func (h *BillingHandler) Path() string {
	return "/temporal.cloud.api.v1.BillingService/"
}

// Handler returns the HTTP handler with interceptors.
func (h *BillingHandler) Handler(opts ...connect.HandlerOption) http.Handler {
	mux := http.NewServeMux()
	return mux
}

// IdentityHandler handles identity API requests.
type IdentityHandler struct {
	service *service.IdentityService
}

// NewIdentityHandler creates a new identity handler.
func NewIdentityHandler(svc *service.IdentityService) *IdentityHandler {
	return &IdentityHandler{service: svc}
}

// Path returns the base path for the handler.
func (h *IdentityHandler) Path() string {
	return "/temporal.cloud.api.v1.IdentityService/"
}

// Handler returns the HTTP handler with interceptors.
func (h *IdentityHandler) Handler(opts ...connect.HandlerOption) http.Handler {
	mux := http.NewServeMux()
	return mux
}

// AuditHandler handles audit API requests.
type AuditHandler struct {
	service *service.AuditService
}

// NewAuditHandler creates a new audit handler.
func NewAuditHandler(svc *service.AuditService) *AuditHandler {
	return &AuditHandler{service: svc}
}

// Path returns the base path for the handler.
func (h *AuditHandler) Path() string {
	return "/temporal.cloud.api.v1.AuditService/"
}

// Handler returns the HTTP handler with interceptors.
func (h *AuditHandler) Handler(opts ...connect.HandlerOption) http.Handler {
	mux := http.NewServeMux()
	return mux
}
