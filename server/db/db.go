package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"
	"sync"

	"github.com/devilsec/btedr/proto/taskpb"
	"github.com/devilsec/btedr/server/db/models"
	"github.com/devilsec/btedr/server/util"
)

type Database struct {
	mu  sync.Mutex
	orm *gorm.DB
	// Key of agent_id and value of Task queue
	tasks map[string]([]*taskpb.Task)
}

func New() (*Database, error) {
	// Create a sqlite3 database stored in btedr.orm
	orm, err := gorm.Open(sqlite.Open(filepath.Join(util.Root, "btedr.db")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Create tables for the Agents
	if err = orm.AutoMigrate(&models.Agent{}); err != nil {
		return nil, err
	}

	db := &Database{
		orm:   orm,
		tasks: make(map[string]([]*taskpb.Task)),
	}

	return db, nil
}
