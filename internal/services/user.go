package services

import (
	"errors"
	"workout-api/internal/models"
	"workout-api/internal/repository"
)

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user models.User) error {
	// Check if user already exists
	existingUser, err := s.repo.GetByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser.ID != 0 {
		return errors.New("user already exists")
	}

	return s.repo.Create(user)
}

func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
