package helper

import (
	"github.com/fardannozami/golang-restful-api/model"
	"github.com/fardannozami/golang-restful-api/response"
)

func ToHabitResponse(habit model.Habit) response.HabitResponse {
	return response.HabitResponse{ID: habit.ID, Name: habit.Name, Description: habit.Description}
}

func ToHabitResponses(habits []model.Habit) []response.HabitResponse {
	var categoryResponses []response.HabitResponse
	for _, habit := range habits {
		categoryResponses = append(categoryResponses, ToHabitResponse(habit))
	}

	return categoryResponses
}
