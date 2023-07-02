package usecase

import (
	"context"
	"farm/backend/domain"
	"time"
)

type NotificationUC struct {
	repo domain.TaskRepo
}

func NewNotificationUC(repo domain.TaskRepo) NotificationUC {
	return NotificationUC{repo: repo}
}

func (c *NotificationUC) UpsertTask(ctx context.Context, notification domain.Task) error {
	return c.repo.UpsertTask(ctx, notification)
}

func (c *NotificationUC) DeleteTask(ctx context.Context, id string) error {
	return c.repo.DeleteTask(ctx, id)
}

func (c *NotificationUC) FetchAllTasks(ctx context.Context) ([]domain.Task, error) {
	return c.repo.GetAllTasks(ctx)
}

func (c *NotificationUC) FetchTasksByDate(ctx context.Context, date time.Time) ([]domain.Task, error) {
	return c.repo.GetTasksByDate(ctx, date)
}
