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

func (c *habitCheckController) Check(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(req.Body)
	var checkHabitRequest request.CheckHabitRequest
	err := decoder.Decode(&checkHabitRequest)
	helper.PanicIfError(err)

	c.habitCheckService.Check(req.Context(), checkHabitRequest)

	apiResponse := response.ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(apiResponse)
	helper.PanicIfError(err)
}

func (c *habitCheckController) GetCheckHistory(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	habitId := params.ByName("id")
	id, err := strconv.Atoi(habitId)
	helper.PanicIfError(err)

	habitChecks := c.habitCheckService.GetCheckHistory(req.Context(), id)

	apiResponse := response.ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    habitChecks,
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(apiResponse)
	helper.PanicIfError(err)
}
