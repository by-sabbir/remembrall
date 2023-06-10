package db

import (
	"context"

	v1 "github.com/by-sabbir/remembrall/internal/types/v1"
	log "github.com/sirupsen/logrus"
)

func (d *DBClient) AddTask(ctx context.Context, t *v1.Task) (*v1.Task, error) {
	d.Client.Create(t)
	return t, nil
}

func (d *DBClient) ListTask(ctx context.Context) ([]v1.Task, error) {
	var tasks []v1.Task
	result := d.Client.Find(&tasks)

	log.Error(result.Error, result.RowsAffected)

	return tasks, nil
}
