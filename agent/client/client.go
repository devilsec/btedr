// Create a client to connect the agent to the server
package client

import (
	"github.com/devilsec/btedr/proto/agentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// A Client connects to a gRPC server to retrive tasks
type Client struct {
	Rpc    agentpb.AgentRPCClient
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

	rpc := agentpb.NewAgentRPCClient(dial)
	client := Client{
		Rpc:    rpc,
		Server: server,
		Dial:   dial,
	}

	return client, nil
}

// TODO: Create a gRPC request to register the agent with the server
func (Client) Register() {

}
