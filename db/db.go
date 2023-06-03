package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Client struct {
}

func NewDBClient() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("task.db"), &gorm.Config{})
	if err != nil {
		log.Error("could not open db", err)
		return &gorm.DB{}, err
	}

	return db, err
}
