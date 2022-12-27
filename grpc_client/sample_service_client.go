package grpcclient

import (
	"context"
	"time"

	pb "github.com/abdullahjankhan/go_sample/proto"
)

type SampleServiceClient interface {
	SampleFunc(string) (*pb.SampleResponse, error)
	setConnection(c pb.SampleServiceClient)
}

type sampleServiceClient struct {
	client pb.SampleServiceClient
}

func NewSampleServiceClient() SampleServiceClient {
	return &sampleServiceClient{
		client: nil,
	}
}

func (client *sampleServiceClient) setConnection(c pb.SampleServiceClient) {
	client.client = c
}

func (client *sampleServiceClient) SampleFunc(msg string) (*pb.SampleResponse, error) {
	// setting deadline to timeout if unable to connect with server after 20 secods
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	// defer to cancle the request
	defer cancel()
	// make the actual call to server
	response, err := client.client.SampleFunc(ctx, &pb.SampleRequest{Msg: msg})
	if err != nil {
		return nil, err
	}
	return response, nil
}
