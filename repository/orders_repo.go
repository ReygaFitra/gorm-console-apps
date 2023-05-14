package repository

import (
	"log"

	"github.com/ReygaFitra/gorm-console-apps/config/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	InsertOrder(order *entity.Orders) (*entity.Orders, error)
	InsertOrderDetails(orderDetail *[]entity.OrderDetails) error
	GetAllOrders() ([]entity.OrderDetails, error)
}

type orderRepository struct {
	db *gorm.DB
}

func (r *orderRepository) InsertOrder(order *entity.Orders) (*entity.Orders, error) {
	res := r.db.Create(order)
	if res.Error != nil {
		panic("Failed Insert Order")
	}
	return order, nil
}

func (r *orderRepository) InsertOrderDetails(orderDetail *[]entity.OrderDetails) error {
	res := r.db.Create(orderDetail)
	if res.Error != nil {
		panic("Failed Insert Order Detail")
	}
	return nil
}

func (r *orderRepository) GetAllOrders() ([]entity.OrderDetails, error) {
	var orders []entity.OrderDetails 
	res := r.db.Joins("Products").Joins("Orders").Find(&orders)
	if res.Error != nil {
		log.Fatal("Failed Show All Orders!")
		return nil, res.Error
	}
	return orders, nil
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}
