package v1

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title    string
	Status   string
	Deadline time.Duration
}
