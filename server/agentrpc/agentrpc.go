package agentrpc

import (
	"context"

	"github.com/devilsec/btedr/proto/agentpb"
	"github.com/devilsec/btedr/proto/taskpb"
	"github.com/devilsec/btedr/server/db"
	"github.com/devilsec/btedr/server/util"
)

type AgentServer struct {
	agentpb.UnimplementedAgentRPCServer
	db *db.Database
}

func New(db *db.Database) *AgentServer {
	return &AgentServer{
		db: db,
	}
}

// TODO: Register implant with the database
func (server *AgentServer) Register(ctx context.Context, req *agentpb.Registration) (*agentpb.Empty, error) {
	util.Log.Infof("Received registration from %s\n", req.Ip)
	err := server.db.AddAgent(req)
	if err != nil {
		util.Log.Error(err)
	}
	return &agentpb.Empty{}, err
}

// TODO: Fetch the task from the database and send it to the agent
func (server *AgentServer) GetTask(ctx context.Context, req *agentpb.Request) (*taskpb.Task, error) {
	return &taskpb.Task{}, nil
}

// TODO: Log the agent's results
func (server *AgentServer) TaskResult(ctx context.Context, req *agentpb.Result) (*agentpb.Empty, error) {
	return &agentpb.Empty{}, nil
}
