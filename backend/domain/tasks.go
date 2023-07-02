package domain

import (
	"context"
	"time"
)

const (
	FertilizationType     = "artificial inseminated"
	PregnantType          = "pregnancy check"
	DryPeriodAfter15dType = "dry period - after 15 days"
	DryPeriodStartType    = "dry period - start"
	GivingBirthType       = "giving birth"
	OvulationType         = "ovulation"

	OvulationAfterFertilizationText = "Today is expected ovulation day of the cow but we made Artificial insemination after last ovulation, is it really in ovulation? will we make Artificial insemination?"
	FertilizationText               = "Today is ovulation day of tha cow and it is ready to be artificial inseminated. Did you inseminated it?"
	PregnantText                    = "Today is time to check is the cow pregnant. Is it pregnant?"
	DryPeriodAfter15dText           = "Left 15 days until the start of the dry period"
	DryPeriodStartText              = "Today should start the dry period. 60 days left to expected giving birth date."
	GivingBirthText                 = "Today we expect the cow is going to giving birth. Did the cow gave birth today?"
)

type Task struct {
	CowID string    `json:"cow_id"`
	Date  time.Time `json:"date"`
	Type  string    `json:"type"`
	Text  string    `json:"text"`
}

type TaskRepo interface {
	UpsertTask(ctx context.Context, task Task) error
	DeleteTask(ctx context.Context, id string) error
	GetAllTasks(ctx context.Context) ([]Task, error)
	GetTasksByDate(ctx context.Context, date time.Time) ([]Task, error)
}
