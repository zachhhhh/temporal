package interceptors

import (
	"context"
	"sync"
	"time"

	"connectrpc.com/connect"
	"go.temporal.io/cloud/internal/config"
	"go.temporal.io/server/common/log"
	"golang.org/x/time/rate"
)

// RateLimitInterceptor handles rate limiting for gRPC requests.
type RateLimitInterceptor struct {
	config   config.RateLimitConfig
	logger   log.Logger
	limiters sync.Map // map[string]*rate.Limiter
}

// NewRateLimitInterceptor creates a new rate limit interceptor.
func NewRateLimitInterceptor(cfg config.RateLimitConfig, logger log.Logger) *RateLimitInterceptor {
	return &RateLimitInterceptor{
		config: cfg,
		logger: logger,
	}
}

// WrapUnary implements connect.Interceptor.
func (i *RateLimitInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		if !i.config.Enabled {
			return next(ctx, req)
		}

		// Get client identifier (from auth or IP)
		clientID := i.getClientID(ctx, req)

		// Get or create limiter for this client
		limiter := i.getLimiter(clientID)

		// Check rate limit
		if !limiter.Allow() {
			return nil, connect.NewError(
				connect.CodeResourceExhausted,
				nil,
			)
		}

		return next(ctx, req)
	}
}

// WrapStreamingClient implements connect.Interceptor.
func (i *RateLimitInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

// WrapStreamingHandler implements connect.Interceptor.
func (i *RateLimitInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
}

func (i *RateLimitInterceptor) getClientID(ctx context.Context, req connect.AnyRequest) string {
	// Try to get from auth context
	if authInfo := GetAuthInfo(ctx); authInfo != nil {
		if authInfo.OrganizationID.String() != "" {
			return "org:" + authInfo.OrganizationID.String()
		}
		return "user:" + authInfo.UserID.String()
	}

	// Fall back to IP address
	if ip := req.Header().Get("X-Forwarded-For"); ip != "" {
		return "ip:" + ip
	}
	if ip := req.Header().Get("X-Real-IP"); ip != "" {
		return "ip:" + ip
	}

	return "unknown"
}

func (i *RateLimitInterceptor) getLimiter(clientID string) *rate.Limiter {
	if limiter, ok := i.limiters.Load(clientID); ok {
		return limiter.(*rate.Limiter)
	}

	// Create new limiter
	limiter := rate.NewLimiter(
		rate.Limit(i.config.RequestsPerSec),
		i.config.BurstSize,
	)
	i.limiters.Store(clientID, limiter)

	// Schedule cleanup after 1 hour of inactivity
	go func() {
		time.Sleep(1 * time.Hour)
		i.limiters.Delete(clientID)
	}()

	return limiter
}
