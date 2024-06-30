package services

import (
	"golang_project_base/config"
	"golang_project_base/models"
)

type HabitService struct{}

func (s *HabitService) InsertHabit(habit models.Habit) (models.Habit, error) {
	if err := config.DB.Create(&habit).Error; err != nil {
		return models.Habit{}, err
	}
	return habit, nil
}
