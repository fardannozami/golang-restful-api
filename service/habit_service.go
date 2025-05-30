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
	GetAll(ctx context.Context) ([]response.HabitResponse, error)
	GetById(ctx context.Context, habitId int) (response.HabitResponse, error)
	Create(ctx context.Context, request request.HabitCreateRequest) (response.HabitResponse, error)
	Update(ctx context.Context, request request.HabitUpdateRequest) (response.HabitResponse, error)
	Delete(ctx context.Context, habitId int) error
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

func (s *habitService) GetAll(ctx context.Context) ([]response.HabitResponse, error) {
	tx, err := s.dB.Begin()
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx)

	habits, err := s.habitRepository.GetAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return helper.ToHabitResponses(habits), nil
}

func (s *habitService) GetById(ctx context.Context, habitId int) (response.HabitResponse, error) {
	tx, err := s.dB.Begin()
	if err != nil {
		return response.HabitResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	habit, err := s.habitRepository.GetById(ctx, tx, habitId)
	if err != nil {
		return response.HabitResponse{}, err
	}

	return helper.ToHabitResponse(habit), nil

}

func (s *habitService) Create(ctx context.Context, request request.HabitCreateRequest) (response.HabitResponse, error) {
	var habitResponse response.HabitResponse
	err := s.validate.Struct(request)
	if err != nil {
		return habitResponse, err
	}

	tx, err := s.dB.Begin()
	if err != nil {
		return habitResponse, err
	}

	habit := model.Habit{Name: request.Name, Description: request.Description}
	defer helper.CommitOrRollback(tx)

	savedHabit, err := s.habitRepository.Create(ctx, tx, habit)
	if err != nil {
		return habitResponse, err
	}

	habitResponse = helper.ToHabitResponse(savedHabit)

	return habitResponse, nil
}

func (s *habitService) Update(ctx context.Context, request request.HabitUpdateRequest) (response.HabitResponse, error) {
	var habitResponse response.HabitResponse
	err := s.validate.Struct(request)
	if err != nil {
		return habitResponse, err
	}

	tx, err := s.dB.Begin()
	if err != nil {
		return habitResponse, err
	}

	defer helper.CommitOrRollback(tx)

	habit, err := s.habitRepository.GetById(ctx, tx, request.ID)
	if err != nil {
		return habitResponse, err
	}

	habit.Name = request.Name
	habit.Description = request.Description

	habit, err = s.habitRepository.Update(ctx, tx, habit)
	if err != nil {
		return habitResponse, err
	}

	return helper.ToHabitResponse(habit), nil
}

func (s *habitService) Delete(ctx context.Context, habitId int) error {
	tx, err := s.dB.Begin()
	if err != nil {
		return err
	}

	defer helper.CommitOrRollback(tx)

	habit, err := s.habitRepository.GetById(ctx, tx, habitId)
	if err != nil {
		return err
	}

	if err := s.habitRepository.Delete(ctx, tx, habit); err != nil {
		return err
	}

	return nil
}
