package Errocode

import (
	pb "Grpclearn/Proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Status struct {
	*status.Status
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}

func TogRPCError(err *Error) error {
	s, _ := status.New(ToRPCCode(err.Code()), err.Msg()).WithDetails(&pb.Error{Code: string(err.Code()), Message: err.Msg()})
	return s.Err()
}

// 内部调用以查询错误状况
func ToRPCStatus(code string, msg string) *Status {
	s, _ := status.New(ToRPCCode(code), msg).WithDetails(&pb.Error{Code: code, Message: msg})
	return &Status{s}
}

func ToRPCCode(code string) codes.Code {
	var statucode codes.Code
	switch code {
	case Fail.Code():
		statucode = codes.Internal
	case InvalidParams.code:
		statucode = codes.InvalidArgument
	case Unauthorized.code:
		statucode = codes.Unavailable
	default:
		statucode = codes.Unknown
	}
	return statucode
}
