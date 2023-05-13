package usecase

import (
	"context"
	"farm/backend/domain"
	"time"
)

type NotificationUC struct {
	repo domain.NotificationRepo
}

func NewNotificationUC(repo domain.NotificationRepo) NotificationUC {
	return NotificationUC{repo: repo}
}

func (c *NotificationUC) UpsertNotification(ctx context.Context, notification domain.Notification) error {
	return c.repo.UpsertNotification(ctx, notification)
}

func (c *NotificationUC) DeleteNotification(ctx context.Context, id string) error {
	return c.repo.DeleteNotification(ctx, id)
}

func (c *NotificationUC) GetAllNotifications(ctx context.Context) ([]domain.Notification, error) {
	return c.repo.GetAllNotification(ctx)
}

func (c *NotificationUC) GetNotificationByCowId(ctx context.Context, id string) (*domain.Notification, error) {
	return c.repo.GetNotificationByCowId(ctx, id)
}
func (c *NotificationUC) GetNotificationsByDate(ctx context.Context, date time.Time) ([]domain.Notification, error) {
	return c.repo.GetNotificationsByDate(ctx, date)
}
