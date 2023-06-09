package db

import (
	"context"

	v1 "github.com/by-sabbir/remembrall/internal/types/v1"
)

func (d *DBClient) AddTask(ctx context.Context, t *v1.Task) (*v1.Task, error) {

	d.Client.Create(t)
	return t, nil
}
