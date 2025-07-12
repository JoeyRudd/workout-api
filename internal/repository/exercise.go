package repository

import (
	"database/sql"
	"workout-api/internal/models"
)

type ExerciseRepository struct {
	db *sql.DB
}

func NewExerciseRepository(db *sql.DB) *ExerciseRepository {
	return &ExerciseRepository{db: db}
}

func (r *ExerciseRepository) Create(exercise models.Exercise) error {
	query := "INSERT INTO exercises (name, muscle_group, equipment_type, notes) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, exercise.Name, exercise.MuscleGroup, exercise.EquipmentType, exercise.Notes)
	return err
}

func (r *ExerciseRepository) GetById(id int) (models.Exercise, error) {
	query := "SELECT id, name, muscle_group, equipment_type, notes, created_at, updated_at FROM exercises WHERE id = $1"
	e := &models.Exercise{}
	err := r.db.QueryRow(query, id).Scan(&e.ID, &e.Name, &e.MuscleGroup, &e.EquipmentType, &e.Notes, &e.CreatedAt, &e.UpdatedAt)
	if err == sql.ErrNoRows {
		return models.Exercise{}, nil
	}
	return *e, err
}

func (r *ExerciseRepository) GetAll() ([]models.Exercise, error) {
	query := "SELECT id, name, muscle_group, equipment_type, notes, created_at, updated_at FROM exercises"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []models.Exercise
	for rows.Next() {
		var e models.Exercise
		err := rows.Scan(&e.ID, &e.Name, &e.MuscleGroup, &e.EquipmentType, &e.Notes, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, e)
	}
	return exercises, nil
}

func (r *ExerciseRepository) GetByMuscleGroup(muscleGroup string) ([]models.Exercise, error) {
	query := "SELECT id, name, muscle_group, equipment_type, notes, created_at, updated_at FROM exercises WHERE muscle_group = $1"
	rows, err := r.db.Query(query, muscleGroup)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []models.Exercise
	for rows.Next() {
		var e models.Exercise
		err := rows.Scan(&e.ID, &e.Name, &e.MuscleGroup, &e.EquipmentType, &e.Notes, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, e)
	}
	return exercises, nil
}

func (r *ExerciseRepository) Update(exercise models.Exercise) error {
	query := "UPDATE exercises SET name = $1, muscle_group = $2, equipment_type = $3, notes = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $5"
	_, err := r.db.Exec(query, exercise.Name, exercise.MuscleGroup, exercise.EquipmentType, exercise.Notes, exercise.ID)
	return err
}

func (r *ExerciseRepository) Delete(id int) error {
	query := "DELETE FROM exercises WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
