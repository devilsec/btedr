// Implements a user-side cli.
// The CLI sends requests to the Operator gRPC server (either in-memory, or over TCP if the user is a client operator)
package console

import (
	"context"
	"net"
	"path/filepath"

	"github.com/desertbit/grumble"
	"github.com/devilsec/btedr/proto/operatorpb"
	"github.com/devilsec/btedr/server/util"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type Console struct {
	App *grumble.App
	Rpc operatorpb.OperatorRPCClient
}

// Create a new console and connect to an in-memory gRPC server
func New(root string, listener *bufconn.Listener) Console {
	app := grumble.New(&grumble.Config{
		Name:                  "BTEDR",
		Description:           "A visibility tool for rapid, panicky, and last minute incident response.",
		HistoryFile:           filepath.Join(root, ".btedr_history"),
		Prompt:                "❯ ",
		PromptColor:           color.New(),
		HelpHeadlineColor:     color.New(),
		HelpHeadlineUnderline: true,
		HelpSubCommands:       true,
	})

	console := Console{
		App: app,
		Rpc: rpcClient(listener),
	}

	addCommands(console)

	return console
}

func rpcClient(listener *bufconn.Listener) operatorpb.OperatorRPCClient {
	ctxDialer := grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		return listener.Dial()
	})

	// Dial in to the in-memory grpc server
	options := []grpc.DialOption{
		ctxDialer,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// TODO: Set the maximum size of an in-memory message
		grpc.WithBlock(),
	}
	conn, err := grpc.DialContext(context.Background(), "", options...)
	if err != nil {
		util.Log.Fatal(err)
	}
	return operatorpb.NewOperatorRPCClient(conn)
}

// TODO: Make the prompt more fancy, to reflect the state of the program
// func (c Console) getPrompt() string {
//   return "❯"
// }
