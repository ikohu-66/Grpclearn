package main

import (
	"Grpclearn/Pkg/MiddeleWare"
	pb "Grpclearn/Proto"
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
)

func main() {
	//在新增 metadata 信息时，务必使用 Append 类别的方法，否则如果直接 New 一个全新的 md，将会导致原有的 metadata 信息丢失
	athu := Atus{
		Appkey:    "grpc_getList",
		AppSecret: "yyds",
	}
	ctx := context.Background()
	clientconn, _ := GetClinetConn(ctx, "localhost:8000", []grpc.DialOption{
		grpc.WithUnaryInterceptor(
			MiddeleWare.UnaryContextTimeout(),
		),
		grpc.WithPerRPCCredentials(&athu),
	})
	defer clientconn.Close()
	tagservice := pb.NewTagServiceClient(clientconn)
	resp, _ := tagservice.GetTageList(ctx, &pb.GetTaglistReques{Name: "weatherInfo"})
	log.Printf("resp: %v", resp)
}
func GetClinetConn(ctx context.Context, targe string, options []grpc.DialOption) (*grpc.ClientConn, error) {
	options = append(options, grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
		grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithMax(2),
			grpc_retry.WithCodes(
				codes.Unknown,
				codes.Internal,
				codes.DeadlineExceeded,
			),
		),
	)))
	//grpc.DialContext 方法是异步建立连接的，并不会马上就成为可用连接了，仅处于 Connecting 状态（需要多久则取决于外部因素，例如：网络），正式要到达 Ready 状态，这个连接才算是真正的可用

	return grpc.NewClient(targe, options...)
}
