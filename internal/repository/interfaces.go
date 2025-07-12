package repository

import "workout-api/internal/models"

// UserRepositoryInterface defines the contract for user repository operations
type UserRepositoryInterface interface {
	Create(user models.User) error
	GetById(id int) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetAll() ([]models.User, error)
	Update(user models.User) error
	Delete(id int) error
}

// ExerciseRepositoryInterface defines the contract for exercise repository operations
type ExerciseRepositoryInterface interface {
	Create(exercise models.Exercise) error
	GetById(id int) (models.Exercise, error)
	GetAll() ([]models.Exercise, error)
	GetByMuscleGroup(muscleGroup string) ([]models.Exercise, error)
	Update(exercise models.Exercise) error
	Delete(id int) error
}
