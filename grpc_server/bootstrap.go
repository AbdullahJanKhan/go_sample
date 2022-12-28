package grpc

import (
	"context"

	"github.com/abdullahjankhan/go_sample/proto"
	"github.com/abdullahjankhan/go_sample/service"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

// NewServer provides new grpc server according to configuration provided
func NewServer(container *service.Container, unaryInterceptors []grpc.UnaryServerInterceptor, streamInterceptors []grpc.StreamServerInterceptor) *grpc.Server {
	opts := []grpcrecovery.Option{
		grpcrecovery.WithRecoveryHandlerContext(func(ctx context.Context, rec interface{}) (err error) {
			container.LoggerService.GetInstance().Fatalf("[gRPC|Server] Recovered in %v", rec)

			return status.Errorf(codes.Internal, "Recovered in %v", rec)
		}),
	}

	server := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             container.GlobalConfigService.GetConfig().Grpc.ServerMinTime, // If a client pings more than once every 5 minutes, terminate the connection
			PermitWithoutStream: true,                                                         // Allow pings even when there are no active streams
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    container.GlobalConfigService.GetConfig().Grpc.ServerMinTime, // Ping the client if it is idle for 2 hours to ensure the connection is still active
			Timeout: container.GlobalConfigService.GetConfig().Grpc.ServerTimeOut, // Wait 20 second for the ping ack before assuming the connection is dead
		}),
		grpcmiddleware.WithUnaryServerChain(
			append([]grpc.UnaryServerInterceptor{
				grpcrecovery.UnaryServerInterceptor(opts...),
			}, unaryInterceptors...)...,
		),
		grpcmiddleware.WithStreamServerChain(
			append([]grpc.StreamServerInterceptor{
				grpcrecovery.StreamServerInterceptor(opts...),
			}, streamInterceptors...)...,
		),
	)

	return server
}

// StartServer create grpc instance and bind servies with grpc
func StartServer(container *service.Container) {

	grpcServer := NewServer(
		container,
		nil,
		nil,
	)

	sampleServiceServer := NewSampleGrpcServer(container)

	proto.RegisterSampleServiceServer(grpcServer, sampleServiceServer)

	go Start(container, grpcServer)

	container.LoggerService.GetInstance().Info("rpc server ok")
}
