package grpcclient

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	pb "github.com/abdullahjankhan/go_sample/proto"
)

var kacp = keepalive.ClientParameters{
	Time:                15 * time.Second,
	Timeout:             20 * time.Second,
	PermitWithoutStream: true,
}

// init all connection at the start of server
func (c *grpcClient) initializeConnections() {
	sampleConn := make(chan bool)
	go c.sampleConnect(sampleConn)
	if val := <-sampleConn; val {
		log.Println("Misbar service connected successfully")
	}
}

// sample client servise connect to sample service and return true/false on connection status
func (c *grpcClient) sampleConnect(connected chan bool) {
	//Set up a connection to the server.
	address := fmt.Sprintf("%v:%v", "127.0.0.1" /*container.GbeConfigService.GetConfig().Grpc.Addr*/, "8000" /*container.GbeConfigService.GetConfig().Grpc.Port*/)
	sampleConn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))

	if err != nil {
		log.Fatalf("did not connect misbar(kyc) service over GRPC: %v", err)
		connected <- false
	}

	client := pb.NewSampleServiceClient(sampleConn)
	c.sampleClient.setConnection(client)
	connected <- true //if connection created successfully
}
