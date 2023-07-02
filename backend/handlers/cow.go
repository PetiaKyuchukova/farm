package handlers

import (
	"farm/backend/domain"
	"farm/backend/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	UpsertCow(gc *gin.Context)
	DeleteCow(gc *gin.Context)
	GetAllCows(gc *gin.Context)
	GetCowById(gc *gin.Context)
	LivenessHandler(gc *gin.Context)
}
type defaultHandler struct {
	uc usecase.CowsUC
}

func NewHandler(uc usecase.CowsUC) Handler {
	return &defaultHandler{uc}
}
func (h *defaultHandler) LivenessHandler(gc *gin.Context) {
	gc.String(http.StatusOK, "all ok from farm manager")
}
func (h *defaultHandler) UpsertCow(gc *gin.Context) {
	var cow domain.Cow

	if err := gc.BindJSON(&cow); err != nil {
		fmt.Errorf("error binding json %w", err)
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.uc.UpsertCow(gc.Request.Context(), cow)
	if err != nil {
		fmt.Errorf("error upserting cow: %w", err)
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, cow)
}

func (h *defaultHandler) DeleteCow(gc *gin.Context) {
	id := gc.Param("id")

	err := h.uc.DeleteCowEntry(gc.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error deleting cow: %w", err)
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, "Successfully deleted!")
}

func (h *defaultHandler) GetAllCows(gc *gin.Context) {
	gc.Header("Access-Control-Allow-Origin", "http://localhost:5173")

	cows, err := h.uc.GetAllCows(gc.Request.Context())
	if err != nil {
		fmt.Errorf("error getting all cows.sql: %w", err)
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, cows)
}
func (h *defaultHandler) GetCowById(gc *gin.Context) {
	id := gc.Param("id")

	cow, err := h.uc.GetCowEntryById(gc.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error getting cow by id: %w", err)
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, cow)
}
