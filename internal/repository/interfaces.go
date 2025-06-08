package repository

import "github.com/alvesph/challenge-clean-architecture/internal/entity"

type OrderInterface interface {
	Create(order *entity.Order) error
	FindAll() ([]entity.Order, error)
}
