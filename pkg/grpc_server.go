package pkg

import "google.golang.org/grpc"

type GrpcServer struct {
	Server *grpc.Server
}

func NewGrpcServer() GrpcServer {
	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*4),
		grpc.MaxSendMsgSize(1024*1024*4),
	)
	return GrpcServer{Server: grpcServer}
}
