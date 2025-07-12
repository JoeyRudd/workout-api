package services

import (
	"errors"
	"workout-api/internal/models"
	"workout-api/internal/repository"
)

type ExerciseService struct {
	repo repository.ExerciseRepositoryInterface
}

func NewExerciseService(repo repository.ExerciseRepositoryInterface) *ExerciseService {
	return &ExerciseService{repo: repo}
}

func (s *ExerciseService) CreateExercise(exercise models.Exercise) error {
	// Basic validation
	if exercise.Name == "" {
		return errors.New("exercise name is required")
	}
	if exercise.MuscleGroup == "" {
		return errors.New("muscle group is required")
	}

	return s.repo.Create(exercise)
}

func (s *ExerciseService) GetExerciseByID(id int) (models.Exercise, error) {
	if id <= 0 {
		return models.Exercise{}, errors.New("invalid exercise ID")
	}
	return s.repo.GetById(id)
}

func (s *ExerciseService) GetAllExercises() ([]models.Exercise, error) {
	return s.repo.GetAll()
}

func (s *ExerciseService) GetExercisesByMuscleGroup(muscleGroup string) ([]models.Exercise, error) {
	if muscleGroup == "" {
		return nil, errors.New("muscle group cannot be empty")
	}
	return s.repo.GetByMuscleGroup(muscleGroup)
}

func (s *ExerciseService) UpdateExercise(exercise models.Exercise) error {
	if exercise.ID <= 0 {
		return errors.New("invalid exercise ID")
	}
	if exercise.Name == "" {
		return errors.New("exercise name is required")
	}
	if exercise.MuscleGroup == "" {
		return errors.New("muscle group is required")
	}

	return s.repo.Update(exercise)
}

func (s *ExerciseService) DeleteExercise(id int) error {
	if id <= 0 {
		return errors.New("invalid exercise ID")
	}
	return s.repo.Delete(id)
}
