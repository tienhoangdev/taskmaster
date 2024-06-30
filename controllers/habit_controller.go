package controllers

import (
	"golang_project_base/models"
	"golang_project_base/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HabitController struct {
	Service *services.HabitService
}

func (h *HabitController) CreateHabit(c *gin.Context) {
	var habit models.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate request payload
	if err := habit.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	createdHabit, err := h.Service.InsertHabit(habit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdHabit)
}
