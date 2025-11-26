package interceptors

import (
	"context"
	"runtime/debug"

	"connectrpc.com/connect"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
)

// RecoveryInterceptor recovers from panics in handlers.
type RecoveryInterceptor struct {
	logger log.Logger
}

// NewRecoveryInterceptor creates a new recovery interceptor.
func NewRecoveryInterceptor(logger log.Logger) *RecoveryInterceptor {
	return &RecoveryInterceptor{logger: logger}
}

// WrapUnary implements connect.Interceptor.
func (i *RecoveryInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (resp connect.AnyResponse, err error) {
		defer func() {
			if r := recover(); r != nil {
				i.logger.Error("Panic recovered in handler",
					tag.NewAnyTag("panic", r),
					tag.NewStringTag("stack", string(debug.Stack())),
					tag.NewStringTag("procedure", req.Spec().Procedure),
				)
				err = connect.NewError(connect.CodeInternal, nil)
			}
		}()
		return next(ctx, req)
	}
}

// WrapStreamingClient implements connect.Interceptor.
func (i *RecoveryInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

// WrapStreamingHandler implements connect.Interceptor.
func (i *RecoveryInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, conn connect.StreamingHandlerConn) (err error) {
		defer func() {
			if r := recover(); r != nil {
				i.logger.Error("Panic recovered in streaming handler",
					tag.NewAnyTag("panic", r),
					tag.NewStringTag("stack", string(debug.Stack())),
				)
				err = connect.NewError(connect.CodeInternal, nil)
			}
		}()
		return next(ctx, conn)
	}
}
