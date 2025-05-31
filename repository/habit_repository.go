package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/fardannozami/golang-restful-api/helper"
	"github.com/fardannozami/golang-restful-api/model"
)

type HabitRepository interface {
	GetAll(ctx context.Context, tx *sql.Tx) []model.Habit
	GetById(ctx context.Context, tx *sql.Tx, habitId int) model.Habit
	Create(ctx context.Context, tx *sql.Tx, habit model.Habit) model.Habit
	Update(ctx context.Context, tx *sql.Tx, habit model.Habit) model.Habit
	Delete(ctx context.Context, tx *sql.Tx, habit model.Habit)
}

type mysqlHabitRepository struct{}

func NewMysqlHabitRepository() HabitRepository {
	return &mysqlHabitRepository{}
}

func (r *mysqlHabitRepository) GetAll(ctx context.Context, tx *sql.Tx) []model.Habit {
	SQL := "SELECT id, name, description, created_at FROM habits"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var habits []model.Habit
	for rows.Next() {
		habit := model.Habit{}
		err := rows.Scan(&habit.ID, &habit.Name, &habit.Description, &habit.CreatedAt)
		helper.PanicIfError(err)

		habits = append(habits, habit)
	}

	return habits
}

func (r *mysqlHabitRepository) GetById(ctx context.Context, tx *sql.Tx, habitId int) model.Habit {
	var habit model.Habit

	SQL := "SELECT id, name, description, created_at FROM habits WHERE id = ?"
	err := tx.QueryRowContext(ctx, SQL, habitId).Scan(&habit.ID, &habit.Name, &habit.Description, &habit.CreatedAt)
	helper.PanicIfError(err)

	return habit
}

func (r *mysqlHabitRepository) Create(ctx context.Context, tx *sql.Tx, habit model.Habit) model.Habit {
	habit.CreatedAt = time.Now()

	SQL := "INSERT INTO habits(name, description, created_at) VALUES(?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, habit.Name, habit.Description, habit.CreatedAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	habit.ID = int(id)

	return habit
}

func (r *mysqlHabitRepository) Update(ctx context.Context, tx *sql.Tx, habit model.Habit) model.Habit {
	SQL := "UPDATE habits SET name = ?, description = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, habit.Name, habit.Description, habit.ID)
	helper.PanicIfError(err)

	return habit
}

func (r *mysqlHabitRepository) Delete(ctx context.Context, tx *sql.Tx, habit model.Habit) {
	SQL := "DELETE FROM habits WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, habit.ID)
	helper.PanicIfError(err)
}
