package MiddeleWare

import (
	"Grpclearn/Pkg/Errocode"
	"Grpclearn/Pkg/Log"
	"context"
	"google.golang.org/grpc"
	"runtime/debug"
	"time"
)

func Acesslog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	requeseLog := "access request log: method: %s, begin_time: %d, request: %v"
	begintime := time.Now().Local().Unix()
	Log.RpcLog.Logger.Info(requeseLog, begintime, info.FullMethod, req)
	resp, err := handler(ctx, req)
	responseLog := "access response log: method: %s, begin_time: %d, end_time: %d, response: %v"
	endTime := time.Now().Local().Unix()
	Log.RpcLog.Logger.Info(responseLog, endTime, resp)
	return resp, err
}
func ErrorLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		errLog := "error log: method: %s, code: %v, message: %v, details: %v"
		s := Errocode.FromError(err)
		Log.RpcLog.Logger.Error(errLog, info.FullMethod, s.Code(), s.Err().Error(), s.Details())
	}
	return resp, err
}
func Recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {

		if e := recover(); e != nil {
			recoveryLog := "recovery log: method: %s, message: %v, stack: %s"
			Log.RpcLog.Logger.Panic(recoveryLog, info.FullMethod, e, string(debug.Stack()[:]))
		}
	}()

	return handler(ctx, req)
}
