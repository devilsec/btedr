package console

import (
	"context"
	"fmt"
	"time"

	"github.com/desertbit/grumble"
	"github.com/devilsec/btedr/proto/operatorpb"
	"github.com/devilsec/btedr/proto/taskpb"
)

// Add user commands for the cli
func addCommands(console Console) {
  ctx := context.Background()
  console.App.AddCommand(&grumble.Command{
    Name: "ping",
    Help: "Check if an implant is alive.",
    Flags: func(f *grumble.Flags) {
      // TODO: Create constants.go for default timeout
      f.Duration("t", "timeout", 5*time.Second, "ping timeout duration")
    },
    Run: func(c *grumble.Context) error {
      response, err := console.Rpc.Ping(ctx, &taskpb.PingReq{
        Duration: uint32(c.Flags.Duration("timeout")),
      })

      if err != nil {
        console.App.Println(err.Error())
        return err
      } else {
        console.App.Println(fmt.Sprintf("Roundtrip: %d", response.GetRoundtrip()))
        return nil
      }
    },
  })

  console.App.AddCommand(&grumble.Command{
    Name: "start",
    Help: "Start listening for implants",
    Flags: func(f *grumble.Flags) {
      // TODO: Create constants.go for default port
      f.Int("p", "port", 50051, "Specify listening port")
    },
    Run: func(c *grumble.Context) error {
      response, err := console.Rpc.Start(ctx, &operatorpb.StartReq{
        Port: uint32(c.Flags.Int("port")),
      })

      if err != nil {
        console.App.Println(err.Error())
        return err
      } else {
        console.App.Println(fmt.Sprintf("Started listener on :%s", response.GetStatus()))
        return nil
      }
    },
  })
}
