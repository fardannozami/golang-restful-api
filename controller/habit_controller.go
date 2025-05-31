package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fardannozami/golang-restful-api/helper"
	"github.com/fardannozami/golang-restful-api/request"
	"github.com/fardannozami/golang-restful-api/response"
	"github.com/fardannozami/golang-restful-api/service"
	"github.com/julienschmidt/httprouter"
)

type HabitController interface {
	Create(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	GetById(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
}

type habitController struct {
	habitService service.HabitService
}

func NewHabitController(habitService service.HabitService) HabitController {
	return &habitController{habitService: habitService}
}

func (c *habitController) Create(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(req.Body)
	var habitCreateRequest request.HabitCreateRequest
	err := decoder.Decode(&habitCreateRequest)
	helper.PanicIfError(err)

	habit := c.habitService.Create(req.Context(), habitCreateRequest)

	apiResponse := response.ApiResponse{
		Code:    http.StatusCreated,
		Message: "created",
		Data:    habit,
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(apiResponse)
	helper.PanicIfError(err)
}

func (c *habitController) Update(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(req.Body)
	var habitUpdateRequest request.HabitUpdateRequest
	err := decoder.Decode(&habitUpdateRequest)
	helper.PanicIfError(err)

	habitId, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	habitUpdateRequest.ID = habitId

	habit := c.habitService.Update(req.Context(), habitUpdateRequest)
	apiResponse := response.ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    habit,
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(apiResponse)
	helper.PanicIfError(err)
}

func (c *habitController) Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	habitId, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	c.habitService.Delete(req.Context(), habitId)

	apiResponse := response.ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(apiResponse)
	helper.PanicIfError(err)
}

func (c *habitController) GetAll(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	habits := c.habitService.GetAll(req.Context())
	apiResponse := response.ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    habits,
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(apiResponse)
	helper.PanicIfError(err)
}

func (c *habitController) GetById(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	habitId, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	habit := c.habitService.GetById(req.Context(), habitId)
	apiResponse := response.ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    habit,
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(apiResponse)
	helper.PanicIfError(err)

}
