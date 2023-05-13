package repository

import (
	"github.com/ReygaFitra/gorm-console-apps/config/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	InsertOrder(order *entity.Orders) (*entity.Orders, error)
	InsertOrderDetails(orderDetail *[]entity.OrderDetails) error
}

type orderRepository struct {
	db *gorm.DB
}

func(r *orderRepository) InsertOrder(order *entity.Orders) (*entity.Orders, error) {
	res := r.db.Create(order)
	if res.Error != nil {
		panic("Failed Insert Order")
	}
	return order, nil
}

func(r *orderRepository) InsertOrderDetails(orderDetail *[]entity.OrderDetails) error {
	res := r.db.Create(orderDetail)
	if res.Error != nil {
		panic("Failed Insert Order Detail")
	}
	return nil
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}
