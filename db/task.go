package db

import (
	"context"

	v1 "github.com/by-sabbir/remembrall/internal/types/v1"
	"github.com/minio/minio-go/v7"
	log "github.com/sirupsen/logrus"
)

func (d *DBClient) AddTask(ctx context.Context, t *v1.Task) (*v1.Task, error) {
	tx := d.Client.Create(t)
	info, err := d.Minio.FPutObject(ctx, "remembrall", "task.sqlite", "task.db", minio.PutObjectOptions{})
	if err != nil {
		log.Error("could not syncdb: ", err)
	} else {
		log.Info("synced db: ", info.Key)
	}
	return t, tx.Error
}

func (d *DBClient) ListTask(ctx context.Context) ([]v1.Task, error) {
	var tasks []v1.Task
	result := d.Client.Find(&tasks)

	return tasks, result.Error
}

func (d *DBClient) RemoveTask(ctx context.Context, id int) error {
	var t v1.Task
	t.ID = uint(id)
	tx := d.Client.Delete(&t)

	return tx.Error
}
