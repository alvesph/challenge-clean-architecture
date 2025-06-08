package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alvesph/challenge-clean-architecture/internal/entity"
	"github.com/alvesph/challenge-clean-architecture/internal/service"
)

type OrderHandlers struct {
	Service *service.OrderService
}

// GET /orders
func (h *OrderHandlers) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.Service.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieve orders: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// POST /order
func (h *OrderHandlers) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order entity.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := h.Service.Create(&order); err != nil {
		http.Error(w, "Failed to create order: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
