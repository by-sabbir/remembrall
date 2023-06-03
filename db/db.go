package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Client struct {
	DB *gorm.DB
}

func NewDBClient() (*Client, error) {
	db, err := gorm.Open(sqlite.Open("task.db"), &gorm.Config{})
	if err != nil {
		log.Error("could not open db", err)
		return &Client{
			DB: &gorm.DB{},
		}, err
	}

	return &Client{
		DB: db,
	}, nil
}
