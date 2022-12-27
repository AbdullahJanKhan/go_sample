package grpc

import (
	"context"
	"fmt"

	"github.com/abdullahjankhan/go_sample/proto"
	"github.com/sirupsen/logrus"
)

type SampleGrpcServer struct {
	proto.UnimplementedSampleServiceServer
	logger *logrus.Logger
}

func NewSampleGrpcServer() *SampleGrpcServer {
	return &SampleGrpcServer{
		// logger: container.LoggerService.GetInstance(),
	}
}

func (grpc *SampleGrpcServer) SampleFunc(ctx context.Context, req *proto.SampleRequest) (*proto.SampleResponse, error) {
	return &proto.SampleResponse{
		MsgRes: fmt.Sprintf("Respose to: %v", req),
	}, nil
}
