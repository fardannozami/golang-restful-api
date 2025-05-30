package model

import "time"

type HabitCheck struct {
	ID        int       `json:"id"`
	HabitID   int       `json:"habit_id"`
	CheckDate time.Time `json:"check_date"`
}
