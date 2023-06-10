package v1

import "time"

type Task struct {
	ID        uint
	Title     string
	Status    string
	CreatedAt time.Time
	Deadline  time.Duration
}
