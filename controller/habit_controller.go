package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	var habitCreateRequest request.HabitCreateRequest
	if err := json.NewDecoder(req.Body).Decode(&habitCreateRequest); err != nil {
		response.WriteValidationError(w, "invalid request body")
		return
	}

	_, err := c.habitService.Create(req.Context(), habitCreateRequest)
	if err != nil {
		response.WriteInternalError(w, err.Error())
		return
	}

	response.WriteCreated(w, "habit created successfully")
}

func (c *habitController) Update(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var habitUpdateRequest request.HabitUpdateRequest
	if err := json.NewDecoder(req.Body).Decode(&habitUpdateRequest); err != nil {
		response.WriteValidationError(w, "invalid request body")
		return
	}

	habitId, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		response.WriteValidationError(w, "invalid habit id")
		return
	}
	habitUpdateRequest.ID = habitId

	habitResponse, err := c.habitService.Update(req.Context(), habitUpdateRequest)
	if err != nil {
		response.WriteInternalError(w, err.Error())
		return
	}

	response.WriteData(w, habitResponse)
}

func (c *habitController) Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	habitId, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		response.WriteValidationError(w, "invalid habit id")
		return
	}

	if err := c.habitService.Delete(req.Context(), habitId); err != nil {
		response.WriteInternalError(w, err.Error())
		return
	}

	response.WriteSuccess(w, "habit deleted successfully")
}

func (c *habitController) GetAll(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	habits, err := c.habitService.GetAll(req.Context())
	if err != nil {
		response.WriteInternalError(w, err.Error())
		return
	}

	response.WriteData(w, habits)
}

func (c *habitController) GetById(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	habitId, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		response.WriteValidationError(w, "invalid habit id")
		return
	}

	habit, err := c.habitService.GetById(req.Context(), habitId)
	if err != nil {
		response.WriteInternalError(w, err.Error())
		return
	}

	response.WriteData(w, habit)
}
