package domain

import (
	"context"
	"time"
)

const (
	FertilizationType     = "AI"
	PregnantType          = "PC"
	DryPeriodAfter15dType = "DP15d"
	DryPeriodStartType    = "DPS"
	GivingBirthType       = "GB"
	OvulationType         = "OVU"

	OvulationAfterFertilizationText = "Expected OVULATION day. ARTIFICIAL INSEMINATION after last ovulation"
	FertilizationText               = "Expected OVULATION day"
	PregnantText                    = "Check cow PREGNANCY"
	DryPeriodAfter15dText           = "Left 15 days until start of the DRY PERIOD"
	DryPeriodStartText              = "Start DRY PERIOD. 60 days left to expected giving birth date."
	GivingBirthText                 = "Expect GIVING BIRTH"
)

type Task struct {
	CowID string     `json:"cow_id"`
	Date  CustomTime `json:"date"`
	Type  string     `json:"type"`
	Text  string     `json:"text"`
	Done  bool       `json:"done"`
}

type TaskRepo interface {
	UpsertTask(ctx context.Context, task Task) error
	DeleteTask(ctx context.Context, id string) error
	GetAllTasks(ctx context.Context) ([]Task, error)
	GetTasksByDate(ctx context.Context, date time.Time) ([]Task, error)
	UpdateTaskStatus(ctx context.Context, cowId string, date time.Time, done bool) error
}
