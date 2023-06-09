package v1

import "time"

type Task struct {
	ID        string
	Title     string
	Status    string
	CreatedAt time.Time
	Deadline  time.Duration
}
