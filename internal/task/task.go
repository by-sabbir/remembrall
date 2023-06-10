package task

import (
	"context"

	v1 "github.com/by-sabbir/remembrall/internal/types/v1"
	log "github.com/sirupsen/logrus"
)

type TaskRepo interface {
	AddTask(context.Context, *v1.Task) (*v1.Task, error)
	ListTask(context.Context) ([]v1.Task, error)
	RemoveTask(context.Context, int) error
}

type TaskService struct {
	Repo TaskRepo
}

func NewTaskService(tp TaskRepo) *TaskService {
	return &TaskService{
		Repo: tp,
	}
}

func (ts *TaskService) AddTask(ctx context.Context, t *v1.Task) (*v1.Task, error) {

	createdTask, err := ts.Repo.AddTask(ctx, t)
	if err != nil {
		log.Error("could not create task: ", err)
		return &v1.Task{}, err
	}

	return createdTask, nil
}

func (ts *TaskService) ListTask(ctx context.Context) ([]v1.Task, error) {

	allTasks, err := ts.Repo.ListTask(ctx)
	if err != nil {
		log.Error("could not fetch task: ", err)
		return []v1.Task{}, err
	}

	return allTasks, nil
}

func (ts *TaskService) RemoveTask(ctx context.Context, id int) error {
	err := ts.Repo.RemoveTask(ctx, id)
	if err != nil {
		log.Error("could not remove task: ", err)
		return err
	}
	return nil
}
