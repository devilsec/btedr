package models

// A single agent
type Agent struct {
	ID       string `gorm:"primaryKey"`
	Os       string
	Ip       string
	Hostname string
	User     string
	Groups   string
	OsUsers  string
}
