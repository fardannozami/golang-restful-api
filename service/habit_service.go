package service

import (
	"context"
	"database/sql"

	"github.com/fardannozami/golang-restful-api/helper"
	"github.com/fardannozami/golang-restful-api/model"
	"github.com/fardannozami/golang-restful-api/repository"
	"github.com/fardannozami/golang-restful-api/request"
	"github.com/fardannozami/golang-restful-api/response"
	"github.com/go-playground/validator/v10"
)

type HabitService interface {
	GetAll(ctx context.Context) []response.HabitResponse
	GetById(ctx context.Context, habitId int) response.HabitResponse
	Create(ctx context.Context, request request.HabitCreateRequest) response.HabitResponse
	Update(ctx context.Context, request request.HabitUpdateRequest) response.HabitResponse
	Delete(ctx context.Context, habitId int)
}

type habitService struct {
	habitRepository repository.HabitRepository
	dB              *sql.DB
	validate        *validator.Validate
}

func NewHabitService(habitRepository repository.HabitRepository, db *sql.DB, validate *validator.Validate) HabitService {
	return &habitService{
		habitRepository: habitRepository,
		dB:              db,
		validate:        validate,
	}
}

func (s *habitService) GetAll(ctx context.Context) []response.HabitResponse {
	tx, err := s.dB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	habits := s.habitRepository.GetAll(ctx, tx)

	return helper.ToHabitResponses(habits)
}

func (s *habitService) GetById(ctx context.Context, habitId int) response.HabitResponse {
	tx, err := s.dB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	habit := s.habitRepository.GetById(ctx, tx, habitId)

	return helper.ToHabitResponse(habit)

}

func (s *habitService) Create(ctx context.Context, request request.HabitCreateRequest) response.HabitResponse {
	var habitResponse response.HabitResponse
	err := s.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.dB.Begin()
	helper.PanicIfError(err)

	habit := model.Habit{Name: request.Name, Description: request.Description}
	defer helper.CommitOrRollback(tx)

	savedHabit := s.habitRepository.Create(ctx, tx, habit)

	habitResponse = helper.ToHabitResponse(savedHabit)

	return habitResponse
}

func (s *habitService) Update(ctx context.Context, request request.HabitUpdateRequest) response.HabitResponse {
	err := s.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.dB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	habit := s.habitRepository.GetById(ctx, tx, request.ID)

	habit.Name = request.Name
	habit.Description = request.Description

	habit = s.habitRepository.Update(ctx, tx, habit)

	return helper.ToHabitResponse(habit)
}

func (s *habitService) Delete(ctx context.Context, habitId int) {
	tx, err := s.dB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	habit := s.habitRepository.GetById(ctx, tx, habitId)

	s.habitRepository.Delete(ctx, tx, habit)
}
