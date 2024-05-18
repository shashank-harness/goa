package middleware

import (
	"context"
	"sync"
	"sync/atomic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StreamCanceler provides a middleware that can be used to gracefully stop
// streaming requests.  To stop streaming requests, simply pass in a context
// with cancellation and cancel the context.  When the context given to the
// StreamCanceler is canceled, it does the following:
//  1. Stops accepting further streaming requests and returns the code
//     Unavailable with message "server is stopping".
//  2. Cancels the context of all streaming requests. Your request handler
//     should obey to the cancelation of request context.
//
// Example:
//
//	var (
//	    ctxCancel  context.Context
//	    cancelFunc context.CancelFunc
//	)
//	ctxCancel, cancelFunc = context.WithCancel(parentCtx)
//	streamInterceptor := StreamCanceler(ctxCancel)
//	// Use the interceptor in your server and when you need to shutdown
//	// your server, simply cancel the context given to the StreamCanceler interceptor.
//	cancelFunc()
//
//	// In your application code, look for context cancellation and respond with proper code.
//	for {
//	    select {
//	    case <-ctx.Done():
//	        return status.Error(codes.Canceled, "canceled")
//	...
func StreamCanceler(ctx context.Context) grpc.StreamServerInterceptor {
	var (
		cancels   sync.Map
		canceling uint32
	)

	go func() {
		<-ctx.Done()
		atomic.StoreUint32(&canceling, 1)
		cancels.Range(func(key any, value any) bool {
			cancel := key.(*context.CancelFunc)
			(*cancel)()
			return true
		})
	}()
	return grpc.StreamServerInterceptor(func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if atomic.LoadUint32(&canceling) == 1 {
			return status.Error(codes.Unavailable, "server is stopping")
		}
		cctx, cancel := context.WithCancel(ss.Context())
		cancels.Store(&cancel, struct{}{})
		err := handler(srv, NewWrappedServerStream(cctx, ss))
		cancels.Delete(&cancel)
		cancel()

		return err
	})
}
