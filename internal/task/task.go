package task

import (
	"context"

	v1 "github.com/by-sabbir/remembrall/internal/types/v1"
	log "github.com/sirupsen/logrus"
)

type TaskRepo interface {
	AddTask(context.Context, *v1.Task) (*v1.Task, error)
}

type TaskService struct {
	Repo TaskRepo
}

func NewTaskService(tp TaskRepo) *TaskService {
	return &TaskService{
		Repo: tp,
	}
}

func (ts TaskService) AddTask(ctx context.Context, t *v1.Task) (*v1.Task, error) {

	createdTask, err := ts.Repo.AddTask(ctx, t)
	log.Info("got task at internal: ", t)
	if err != nil {
		log.Error("could not create task: ", err)
		return &v1.Task{}, err
	}

	return createdTask, nil
}
