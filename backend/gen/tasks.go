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
		Date:  task.Date.Time,
		Type:  task.Type,
		Text:  task.Text,
		Done:  makeNullBool(&task.Done),
	})
	return err
}
func (r *TaskRepo) DeleteTask(ctx context.Context, id string) error {
	err := r.querier.DeleteTask(ctx, id)
	return err
}
func (r *TaskRepo) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	rows, err := r.querier.GetAllTasks(ctx)
	tasks := []domain.Task{}
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		tasks = append(tasks, domain.Task{
			CowID: row.Cowid,
			Date:  domain.CustomTime{row.Date},
			Type:  row.Type,
			Text:  row.Text,
			Done:  row.Done.Bool,
		})
	}
	return tasks, err
}
func (r *TaskRepo) GetTasksByDate(ctx context.Context, date time.Time) ([]domain.Task, error) {
	rows, err := r.querier.GetTasksByDate(ctx, date)
	if err != nil {
		return nil, err
	}
	tasks := []domain.Task{}
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		tasks = append(tasks, domain.Task{
			CowID: row.Cowid,
			Date:  domain.CustomTime{row.Date},
			Type:  row.Type,
			Text:  row.Text,
			Done:  row.Done.Bool,
		})
	}
	return tasks, err
}

func (r *TaskRepo) UpdateTaskStatus(ctx context.Context, cowId string, date time.Time, done bool) error {
	err := r.querier.UpdateTaskStatus(ctx, db.UpdateTaskStatusParams{
		Done:  makeNullBool(&done),
		Cowid: cowId,
		Date:  date,
	})
	if err != nil {
		return err
	}

	return nil
}
