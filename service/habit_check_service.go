package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/fardannozami/golang-restful-api/helper"
	"github.com/fardannozami/golang-restful-api/repository"
	"github.com/fardannozami/golang-restful-api/request"
	"github.com/go-playground/validator/v10"
)

type HabitCheckService interface {
	Check(ctx context.Context, req request.CheckHabitRequest) error
	GetCheckHistory(ctx context.Context, habitId int) ([]time.Time, error)
}

type habitCheckService struct {
	dB                   *sql.DB
	habitCheckRepository repository.HabitCheckRepository
	habitRepository      repository.HabitRepository
	validate             *validator.Validate
}

func NewHabitCheckService(db *sql.DB, habitCheckRepository repository.HabitCheckRepository, habitRepository repository.HabitRepository, validate *validator.Validate) HabitCheckService {
	return &habitCheckService{
		dB:                   db,
		habitCheckRepository: habitCheckRepository,
		habitRepository:      habitRepository,
		validate:             validate,
	}
}

func (s *habitCheckService) Check(ctx context.Context, req request.CheckHabitRequest) error {
	if err := s.validate.Struct(req); err != nil {
		return err
	}

	tx, err := s.dB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	_, err = s.habitRepository.GetById(ctx, tx, req.ID)
	if err != nil {
		return err
	}

	checkDate, err := time.Parse(helper.DateLayout, req.CheckDate)
	if err != nil {
		return err
	}

	return s.habitCheckRepository.Check(ctx, tx, req.ID, checkDate)
}

func (s *habitCheckService) GetCheckHistory(ctx context.Context, habitId int) ([]time.Time, error) {
	if habitId <= 0 {
		return nil, helper.ErrHabitIdNotValid
	}

	tx, err := s.dB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	habit, err := s.habitRepository.GetById(ctx, tx, habitId)
	if err != nil {
		return nil, err
	}

	checkDates, err := s.habitCheckRepository.GetCheckHistory(ctx, tx, habit.ID)
	if err != nil {
		return nil, err
	}

	return checkDates, nil
}
