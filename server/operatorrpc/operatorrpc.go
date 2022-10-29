// This implements the operator gRPC service.
// The server console communicates with an in-memory Operator gRPC server.
// Remote clients/operators will communicate with an Operator gRPC server over a TCP port instead.
package operatorrpc

import (
	"context"
	"fmt"
	"net"

	"github.com/devilsec/btedr/proto/agentpb"
	"github.com/devilsec/btedr/proto/operatorpb"
	"github.com/devilsec/btedr/proto/taskpb"
	"github.com/devilsec/btedr/server/agentrpc"
	"github.com/devilsec/btedr/server/db"
	"github.com/devilsec/btedr/server/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

// Implements the operator gRPC service functions
type OperatorServer struct {
	operatorpb.UnimplementedOperatorRPCServer
	db *db.Database
}

// Start a gRPC server over an in-memory buffer
// This is to be used only by the server for communicating with itself
func Start(db *db.Database) (*grpc.Server, *bufconn.Listener) {
	// TODO: Create constants.go to store this default value
	// 2 MiB buffer
	listener := bufconn.Listen(2 * 1024 * 1024)

	// TODO: Set the maximum size of a gRPC message
	srv := grpc.NewServer()
	operatorpb.RegisterOperatorRPCServer(srv, &OperatorServer{
		db: db,
	})

	go func() {
		if err := srv.Serve(listener); err != nil {
			util.Log.Fatal(err)
		}
	}()

	return srv, listener
}

// RPC Service for starting a listener for agents on a specified port
func (server *OperatorServer) Start(ctx context.Context, req *operatorpb.StartReq) (*operatorpb.StartResp, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", req.GetPort()))
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer()
	agentpb.RegisterAgentRPCServer(s, agentrpc.New(server.db))

	// Start listening for agents
	// TODO: Create a `Job` struct, to notify when this job has ended
	go func() {
		if err := s.Serve(listener); err != nil {
			// TODO: Get rid of printlns here, send error to console over the Job's channel as an event
			fmt.Println()
			util.Log.Error(err)
		}
	}()

	resp := &operatorpb.StartResp{
		Status: fmt.Sprintf("Status: started on :%d", req.GetPort()),
	}
	return resp, nil
}

// RPC Service for pinging an agent
// TODO: Communicate with an agent
func (server *OperatorServer) Ping(ctx context.Context, req *taskpb.PingReq) (*operatorpb.PingResp, error) {
	resp := &operatorpb.PingResp{
		Roundtrip: 1337,
		Error:     "",
	}

	return resp, nil
}
