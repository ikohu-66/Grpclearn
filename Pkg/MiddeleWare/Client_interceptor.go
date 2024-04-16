package MiddeleWare

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

func defaultContextTime(ctx context.Context) (context.Context, context.CancelFunc) {
	var cancle context.CancelFunc
	if _, ok := ctx.Deadline(); ok {
		defaultTimeout := 60 * time.Second
		ctx, cancle = context.WithTimeout(ctx, defaultTimeout)
		//通过调用 cancel() 函数，可以手动取消该上下文，释放相关资源并停止与该上下文关联的操作。
	}
	return ctx, cancle
}
func UnaryContextTimeout() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, resp interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, cancel := defaultContextTime(ctx)
		if cancel != nil {
			defer cancel()
		}

		return invoker(ctx, method, req, resp, cc, opts...)
	}
}

func StreamContextTimeout() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx, cancel := defaultContextTime(ctx)
		if cancel != nil {
			defer cancel()
		}

		return streamer(ctx, desc, cc, method, opts...)
	}
}
