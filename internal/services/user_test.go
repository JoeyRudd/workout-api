package services

import (
	"errors"
	"testing"
	"time"
	"workout-api/internal/models"
	"workout-api/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock UserRepository that implements repository.UserRepositoryInterface
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetById(id int) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(email string) (models.User, error) {
	args := m.Called(email)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) GetAll() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) Update(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Ensure MockUserRepository implements the interface
var _ repository.UserRepositoryInterface = (*MockUserRepository)(nil)

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := models.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "hashedpassword",
	}

	// Mock that user doesn't exist
	mockRepo.On("GetByEmail", user.Email).Return(models.User{}, nil)
	mockRepo.On("Create", user).Return(nil)

	err := service.CreateUser(user)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser_UserAlreadyExists(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := models.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "hashedpassword",
	}

	existingUser := models.User{
		ID:    1,
		Name:  "Existing John",
		Email: "john@example.com",
	}

	// Mock that user already exists
	mockRepo.On("GetByEmail", user.Email).Return(existingUser, nil)

	err := service.CreateUser(user)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user already exists")
	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser_GetByEmailError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := models.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "hashedpassword",
	}

	// Mock database error during GetByEmail
	mockRepo.On("GetByEmail", user.Email).Return(models.User{}, errors.New("database error"))

	err := service.CreateUser(user)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database error")
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	expectedUser := models.User{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetById", 1).Return(expectedUser, nil)

	user, err := service.GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	mockRepo.On("GetById", 999).Return(models.User{}, errors.New("user not found"))

	user, err := service.GetUserByID(999)
	assert.Error(t, err)
	assert.Equal(t, models.User{}, user)
	assert.Contains(t, err.Error(), "user not found")
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	expectedUsers := []models.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
	}

	mockRepo.On("GetAll").Return(expectedUsers, nil)

	users, err := service.GetAllUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, expectedUsers, users)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetAllUsers_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	mockRepo.On("GetAll").Return([]models.User{}, errors.New("database connection failed"))

	users, err := service.GetAllUsers()
	assert.Error(t, err)
	assert.Empty(t, users)
	assert.Contains(t, err.Error(), "database connection failed")
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	mockRepo.On("Delete", 1).Return(nil)

	err := service.DeleteUser(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	mockRepo.On("Delete", 999).Return(errors.New("user not found"))

	err := service.DeleteUser(999)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")
	mockRepo.AssertExpectations(t)
}
