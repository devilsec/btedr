// Utility package for useful project-wide variables and functions
package util

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// A logger for the project.
// Saves logs to $HOME/.btedr/log.txt
var Log = logrus.New()

// Root directory for files
// Defaults to $HOME/.btedr/
var Root = func () string {
	// Create a HOME/.btedr/ directory
	dir, err := os.UserHomeDir()
	if err != nil {
    // Use current dir if HOME diesn't exist
		if dir, err = os.Getwd(); err != nil {
			Log.Fatal(err)
		}
	}

	dir = filepath.Join(dir, ".btedr")
	if err := os.MkdirAll(dir, 0700); err != nil {
		Log.Fatal(err)
	}

  // Create log file
  if file, err := os.OpenFile(filepath.Join(dir, "log.txt"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600); err != nil {
    // Don't bother logging if OpenFile fails
    Log.Out = io.Discard
  } else {
    Log.Out = file
  }

  return dir
}()
