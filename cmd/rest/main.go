package main

import (
	"net/http"

	"github.com/alvesph/challenge-clean-architecture/configs"
	"github.com/alvesph/challenge-clean-architecture/internal/entity"
	"github.com/alvesph/challenge-clean-architecture/internal/handler"
	"github.com/alvesph/challenge-clean-architecture/internal/repository"
	"github.com/alvesph/challenge-clean-architecture/internal/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	db.AutoMigrate(&entity.Order{})

	orderRepo := repository.NewOrder(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandlers := handler.OrderHandlers{Service: orderService}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/orders", orderHandlers.CreateOrder)
	r.Get("/orders", orderHandlers.GetOrders)

	http.ListenAndServe(":"+cfg.RestPort, r)
}
