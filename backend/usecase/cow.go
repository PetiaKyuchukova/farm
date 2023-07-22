package usecase

import (
	"context"
	"farm/backend/domain"
	"fmt"
	"sort"
)

type CowsUC struct {
	repo             domain.CowRepo
	pregnancyRepo    domain.PregnancyRepo
	inseminationRepo domain.InseminationRepo
}

func NewCowUC(repo domain.CowRepo, pregnancyRepo domain.PregnancyRepo, inseminationRepo domain.InseminationRepo) CowsUC {
	return CowsUC{
		repo:             repo,
		pregnancyRepo:    pregnancyRepo,
		inseminationRepo: inseminationRepo,
	}
}

func (c *CowsUC) UpsertCow(ctx context.Context, cow domain.Cow) error {
	for _, pregnancy := range cow.Pregnancies {
		err := c.pregnancyRepo.UpsertPregnancy(ctx, pregnancy, cow.ID)
		if err != nil {
			fmt.Errorf("error upserting pregnancy: %s", pregnancy)
			return err
		}
	}

	for _, insemination := range cow.Inseminations {
		err := c.inseminationRepo.UpsertInsemination(ctx, insemination, cow.ID)
		if err != nil {
			fmt.Errorf("error upserting inemination: %s", insemination)
			return err
		}
	}

	err := c.repo.UpsertCow(ctx, cow)
	if err != nil {
		fmt.Errorf("error upserting cow: %s", cow)
		return err
	}

	return nil
}

func (c *CowsUC) DeleteCowEntry(ctx context.Context, id string) error {
	err := c.repo.DeleteCow(ctx, id)
	if err != nil {
		fmt.Errorf("error deleting cow: %w", err)
		return err
	}

	err = c.pregnancyRepo.DeletePregnancies(ctx, id)
	if err != nil {
		fmt.Errorf("error deleting pregnancies of cow: %w", err)
		return err
	}

	err = c.inseminationRepo.DeleteInseminations(ctx, id)
	if err != nil {
		fmt.Errorf("error deleting inseminations of coe: %w", err)
		return err
	}

	return nil
}

func (c *CowsUC) GetAllCows(ctx context.Context) ([]domain.Cow, error) {
	cows, err := c.repo.GetAllCows(ctx)
	if err != nil {
		fmt.Errorf("error getting all cows: %w", err)
		return nil, err
	}

	for _, cow := range cows {
		pregnancy, err := c.pregnancyRepo.GetPregnanciesByCowID(ctx, cow.ID)
		if err != nil {
			fmt.Errorf("error getting prgnancy for cow: %w", err)
			return nil, err
		}

		insemination, err := c.inseminationRepo.GetInseminationsByCowID(ctx, cow.ID)
		if err != nil {
			fmt.Errorf("error getting insemination for cow: %w", err)
			return nil, err
		}

		cow.Pregnancies = pregnancy
		cow.Inseminations = insemination
	}

	return cows, nil
}

func (c *CowsUC) GetCowEntryById(ctx context.Context, id string) (*domain.Cow, error) {
	cow, err := c.repo.GetCowById(ctx, id)
	if err != nil {
		fmt.Errorf("error getting cow: %w", err)
		return nil, err
	}

	pregnancies, err := c.pregnancyRepo.GetPregnanciesByCowID(ctx, id)
	if err != nil {
		fmt.Errorf("err getting pregnancies: %w", err)
	}

	sort.Slice(pregnancies, func(aa, ab int) bool {
		return pregnancies[aa].DetectedAt.Time.Before(pregnancies[ab].DetectedAt.Time)
	})

	insemination, err := c.inseminationRepo.GetInseminationsByCowID(ctx, id)
	if err != nil {
		fmt.Errorf("err getting inseminations: %w", err)
	}

	sort.Slice(insemination, func(aa, ab int) bool {
		return insemination[aa].Date.Time.Before(insemination[ab].Date.Time)
	})

	cow.Pregnancies = pregnancies
	cow.Inseminations = insemination

	return cow, nil

}
