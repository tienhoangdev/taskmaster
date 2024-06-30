package models

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Habit struct {
	gorm.Model
	HabitName     string         `json:"habit_name" binding:"required"`
	Configuration datatypes.JSON `json:"configuration"`
	UserId        int            `json:"user_id" binding:"required"`
	Description   string         `json:"description"`
}

type HabitConfiguration struct {
	Type      string   `json:"type" binding:"required, oneof=daily weekly monthly"`
	Exception []string `json:"exception"`
	Days      []string `json:"days"`
}

func (h *Habit) Validate() error {
	validate := validator.New()

	if err := validate.Struct(h); err != nil {
		return err
	}

	var config HabitConfiguration
	if err := json.Unmarshal(h.Configuration, &config); err != nil {
		return err
	}

	if err := validate.Struct(&config); err != nil {
		return err
	}

	return nil
}

// Method to set default values before creating a record
func (h *Habit) BeforeCreate(tx *gorm.DB) (err error) {
	var config HabitConfiguration
	if err := json.Unmarshal(h.Configuration, &config); err != nil {
		return err
	}

	// Set default values if not provided
	if config.Exception == nil {
		config.Exception = []string{}
	}
	if config.Days == nil {
		config.Days = []string{}
	}

	// Marshal the updated configuration back to JSON
	updatedConfig, err := json.Marshal(config)
	if err != nil {
		return err
	}

	h.Configuration = datatypes.JSON(updatedConfig)
	return nil
}

// Method to set default values before updating a record
func (h *Habit) BeforeUpdate(tx *gorm.DB) (err error) {
	return h.BeforeCreate(tx) // Use the same logic as BeforeCreate
}
