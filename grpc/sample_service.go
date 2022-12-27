package grpc

import (
	"context"
	"fmt"

	"github.com/abdullahjankhan/go_sample/proto"
	"github.com/abdullahjankhan/go_sample/service"
	"github.com/sirupsen/logrus"
)

type SampleGrpcServer struct {
	proto.UnimplementedSampleServiceServer
	logger *logrus.Logger
}

func NewSampleGrpcServer(container *service.Container) *SampleGrpcServer {
	return &SampleGrpcServer{
		logger: container.LoggerService.GetInstance(),
	}
}

func (grpc *SampleGrpcServer) SampleFunc(ctx context.Context, req *proto.SampleRequest) (*proto.SampleResponse, error) {
	grpc.logger.Infof("Request Receive at SampleFunc: %v", req)

	return &proto.SampleResponse{
		MsgRes: fmt.Sprintf("Respose to: %v", req),
	}, nil
}
