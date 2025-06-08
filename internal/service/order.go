package service

import (
	"github.com/alvesph/challenge-clean-architecture/internal/entity"
	"github.com/alvesph/challenge-clean-architecture/internal/repository"
)

type OrderService struct {
	OrderRepo repository.OrderInterface
}

func NewOrderService(order repository.OrderInterface) *OrderService {
	return &OrderService{
		OrderRepo: order,
	}
}

func (s *OrderService) Create(order *entity.Order) error {
	return s.OrderRepo.Create(order)
}

func (s *OrderService) FindAll() ([]entity.Order, error) {
	return s.OrderRepo.FindAll()
}
