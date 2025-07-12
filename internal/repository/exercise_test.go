package repository

import (
	"database/sql"
	"testing"
	"time"
	"workout-api/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestExerciseRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewExerciseRepository(db)
	exercise := models.Exercise{
		Name:          "Push-ups",
		MuscleGroup:   "Chest",
		EquipmentType: "Bodyweight",
		Notes:         "Standard push-ups",
	}

	mock.ExpectExec("INSERT INTO exercises").
		WithArgs(exercise.Name, exercise.MuscleGroup, exercise.EquipmentType, exercise.Notes).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(exercise)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestExerciseRepository_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewExerciseRepository(db)
	expectedTime := time.Now()
	expectedExercise := models.Exercise{
		ID:            1,
		Name:          "Push-ups",
		MuscleGroup:   "Chest",
		EquipmentType: "Bodyweight",
		Notes:         "Standard push-ups",
		CreatedAt:     expectedTime,
		UpdatedAt:     expectedTime,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "muscle_group", "equipment_type", "notes", "created_at", "updated_at"}).
		AddRow(expectedExercise.ID, expectedExercise.Name, expectedExercise.MuscleGroup,
			expectedExercise.EquipmentType, expectedExercise.Notes, expectedExercise.CreatedAt, expectedExercise.UpdatedAt)

	mock.ExpectQuery("SELECT (.+) FROM exercises WHERE id = \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	exercise, err := repo.GetById(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedExercise, exercise)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestExerciseRepository_GetById_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewExerciseRepository(db)

	mock.ExpectQuery("SELECT (.+) FROM exercises WHERE id = \\$1").
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	exercise, err := repo.GetById(1)
	assert.NoError(t, err)
	assert.Equal(t, models.Exercise{}, exercise)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestExerciseRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewExerciseRepository(db)
	expectedTime := time.Now()

	rows := sqlmock.NewRows([]string{"id", "name", "muscle_group", "equipment_type", "notes", "created_at", "updated_at"}).
		AddRow(1, "Push-ups", "Chest", "Bodyweight", "Standard push-ups", expectedTime, expectedTime).
		AddRow(2, "Squats", "Legs", "Bodyweight", "Basic squats", expectedTime, expectedTime)

	mock.ExpectQuery("SELECT (.+) FROM exercises").
		WillReturnRows(rows)

	exercises, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, exercises, 2)
	assert.Equal(t, "Push-ups", exercises[0].Name)
	assert.Equal(t, "Squats", exercises[1].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestExerciseRepository_GetByMuscleGroup(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewExerciseRepository(db)
	expectedTime := time.Now()

	rows := sqlmock.NewRows([]string{"id", "name", "muscle_group", "equipment_type", "notes", "created_at", "updated_at"}).
		AddRow(1, "Push-ups", "Chest", "Bodyweight", "Standard push-ups", expectedTime, expectedTime).
		AddRow(2, "Bench Press", "Chest", "Barbell", "Heavy bench press", expectedTime, expectedTime)

	mock.ExpectQuery("SELECT (.+) FROM exercises WHERE muscle_group = \\$1").
		WithArgs("Chest").
		WillReturnRows(rows)

	exercises, err := repo.GetByMuscleGroup("Chest")
	assert.NoError(t, err)
	assert.Len(t, exercises, 2)
	assert.Equal(t, "Chest", exercises[0].MuscleGroup)
	assert.Equal(t, "Chest", exercises[1].MuscleGroup)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestExerciseRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewExerciseRepository(db)
	exercise := models.Exercise{
		ID:            1,
		Name:          "Modified Push-ups",
		MuscleGroup:   "Chest",
		EquipmentType: "Bodyweight",
		Notes:         "Modified push-ups for beginners",
	}

	mock.ExpectExec("UPDATE exercises SET").
		WithArgs(exercise.Name, exercise.MuscleGroup, exercise.EquipmentType, exercise.Notes, exercise.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(exercise)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestExerciseRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewExerciseRepository(db)

	mock.ExpectExec("DELETE FROM exercises WHERE id = \\$1").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete(1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
