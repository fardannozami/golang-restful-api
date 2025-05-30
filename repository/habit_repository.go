package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/fardannozami/golang-restful-api/helper"
	"github.com/fardannozami/golang-restful-api/model"
)

type HabitRepository interface {
	GetAll(ctx context.Context, tx *sql.Tx) ([]model.Habit, error)
	GetById(ctx context.Context, tx *sql.Tx, habitId int) (model.Habit, error)
	Create(ctx context.Context, tx *sql.Tx, habit model.Habit) (model.Habit, error)
	Update(ctx context.Context, tx *sql.Tx, habit model.Habit) (model.Habit, error)
	Delete(ctx context.Context, tx *sql.Tx, habit model.Habit) error
}

type mysqlHabitRepository struct{}

func NewMysqlHabitRepository() HabitRepository {
	return &mysqlHabitRepository{}
}

func (r *mysqlHabitRepository) GetAll(ctx context.Context, tx *sql.Tx) ([]model.Habit, error) {
	SQL := "SELECT id, name, description, created_at FROM habits"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var habits []model.Habit
	for rows.Next() {
		habit := model.Habit{}
		err := rows.Scan(&habit.ID, &habit.Name, &habit.Description, &habit.CreatedAt)
		if err != nil {
			return nil, err
		}
		habits = append(habits, habit)
	}

	return habits, nil
}

func (r *mysqlHabitRepository) GetById(ctx context.Context, tx *sql.Tx, habitId int) (model.Habit, error) {
	var habit model.Habit

	SQL := "SELECT id, name, description, created_at FROM habits WHERE id = ?"
	err := tx.QueryRowContext(ctx, SQL, habitId).Scan(&habit.ID, &habit.Name, &habit.Description, &habit.CreatedAt)
	if err == sql.ErrNoRows {
		return habit, helper.ErrHabitNotFound
	}
	return habit, err
}

func (r *mysqlHabitRepository) Create(ctx context.Context, tx *sql.Tx, habit model.Habit) (model.Habit, error) {
	// Set CreatedAt langsung di sini
	habit.CreatedAt = time.Now()

	SQL := "INSERT INTO habits(name, description, created_at) VALUES(?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, habit.Name, habit.Description, habit.CreatedAt)
	if err != nil {
		return model.Habit{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.Habit{}, err
	}

	habit.ID = int(id)

	return habit, nil
}

func (r *mysqlHabitRepository) Update(ctx context.Context, tx *sql.Tx, habit model.Habit) (model.Habit, error) {
	SQL := "UPDATE habits SET name = ?, description = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, habit.Name, habit.Description, habit.ID)
	if err != nil {
		return model.Habit{}, err
	}

	return habit, nil
}

func (r *mysqlHabitRepository) Delete(ctx context.Context, tx *sql.Tx, habit model.Habit) error {
	SQL := "DELETE FROM habits WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, habit.ID)
	if err != nil {
		return err
	}

	return nil
}
