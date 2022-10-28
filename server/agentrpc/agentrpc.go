package agentrpc

import (
	"context"

	"github.com/devilsec/btedr/proto/agentpb"
	"github.com/devilsec/btedr/proto/taskpb"
)

type AgentServer struct {
	agentpb.UnimplementedAgentRPCServer
}

// TODO: Register implant with the database
func (server *AgentServer) Register(ctx context.Context, req *agentpb.Registration) (*agentpb.Empty, error) {
  return &agentpb.Empty{}, nil
}

// TODO: Fetch the task from the database and send it to the agent
func (server *AgentServer) GetTask(ctx context.Context, req *agentpb.Request) (*taskpb.Task, error) {
  return &taskpb.Task{}, nil
}

// TODO: Log the agent's results
func (server *AgentServer) TaskResult(ctx context.Context, req *agentpb.Result) (*agentpb.Empty, error) {
  return &agentpb.Empty{}, nil
}
