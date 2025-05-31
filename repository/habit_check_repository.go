package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/fardannozami/golang-restful-api/helper"
)

type HabitCheckRepository interface {
	Check(ctx context.Context, tx *sql.Tx, habitId int, checkDate time.Time)
	GetCheckHistory(ctx context.Context, tx *sql.Tx, habitId int) []time.Time
}

type habitCheckRepository struct{}

func NewHabitCheckRepository() HabitCheckRepository {
	return &habitCheckRepository{}
}

func (r *habitCheckRepository) Check(ctx context.Context, tx *sql.Tx, habitId int, checkDate time.Time) {
	SQL := "INSERT INTO habit_checks(habit_id, check_date) VALUES(?, ?)"

	_, err := tx.ExecContext(ctx, SQL, habitId, checkDate)
	helper.PanicIfError(err)
}

func (r *habitCheckRepository) GetCheckHistory(ctx context.Context, tx *sql.Tx, habitId int) []time.Time {
	SQL := "SELECT check_date FROM habit_checks WHERE habit_id = ?"

	rows, err := tx.QueryContext(ctx, SQL, habitId)
	helper.PanicIfError(err)

	defer rows.Close()

	var checkDates []time.Time
	for rows.Next() {
		var checkDate time.Time
		err := rows.Scan(&checkDate)
		helper.PanicIfError(err)

		checkDates = append(checkDates, checkDate)
	}

	return checkDates
}
