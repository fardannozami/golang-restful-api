package router

import (
	"github.com/fardannozami/golang-restful-api/controller"
	"github.com/julienschmidt/httprouter"
)

func HabitRoutes(router *httprouter.Router, habitController controller.HabitController) {
	// Habit routes
	router.GET("/api/habits", habitController.GetAll)
	router.GET("/api/habits/:id", habitController.GetById)
	router.POST("/api/habits", habitController.Create)
	router.PUT("/api/habits/:id", habitController.Update)
	router.DELETE("/api/habits/:id", habitController.Delete)

}
