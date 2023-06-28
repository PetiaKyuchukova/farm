package gen

import (
	"context"
	"farm/backend/domain"
	db "farm/backend/gen/db"
	"time"
)

type TaskRepo struct {
	querier db.Querier
}

func NewTaskRepo(querier db.Querier) *TaskRepo {
	return &TaskRepo{querier: querier}
}

func (r *TaskRepo) UpsertTask(ctx context.Context, task domain.Task) error {
	err := r.querier.UpsertTasks(ctx, db.UpsertTasksParams{
		Cowid: task.CowID,
		Date:  task.Date,
		Type:  task.Type,
		Text:  task.Text,
	})
	return err
}

func (r *TaskRepo) DeleteTask(ctx context.Context, id string) error {
	err := r.querier.DeleteTask(ctx, id)
	return err
}
func (r *TaskRepo) GetAllTask(ctx context.Context) ([]domain.Task, error) {
	rows, err := r.querier.GetAllTasks(ctx)
	notifications := []domain.Task{}
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		notifications = append(notifications, domain.Task{
			CowID: row.Cowid,
			Date:  row.Date,
			Type:  row.Type,
			Text:  row.Text,
		})
	}
	return notifications, err
}
func (r *TaskRepo) GetTaskByCowId(ctx context.Context, cowId string) (*domain.Task, error) {
	row, err := r.querier.GetTaskByCowId(ctx, cowId)
	if err != nil {
		return nil, err
	}
	notification := domain.Task{
		CowID: row.Cowid,
		Date:  row.Date,
		Type:  row.Type,
		Text:  row.Text,
	}

	return &notification, err
}

func (r *TaskRepo) GetTasksByDate(ctx context.Context, date time.Time) ([]domain.Task, error) {
	rows, err := r.querier.GetTasksByDate(ctx, date)
	if err != nil {
		return nil, err
	}
	notifications := []domain.Task{}
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		notifications = append(notifications, domain.Task{
			CowID: row.Cowid,
			Date:  row.Date,
			Type:  row.Type,
			Text:  row.Text,
		})
	}
	return notifications, err
}
