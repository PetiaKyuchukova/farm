package handlers

import (
	"farm/backend/domain"
	"farm/backend/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MilkHandler interface {
	FetchMilkSeriesInTimeframe(gc *gin.Context)
	UpsertMilk(gc *gin.Context)
}

type defaultMilkHandler struct {
	milkUC usecase.MilkUC
}

func NewMilkHandler(milk usecase.MilkUC) MilkHandler {
	return &defaultMilkHandler{milkUC: milk}
}

func (m *defaultMilkHandler) FetchMilkSeriesInTimeframe(gc *gin.Context) {
	fromString := gc.Query("from")
	toString := gc.Query("to")

	from, err := time.Parse("2006-01-02", fromString)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	to, err := time.Parse("2006-01-02", toString)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	milk, err := m.milkUC.GetMilkInTimeframe(gc.Request.Context(), from, to)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, milk)
}

func (h *defaultMilkHandler) UpsertMilk(gc *gin.Context) {
	var milk domain.Milk

	if err := gc.BindJSON(&milk); err != nil {
		fmt.Errorf("error binding json %w", err)
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.milkUC.UpsertMilk(gc.Request.Context(), milk)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, milk)
}
