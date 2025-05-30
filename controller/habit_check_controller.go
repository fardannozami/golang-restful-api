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

type HabitCheckController interface {
	Check(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	GetCheckHistory(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
}

type habitCheckController struct {
	habitCheckService service.HabitCheckService
}

func NewHabitCheckController(habitCheckService service.HabitCheckService) HabitCheckController {
	return &habitCheckController{habitCheckService: habitCheckService}
}

func (c *habitCheckController) Check(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var checkHabitRequest request.CheckHabitRequest
	if err := json.NewDecoder(req.Body).Decode(&checkHabitRequest); err != nil {
		response.WriteValidationError(writer, "invalid request body")
		return
	}

	if err := c.habitCheckService.Check(req.Context(), checkHabitRequest); err != nil {
		response.WriteValidationError(writer, err.Error())
		return
	}

	response.WriteSuccess(writer, "habit checked successfully")
}

func (c *habitCheckController) GetCheckHistory(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	habitId := params.ByName("id")
	id, err := strconv.Atoi(habitId)
	if err != nil {
		response.WriteValidationError(writer, "invalid habit id")
		return
	}

	habitChecks, err := c.habitCheckService.GetCheckHistory(req.Context(), id)
	if err != nil {
		response.WriteValidationError(writer, err.Error())
		return
	}

	response.WriteData(writer, habitChecks)
}
