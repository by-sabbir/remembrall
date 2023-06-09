package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBClient struct {
	Client *gorm.DB
}

func NewDBClient() (*DBClient, error) {
	db, err := gorm.Open(sqlite.Open("task.sqlite"), &gorm.Config{})
	if err != nil {
		log.Error("could not open db", err)
		return &DBClient{
			Client: &gorm.DB{},
		}, err
	}
	return &DBClient{
		Client: db,
	}, nil
}
