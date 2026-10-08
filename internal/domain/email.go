package domain

import "time"

type Email struct {
	UID     uint32
	From    string
	To      []string
	Subject string
	Body    string
	Date    time.Time
	Flags   []string
}
