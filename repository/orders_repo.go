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
	UpdateOrder(id int, order *entity.Orders) error
}

type orderRepository struct {
	db *gorm.DB
}

func (r *orderRepository) InsertOrder(order *entity.Orders) (*entity.Orders, error) {
	res := r.db.Create(order)
	if res.Error != nil {
		log.Fatal("Failed Insert Order")
		return nil, res.Error
	}
	return order, nil
}

func (r *orderRepository) InsertOrderDetails(orderDetail *[]entity.OrderDetails) error {
	res := r.db.Create(orderDetail)
	if res.Error != nil {
		log.Fatal("Failed Insert Order Detail")
		return res.Error
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

func (r *orderRepository) UpdateOrder(id int, order *entity.Orders) error {
	res := r.db.Where("id_order = ?", id).Updates(order)
	if res.Error != nil {
		log.Fatal("Failed Updating Order!")
		return res.Error
	}
	return nil
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}
