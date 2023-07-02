package worker

import (
	"context"
	"farm/backend/domain"
	"farm/backend/usecase"
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

type Worker struct {
	taskUC usecase.TaskUC
	cowUC  usecase.CowsUC
}

func NewWorker(taskUC usecase.TaskUC, cowUC usecase.CowsUC) *Worker {
	return &Worker{taskUC: taskUC, cowUC: cowUC}
}

func (w *Worker) Schedule(ctx context.Context, cronExpression string) error {
	s := gocron.NewScheduler(time.UTC)

	_, err := s.Cron(cronExpression).SingletonMode().Do(func() {
		w.TaskWorker(ctx)
	})
	if err != nil {
		return fmt.Errorf("error setting up cron scheduler. caused by - %v", err)
	}

	s.StartAsync()

	return nil
}

func (w *Worker) TaskWorker(ctx context.Context) {
	//get all cows.sql from DB
	cows, err := w.cowUC.GetAllCows(ctx)
	if err != nil {
		if err != nil {
			fmt.Errorf("error getting all cows.sql from database: %w", err)
			return
		}
	}

	today := time.Now()
	day := time.Hour * 24
	//save alarms to db then frontend will fetch them for current day
	for _, cow := range cows {
		if len(cow.Inseminations) == 0 {
			continue
		}

		if cow.Inseminations[0].Date.Add(36*day) == today {
			//sent alarm: to check pregnancy. Is this cow pregnant? --> true or false (if user press yes, sent a put req to update IsPregnant to true)
			w.taskUC.UpsertTask(ctx, domain.Task{
				CowID: cow.ID,
				Date:  time.Now(),
				Type:  domain.PregnantType,
				Text:  domain.PregnantText,
			})

		} else if cow.IsPregnant && cow.Inseminations[0].Date.Add(208*day) == today {
			//sent alarm: dry period after 15 days
			w.taskUC.UpsertTask(ctx, domain.Task{
				CowID: cow.ID,
				Date:  time.Now(),
				Type:  domain.DryPeriodAfter15dType,
				Text:  domain.DryPeriodAfter15dText,
			})
		} else if cow.IsPregnant && cow.Inseminations[0].Date.Add(223*day) == today {
			//sent alarm: time for dry period, birth after 60 days
			w.taskUC.UpsertTask(ctx, domain.Task{
				CowID: cow.ID,
				Date:  time.Now(),
				Type:  domain.DryPeriodStartType,
				Text:  domain.DryPeriodStartText,
			})
		} else if cow.IsPregnant && cow.Inseminations[0].Date.Add(283*day) == today {
			//sent alarm: today we expect birth! did the cow give birth today? --> true or false (if user press yes, sent a put req to update IsPregnant to false LastOvulation = today )
			//should add func to set give birthdate manually

			w.taskUC.UpsertTask(ctx, domain.Task{
				CowID: cow.ID,
				Date:  time.Now(),
				Type:  domain.GivingBirthType,
				Text:  domain.GivingBirthText,
			})
		}

		if !cow.IsPregnant && cow.Inseminations[0].Date.Before(cow.Ovulation) && cow.Ovulation.Add(21*day) == today {
			cow.Ovulation = today
			//if the cow is not pregnant AND we did not make Artificial insemination on last ovulation AND today is 21 days after the last ovu - today is ovulation
			//sent alarm: today is {cow number} ovulation day, will we make Artificial insemination? --> true or false (if user press true sent a put req that will update the LastFertilization)
			w.taskUC.UpsertTask(ctx, domain.Task{
				CowID: cow.ID,
				Date:  time.Now(),
				Type:  domain.FertilizationType,
				Text:  domain.FertilizationText,
			})
		} else if !cow.IsPregnant && cow.Inseminations[0].Date.After(cow.Ovulation) && cow.Ovulation.Add(21*day) == today {
			cow.Ovulation = today
			//sent alarm: today is {cow number} ovulation day, we made Artificial insemination after last ovulation, is it really in ovulation? will we make Artificial insemination? --> true or false (if user press true sent a put req that will update the LastFertilization)
			w.taskUC.UpsertTask(ctx, domain.Task{
				CowID: cow.ID,
				Date:  time.Now(),
				Type:  domain.OvulationType,
				Text:  domain.OvulationAfterFertilizationText,
			})
		}
	}
}
