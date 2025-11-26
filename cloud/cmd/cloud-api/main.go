// Package main is the entry point for the Cloud API service.
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/rs/cors"
	"go.temporal.io/cloud/internal/api/v1"
	"go.temporal.io/cloud/internal/config"
	"go.temporal.io/cloud/internal/interceptors"
	"go.temporal.io/cloud/internal/repository"
	"go.temporal.io/cloud/internal/service"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	logger := log.NewZapLogger(log.BuildZapLogger(log.Config{
		Level:  "info",
		Format: "json",
	}))

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load config", tag.Error(err))
	}

	// Initialize database connection
	db, err := repository.NewPostgresDB(cfg.Database)
	if err != nil {
		logger.Fatal("Failed to connect to database", tag.Error(err))
	}
	defer db.Close()

	// Initialize repositories
	repos := repository.NewRepositories(db)

	// Initialize services
	orgService := service.NewOrganizationService(repos, logger)
	nsService := service.NewNamespaceService(repos, logger)
	billingService := service.NewBillingService(repos, cfg.Stripe, logger)
	identityService := service.NewIdentityService(repos, cfg.JWT, logger)
	auditService := service.NewAuditService(repos, logger)

	// Create interceptors
	authInterceptor := interceptors.NewAuthInterceptor(identityService, logger)
	auditInterceptor := interceptors.NewAuditInterceptor(auditService, logger)
	rateLimitInterceptor := interceptors.NewRateLimitInterceptor(cfg.RateLimit, logger)
	recoveryInterceptor := interceptors.NewRecoveryInterceptor(logger)

	interceptorChain := connect.WithInterceptors(
		recoveryInterceptor,
		rateLimitInterceptor,
		authInterceptor,
		auditInterceptor,
	)

	// Create HTTP mux
	mux := http.NewServeMux()

	// Register services
	orgHandler := api.NewOrganizationHandler(orgService)
	mux.Handle(orgHandler.Path(), orgHandler.Handler(interceptorChain))

	nsHandler := api.NewNamespaceHandler(nsService)
	mux.Handle(nsHandler.Path(), nsHandler.Handler(interceptorChain))

	billingHandler := api.NewBillingHandler(billingService)
	mux.Handle(billingHandler.Path(), billingHandler.Handler(interceptorChain))

	identityHandler := api.NewIdentityHandler(identityService)
	mux.Handle(identityHandler.Path(), identityHandler.Handler(interceptorChain))

	auditHandler := api.NewAuditHandler(auditService)
	mux.Handle(auditHandler.Path(), auditHandler.Handler(interceptorChain))

	// Register reflection for grpcurl/grpcui
	reflector := grpcreflect.NewStaticReflector(
		"temporal.cloud.api.v1.OrganizationService",
		"temporal.cloud.api.v1.NamespaceService",
		"temporal.cloud.api.v1.BillingService",
		"temporal.cloud.api.v1.IdentityService",
		"temporal.cloud.api.v1.AuditService",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	// Health check endpoints
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Database unavailable"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ready"))
	})

	// CORS configuration
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   cfg.CORS.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Connect-Protocol-Version"},
		ExposedHeaders:   []string{"Grpc-Status", "Grpc-Message", "Grpc-Status-Details-Bin"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Create server with HTTP/2 support
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      corsHandler.Handler(h2c.NewHandler(mux, &http2.Server{})),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server in goroutine
	go func() {
		logger.Info("Starting Cloud API server", tag.NewInt("port", cfg.Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server failed", tag.Error(err))
		}
	}()

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", tag.Error(err))
	}

	logger.Info("Server stopped")
}
