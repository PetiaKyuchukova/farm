package handlers

import (
	"farm/backend/usecase"
)

//put handler to update IsPregnant to true (need cowId)
//put handler to update IsPregnant to false LastOvulation = today (0day of cow`s period) (need cowId)
//put handler to update the LastFertilization (need cowId)

type NotificationHandler interface {
}

type defaultNotificationHandler struct {
	uc *usecase.CowsUC
}

func NewNotificationHandler(uc *usecase.CowsUC) NotificationHandler {
	return &defaultNotificationHandler{uc}
}
