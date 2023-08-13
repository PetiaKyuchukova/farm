package usecase

import (
	"context"
	"farm/backend/domain"
	"time"
)

type TaskUC struct {
	repo domain.TaskRepo
}

func NewTaskUC(repo domain.TaskRepo) TaskUC {
	return TaskUC{repo: repo}
}

func (c *TaskUC) UpsertTask(ctx context.Context, task domain.Task) error {
	return c.repo.UpsertTask(ctx, task)
}

func (c *TaskUC) UpdateTaskStatus(ctx context.Context, task domain.Task) error {
	return c.repo.UpdateTaskStatus(ctx, task.CowID, task.Date.Time, task.Done)
}

func (c *TaskUC) DeleteTask(ctx context.Context, id string) error {
	return c.repo.DeleteTask(ctx, id)
}

func (c *TaskUC) FetchAllTasks(ctx context.Context) ([]domain.Task, error) {
	return c.repo.GetAllTasks(ctx)
}

func (c *TaskUC) FetchTasksByDate(ctx context.Context, date time.Time) ([]domain.Task, error) {
	return c.repo.GetTasksByDate(ctx, date)
}
