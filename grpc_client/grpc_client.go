package grpcclient

import "github.com/abdullahjankhan/go_sample/service"

type GrpcClient interface {
	SampleFunc(string) (string, error)
}

type grpcClient struct {
	sampleClient SampleServiceClient
	container    service.Container
}

func NewGrpcClient() GrpcClient {
	sampleClient := NewSampleServiceClient()

	grpcClient := &grpcClient{
		sampleClient: sampleClient,
	}
	go grpcClient.initializeConnections()
	return grpcClient
}

func (c *grpcClient) SampleFunc(msg string) (string, error) {
	return "", nil
}
