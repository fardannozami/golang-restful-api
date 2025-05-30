package helper

import "errors"

var (
	ErrHabitNotFound   = errors.New("habit not found")
	ErrHabitIdNotValid = errors.New("habitId not valid")
)

func PanicIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
