package repository

import (
	"database/sql"
	"workout-api/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user models.User) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password)
	return err
}

func (r *UserRepository) GetById(id int) (models.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE id = $1"
	u := &models.User{}
	err := r.db.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err == sql.ErrNoRows {
		return models.User{}, nil
	}
	return *u, err
}

func (r *UserRepository) GetByEmail(email string) (models.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = $1"
	u := &models.User{}
	err := r.db.QueryRow(query, email).Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err == sql.ErrNoRows {
		return models.User{}, nil
	}
	return *u, err
}

func (r *UserRepository) Update(user models.User) error {
	query := "UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	query := "SELECT id, name, email, password FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}
