# Proto Definition

gRPC is a high performance, open source framework developed by Google to handle remote procedure calls (RPCs). gRPC is Google's approach to a client-server application. It lets client and server applications communicate transparently, simplifying the process for developers to build connected systems.

gRPC calls requires a protobuf file to be used to implement / call the functions.

This folder will contain all the proto files and using the protoc tool they can be translated to any the protobuf file for gRPC calls.

R(emote) P(rocedure) C(all) are almost 7 times faster when receiving data and almost 10 times faster when sending data, in comparision with Rest.

Hence making them a goto for inter-service communication.

## How to generate a protobuf (.pb) file

This will generate protobuf files for golang.

To learn about other languages.

Read more at

`https://grpc.io/docs/languages/`

Note: proto/sample.proto refers to the location where we have stored the .proto we want to use. It will output 2 files \_grpc.pb.go and pb.go. pb.go file contains all the message description where as \_grpc.pb.go contains all the fucntion for grpc call that must be implemented

```cli
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/sample.proto
```

In case of

```cli
protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go_out: protoc-gen-go: Plugin failed with status code 1.
```

Download the protoc version macthing your OS from

`https://github.com/protocolbuffers/protobuf/releases`

UnZip the folder and place the bin/protoc to your GOPATH/bin

Then Run Commands to install protoc-gen-go

```cli
 go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
