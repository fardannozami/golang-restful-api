package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/fardannozami/golang-restful-api/app"
	"github.com/fardannozami/golang-restful-api/controller"
	"github.com/fardannozami/golang-restful-api/helper"
	"github.com/fardannozami/golang-restful-api/repository"
	"github.com/fardannozami/golang-restful-api/router"
	"github.com/fardannozami/golang-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Setup dependencies
	validate := validator.New()
	db, err := app.NewMySqlDB()
	helper.PanicIfError(err)

	// Setup repositories
	habitRepository := repository.NewMysqlHabitRepository()
	habitCheckRepository := repository.NewHabitCheckRepository()

	// Setup services
	habitService := service.NewHabitService(habitRepository, db, validate)
	habitCheckService := service.NewHabitCheckService(db, habitCheckRepository, habitRepository, validate)

	// Setup controllers
	habitController := controller.NewHabitController(habitService)
	habitCheckController := controller.NewHabitCheckController(habitCheckService)

	// Setup router
	r := httprouter.New()
	router.HabitRoutes(r, habitController)
	router.HabitCheckRoutes(r, habitCheckController)

	// Setup HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Run server asynchronously
	go func() {
		fmt.Println("🚀 Server running on http://localhost:" + port)
		helper.PanicIfError(server.ListenAndServe())
	}()

	// Channel untuk menangkap sinyal interrupt
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit // Tunggu sampai user tekan CTRL+C

	fmt.Println("\n🛑 Server shutting down...")

	// Buat context dengan timeout 5 detik untuk shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server secara graceful
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("❌ Server forced to shutdown: %s\n", err)
	} else {
		fmt.Println("✅ Server exited gracefully")
	}
}
