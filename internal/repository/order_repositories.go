package repository

import (
	"github.com/alvesph/challenge-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type Order struct {
	DB *gorm.DB
}

func NewOrder(db *gorm.DB) *Order {
	return &Order{DB: db}
}

func (o *Order) Create(order *entity.Order) error {
	return o.DB.Create(order).Error
}

func (o *Order) FindAll() ([]entity.Order, error) {
	var orders []entity.Order
	if err := o.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
