package models

import "time"

type Exercise struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	MuscleGroup   string    `json:"muscle_group"`
	EquipmentType string    `json:"equipment_type"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
