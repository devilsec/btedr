package db

import (
	"path/filepath"
	"sync"

	"github.com/devilsec/btedr/server/db/models"
	"github.com/devilsec/btedr/server/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Gorm is thread-safe, but this mutex should be locked when making successive calls to the Db
var Mu = &sync.Mutex{}

// Create a sqlite3 database
var Db = func() *gorm.DB {
	// Stored in $HOME/.btedr/btedr.db
	db, err := gorm.Open(sqlite.Open(filepath.Join(util.Root, "btedr.db")), &gorm.Config{})
	if err != nil {
		util.Log.Fatal(err)
	}

	// Create tables for the Tasks and Implants
	if err = db.AutoMigrate(&models.Task{}, &models.Implant{}); err != nil {
		util.Log.Fatal(err)
	}
	return db
}()
