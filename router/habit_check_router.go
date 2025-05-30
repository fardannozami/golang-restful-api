package router

import (
	"github.com/fardannozami/golang-restful-api/controller"
	"github.com/julienschmidt/httprouter"
)

func HabitCheckRoutes(router *httprouter.Router, habitCheckController controller.HabitCheckController) {
	// Habit check routes
	router.GET("/api/habit-checks/:id", habitCheckController.GetCheckHistory)
	router.POST("/api/habit-checks", habitCheckController.Check)
}
