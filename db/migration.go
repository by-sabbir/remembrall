package db

import (
	log "github.com/sirupsen/logrus"

	v1 "github.com/by-sabbir/remembrall/internal/types/v1"
)

func (d *DBClient) Migrate() error {

	var task v1.Task

	if err := d.Client.AutoMigrate(task); err != nil {
		log.Error("could not migrate db: ", err)
		return err
	}

	return nil
}
