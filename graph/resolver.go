package graph

import "github.com/alvesph/challenge-clean-architecture/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	OrderService *service.OrderService
}
