package models

import "github.com/devilsec/btedr/proto/taskpb"

// A single task to be executed by an implant
type Task struct {
	ID        string `gorm:"primaryKey"`
	Type      taskpb.Type
	ImplantId string // The implant that will execute this task
}

// A single implant
type Implant struct {
	ID       string `gorm:"primaryKey"`
	Ip       string
	User     string
	Hostname string
	TaskId   string // The task the implant will execute next
}
