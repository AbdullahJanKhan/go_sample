package grpc

import (
	"fmt"
	"net"

	"github.com/abdullahjankhan/go_sample/service"
	"google.golang.org/grpc"
)

// Start bind grpc server with specifc address
func Start(container *service.Container, grpcServer *grpc.Server) {
	address := fmt.Sprintf("%v:%v", "127.0.0.1" /*container.GbeConfigService.GetConfig().Grpc.Addr*/, "8000" /*container.GbeConfigService.GetConfig().Grpc.Port*/)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		container.LoggerService.GetInstance().Fatalf("failed to listen: %v", err)
	}

	//server rpc server over specific port
	if err := grpcServer.Serve(lis); err != nil {
		container.LoggerService.GetInstance().Fatalf("failed to serve: %v", err)
	}
}
