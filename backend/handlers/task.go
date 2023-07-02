package handlers

import (
	"farm/backend/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//put handler to update IsPregnant to true (need cowId)
//put handler to update IsPregnant to false LastOvulation = today (0day of cow`s period) (need cowId)
//put handler to update the LastFertilization (need cowId)

type TaskHandler interface {
	GetTasksByDate(gc *gin.Context)
}

type defaultTaskHandler struct {
	taskUC usecase.TaskUC
}

func NewTaskHandler(task usecase.TaskUC) TaskHandler {
	return &defaultTaskHandler{task}
}

func (h *defaultTaskHandler) GetTasksByDate(gc *gin.Context) {
	dateString := gc.Query("date")

	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks, err := h.taskUC.FetchTasksByDate(gc.Request.Context(), date)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, tasks)
}
