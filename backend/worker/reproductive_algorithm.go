package worker

import (
	"farm/backend/domain"
	"time"
)

func AlarmWorker() {
	//get all cows from DB
	cows := []domain.Cow{}
	today := time.Now()

	day := time.Hour * 24
	//тези аларми ги записваме в базата данни и след това фронтенда събира всички аларми за този ден
	for _, cow := range cows {
		if cow.LastFertilization.Add(36*day) == today {
			//sent alarm to check pregnancy
		} else if cow.IsPregnant && cow.LastFertilization.Add(208*day) == today {
			//sent alarm засушаване след 15 дни
		} else if cow.IsPregnant && cow.LastFertilization.Add(223*day) == today {
			//sent alarm време е за засушаване, след 60 дни кравата очакваме да роди
		} else if cow.IsPregnant && cow.LastFertilization.Add(283*day) == today {
			//sent alarm очакваме крават днес да роди
		}

		//ако кравата не е бременна И след последната овулация нямаме заплождане и днес е 21 дена след предната овулация прави аларма кравата е в овулация
		if !cow.IsPregnant && cow.LastFertilization.Before(cow.LastOvulation) && cow.LastOvulation.Add(21*day) == today {
			//sent alarm Кравата е в овулация и готова за заплождане ще оплождане ли ако да сетваме датата на оплождане
		} else if !cow.IsPregnant && cow.LastFertilization.After(cow.LastOvulation) && cow.LastOvulation.Add(21*day) == today {
			//изпрашаме аларма кравата не е бреммена заплодили сме я след последната овулация, но днес е окачвана овулация ако не е бременна. Имаме ли овулация???ако да правим датата на
		}
	}
}