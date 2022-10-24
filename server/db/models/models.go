package models

// A single agent
type Agent struct {
	ID       string `gorm:"primaryKey"`
  Os       string
	Ip       string
	Hostname string
	Uid      string
  Gid      string
}
