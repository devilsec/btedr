package db

import (
	"errors"

	"github.com/devilsec/btedr/proto/taskpb"
)

// Pop a task from an agent
func (db *database) PopTask(agent string) (*taskpb.Task, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

  // Get the queue. Return some error if the agent doesn't have a queue
  queue, ok := db.tasks[agent]
  if !ok {
    return nil, errors.New("Agent not registered")
  }

  // Get the first task in the queue. Return some error if the queue is empty
  if len(queue) > 0 {
    task := queue[0]
    db.tasks[agent] = db.tasks[agent][1:]
    return task, nil
  } else {
    return nil, errors.New("Queue is empty")
  }
}

// TODO: Get an agent from Db
func (db *database) Agent() {
	db.mu.Lock()
	defer db.mu.Unlock()
	// ...
}

// TODO: Add an agent to the Db
func (db *database) AddAgent() {
	db.mu.Lock()
	defer db.mu.Unlock()
}

// TODO: Add a Task to the Db
func (db *database) AddTask() {

}
