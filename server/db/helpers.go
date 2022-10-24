package db

import (
	"errors"

	"github.com/devilsec/btedr/proto/taskpb"
)

// Pop a task from an implant
func (db *database) PopTask(implant string) (*taskpb.Task, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

  // Get the queue. Return some error if the implant doesn't have a queue
  queue, ok := db.tasks[implant]
  if !ok {
    return nil, errors.New("Implant not registered")
  }

  // Get the first task in the queue. Return some error if the queue is empty
  if len(queue) > 0 {
    task := queue[0]
    db.tasks[implant] = db.tasks[implant][1:]
    return task, nil
  } else {
    return nil, errors.New("Queue is empty")
  }
}

// TODO: Get an Implant from Db
func (db *database) Implant() {
	db.mu.Lock()
	defer db.mu.Unlock()
	// ...
}

// TODO: Add an Implant to the Db
func (db *database) AddImplant() {
	db.mu.Lock()
	defer db.mu.Unlock()
}

// TODO: Add a Task to the Db
func (db *database) AddTask() {

}
