// Create a client to connect the implant to the server
package client

import (
	"github.com/devilsec/btedr/proto/implantpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// A Client connects to a gRPC server to retrive tasks
type Client struct {
	Rpc    implantpb.ImplantRPCClient
	Server Server
	Dial   *grpc.ClientConn
}

// Create a client and connect to the server
func New(serverIP string, serverPort int16) (Client, error) {
	server := Server{
		ip:   serverIP,
		port: serverPort,
	}

	// Connect to the server
	// TODO: Connect via HTTPS (mutual TLS?)
	dial, err := grpc.Dial(server.ip, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return Client{}, nil
	}

	rpc := implantpb.NewImplantRPCClient(dial)
	client := Client{
		Rpc:    rpc,
		Server: server,
		Dial:   dial,
	}

	return client, nil
}

// TODO: Create a gRPC request to register the implant with the server
func (Client) Register() {

}
