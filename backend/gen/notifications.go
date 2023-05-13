package gen

import (
	"context"
	"farm/backend/domain"
	db "farm/backend/gen/db"
	"time"
)

type NotificationRepo struct {
	querier db.Querier
}

func NewNotificationRepo(querier db.Querier) *NotificationRepo {
	return &NotificationRepo{querier: querier}
}

func (r *NotificationRepo) UpsertNotification(ctx context.Context, notification domain.Notification) error {
	err := r.querier.UpsertNotification(ctx, db.UpsertNotificationParams{
		Cowid: notification.CowID,
		Date:  notification.Date,
		Type:  notification.Type,
		Text:  notification.Text,
	})
	return err
}

func (r *NotificationRepo) DeleteNotification(ctx context.Context, id string) error {
	err := r.querier.DeleteNotification(ctx, id)
	return err
}
func (r *NotificationRepo) GetAllNotification(ctx context.Context) ([]domain.Notification, error) {
	rows, err := r.querier.GetAllNotification(ctx)
	notifications := []domain.Notification{}
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		notifications = append(notifications, domain.Notification{
			CowID: row.Cowid,
			Date:  row.Date,
			Type:  row.Type,
			Text:  row.Text,
		})
	}
	return notifications, err
}
func (r *NotificationRepo) GetNotificationByCowId(ctx context.Context, cowId string) (*domain.Notification, error) {
	row, err := r.querier.GetNotificationByCowId(ctx, cowId)
	if err != nil {
		return nil, err
	}
	notification := domain.Notification{
		CowID: row.Cowid,
		Date:  row.Date,
		Type:  row.Type,
		Text:  row.Text,
	}

	return &notification, err
}

func (r *NotificationRepo) GetNotificationsByDate(ctx context.Context, date time.Time) ([]domain.Notification, error) {
	rows, err := r.querier.GetNotificationsByDate(ctx, date)
	if err != nil {
		return nil, err
	}
	notifications := []domain.Notification{}
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		notifications = append(notifications, domain.Notification{
			CowID: row.Cowid,
			Date:  row.Date,
			Type:  row.Type,
			Text:  row.Text,
		})
	}
	return notifications, err
}
