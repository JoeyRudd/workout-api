package services

import (
	"testing"
	"time"
	"workout-api/internal/models"
	"workout-api/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock ExerciseRepository that implements repository.ExerciseRepositoryInterface
type MockExerciseRepository struct {
	mock.Mock
}

func (m *MockExerciseRepository) Create(exercise models.Exercise) error {
	args := m.Called(exercise)
	return args.Error(0)
}

func (m *MockExerciseRepository) GetById(id int) (models.Exercise, error) {
	args := m.Called(id)
	return args.Get(0).(models.Exercise), args.Error(1)
}

func (m *MockExerciseRepository) GetAll() ([]models.Exercise, error) {
	args := m.Called()
	return args.Get(0).([]models.Exercise), args.Error(1)
}

func (m *MockExerciseRepository) GetByMuscleGroup(muscleGroup string) ([]models.Exercise, error) {
	args := m.Called(muscleGroup)
	return args.Get(0).([]models.Exercise), args.Error(1)
}

func (m *MockExerciseRepository) Update(exercise models.Exercise) error {
	args := m.Called(exercise)
	return args.Error(0)
}

func (m *MockExerciseRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Ensure MockExerciseRepository implements the interface
var _ repository.ExerciseRepositoryInterface = (*MockExerciseRepository)(nil)

func TestExerciseService_CreateExercise(t *testing.T) {
	mockRepo := new(MockExerciseRepository)
	service := NewExerciseService(mockRepo)

	exercise := models.Exercise{
		Name:          "Push-ups",
		MuscleGroup:   "Chest",
		EquipmentType: "Bodyweight",
		Notes:         "Standard push-ups",
	}

	mockRepo.On("Create", exercise).Return(nil)

	err := service.CreateExercise(exercise)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestExerciseService_CreateExercise_ValidationErrors(t *testing.T) {
	mockRepo := new(MockExerciseRepository)
	service := NewExerciseService(mockRepo)

	// Test empty name
	exercise := models.Exercise{
		Name:        "",
		MuscleGroup: "Chest",
	}
	err := service.CreateExercise(exercise)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "exercise name is required")

	// Test empty muscle group
	exercise = models.Exercise{
		Name:        "Push-ups",
		MuscleGroup: "",
	}
	err = service.CreateExercise(exercise)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "muscle group is required")
}

func TestExerciseService_GetExerciseByID(t *testing.T) {
	mockRepo := new(MockExerciseRepository)
	service := NewExerciseService(mockRepo)

	expectedExercise := models.Exercise{
		ID:            1,
		Name:          "Push-ups",
		MuscleGroup:   "Chest",
		EquipmentType: "Bodyweight",
		Notes:         "Standard push-ups",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	mockRepo.On("GetById", 1).Return(expectedExercise, nil)

	exercise, err := service.GetExerciseByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedExercise, exercise)
	mockRepo.AssertExpectations(t)
}

func TestExerciseService_GetExerciseByID_InvalidID(t *testing.T) {
	mockRepo := new(MockExerciseRepository)
	service := NewExerciseService(mockRepo)

	exercise, err := service.GetExerciseByID(0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid exercise ID")
	assert.Equal(t, models.Exercise{}, exercise)
}

func TestExerciseService_GetAllExercises(t *testing.T) {
	mockRepo := new(MockExerciseRepository)
	service := NewExerciseService(mockRepo)

	expectedExercises := []models.Exercise{
		{ID: 1, Name: "Push-ups", MuscleGroup: "Chest"},
		{ID: 2, Name: "Squats", MuscleGroup: "Legs"},
	}

	mockRepo.On("GetAll").Return(expectedExercises, nil)

	exercises, err := service.GetAllExercises()
	assert.NoError(t, err)
	assert.Len(t, exercises, 2)
	assert.Equal(t, expectedExercises, exercises)
	mockRepo.AssertExpectations(t)
}

func TestExerciseService_GetExercisesByMuscleGroup(t *testing.T) {
	mockRepo := new(MockExerciseRepository)
	service := NewExerciseService(mockRepo)

	expectedExercises := []models.Exercise{
		{ID: 1, Name: "Push-ups", MuscleGroup: "Chest"},
		{ID: 3, Name: "Bench Press", MuscleGroup: "Chest"},
	}

	mockRepo.On("GetByMuscleGroup", "Chest").Return(expectedExercises, nil)

	exercises, err := service.GetExercisesByMuscleGroup("Chest")
	assert.NoError(t, err)
	assert.Len(t, exercises, 2)
	assert.Equal(t, expectedExercises, exercises)
	mockRepo.AssertExpectations(t)
}

func TestExerciseService_GetExercisesByMuscleGroup_EmptyMuscleGroup(t *testing.T) {
	mockRepo := new(MockExerciseRepository)
	service := NewExerciseService(mockRepo)

	exercises, err := service.GetExercisesByMuscleGroup("")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "muscle group cannot be empty")
	assert.Nil(t, exercises)
}

func TestExerciseService_UpdateExercise(t *testing.T) {
	mockRepo := new(MockExerciseRepository)
	service := NewExerciseService(mockRepo)

	exercise := models.Exercise{
		ID:            1,
		Name:          "Modified Push-ups",
		MuscleGroup:   "Chest",
		EquipmentType: "Bodyweight",
		Notes:         "Modified for beginners",
	}

	mockRepo.On("Update", exercise).Return(nil)

	err := service.UpdateExercise(exercise)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestExerciseService_UpdateExercise_ValidationErrors(t *testing.T) {
	mockRepo := new(MockExerciseRepository)
	service := NewExerciseService(mockRepo)

	// Test invalid ID
	exercise := models.Exercise{
		ID:          0,
		Name:        "Push-ups",
		MuscleGroup: "Chest",
	}
	err := service.UpdateExercise(exercise)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid exercise ID")

	// Test empty name
	exercise = models.Exercise{
		ID:          1,
		Name:        "",
		MuscleGroup: "Chest",
	}
	err = service.UpdateExercise(exercise)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "exercise name is required")

	// Test empty muscle group
	exercise = models.Exercise{
		ID:          1,
		Name:        "Push-ups",
		MuscleGroup: "",
	}
	err = service.UpdateExercise(exercise)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "muscle group is required")
}

func TestExerciseService_DeleteExercise(t *testing.T) {
	mockRepo := new(MockExerciseRepository)
	service := NewExerciseService(mockRepo)

	mockRepo.On("Delete", 1).Return(nil)

	err := service.DeleteExercise(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestExerciseService_DeleteExercise_InvalidID(t *testing.T) {
	mockRepo := new(MockExerciseRepository)
	service := NewExerciseService(mockRepo)

	err := service.DeleteExercise(0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid exercise ID")
}
