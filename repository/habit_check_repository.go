package repository

import (
	"context"
	"database/sql"
	"time"
)

type HabitCheckRepository interface {
	Check(ctx context.Context, tx *sql.Tx, habitId int, checkDate time.Time) error
	GetCheckHistory(ctx context.Context, tx *sql.Tx, habitId int) ([]time.Time, error)
}

type habitCheckRepository struct{}

func NewHabitCheckRepository() HabitCheckRepository {
	return &habitCheckRepository{}
}

func (r *habitCheckRepository) Check(ctx context.Context, tx *sql.Tx, habitId int, checkDate time.Time) error {
	SQL := "INSERT INTO habit_checks(habit_id, check_date) VALUES(?, ?)"

	_, err := tx.ExecContext(ctx, SQL, habitId, checkDate)
	if err != nil {
		return err
	}

	return nil
}

func (r *habitCheckRepository) GetCheckHistory(ctx context.Context, tx *sql.Tx, habitId int) ([]time.Time, error) {
	SQL := "SELECT check_date FROM habit_checks WHERE habit_id = ?"

	rows, err := tx.QueryContext(ctx, SQL, habitId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var checkDates []time.Time
	for rows.Next() {
		var checkDate time.Time
		err := rows.Scan(&checkDate)
		if err != nil {
			return nil, err
		}

		checkDates = append(checkDates, checkDate)
	}

	return checkDates, nil
}
