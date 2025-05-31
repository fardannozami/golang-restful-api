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
	Check(ctx context.Context, req request.CheckHabitRequest)
	GetCheckHistory(ctx context.Context, habitId int) []time.Time
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

func (s *habitCheckService) Check(ctx context.Context, req request.CheckHabitRequest) {
	err := s.validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := s.dB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	_ = s.habitRepository.GetById(ctx, tx, req.ID)

	checkDate, err := time.Parse(helper.DateLayout, req.CheckDate)
	helper.PanicIfError(err)

	s.habitCheckRepository.Check(ctx, tx, req.ID, checkDate)
}

func (s *habitCheckService) GetCheckHistory(ctx context.Context, habitId int) []time.Time {
	tx, err := s.dB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	habit := s.habitRepository.GetById(ctx, tx, habitId)

	checkDates := s.habitCheckRepository.GetCheckHistory(ctx, tx, habit.ID)

	return checkDates
}
