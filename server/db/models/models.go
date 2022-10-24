package models

// A single implant
type Implant struct {
	ID       string `gorm:"primaryKey"`
  Os       string
	Ip       string
	Hostname string
	Uid      string
  Gid      string
}
