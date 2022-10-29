package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/devilsec/btedr/proto/agentpb"
	"github.com/devilsec/btedr/proto/taskpb"
	"github.com/devilsec/btedr/server/db/models"
	"github.com/google/uuid"
)

// Pop a task from an agent
func (db *Database) PopTask(agent string) (*taskpb.Task, error) {
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
func (db *Database) Agent() {
	db.mu.Lock()
	defer db.mu.Unlock()
	// ...
}

// Add an agent to the Db
func (db *Database) AddAgent(agent *agentpb.Registration) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Generate a new UUID for the agent
	id := uuid.New().String()
	if id == "" {
		return errors.New("Could not create new UUID.")
	}

	// Add the agent to the db
	db.orm.Create(&models.Agent{
		ID:       id,
		Os:       agent.Os,
		Ip:       agent.Ip,
		Hostname: agent.Hostname,
		User:     users2string([]*agentpb.User{agent.User}),
		Groups:   users2string(agent.Groups),
		OsUsers:  users2string(agent.Users),
	})

	return nil
}

// TODO: Add a Task to the Db
func (db *Database) AddTask() {

}

func users2string(users []*agentpb.User) string {
	g := []string{}
	for _, group := range users {
		g = append(g, fmt.Sprintf("%d(%s)", group.Id, group.Name))
	}
	return strings.Join(g, ",")
}
