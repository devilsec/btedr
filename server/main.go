package main

import (
	"github.com/devilsec/btedr/server/console"
	"github.com/devilsec/btedr/server/db"
	"github.com/devilsec/btedr/server/operatorrpc"
	"github.com/devilsec/btedr/server/util"
)

func main() {
	db, err := db.New()
	if err != nil {
		util.Log.Fatal(err)
	}

	srv, listener := operatorrpc.Start(db)
	defer srv.GracefulStop()
	c := console.New(util.Root, listener)

	if err := c.App.Run(); err != nil {
		util.Log.Fatal(err)
	}
}
