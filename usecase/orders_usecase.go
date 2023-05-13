package usecase

import (
	"gorm-basic/config/entity"
	"gorm-basic/repository"
)

type OrderUsecase interface {
	AddOrder(order *entity.Orders) (*entity.Orders, error)
	AddOrderDetails(orderDetail *[]entity.OrderDetails) error
}

type orderUsecase struct {
	orderRepo repository.OrderRepository
}

func (u *orderUsecase) AddOrder(order *entity.Orders) (*entity.Orders, error) {
	res, err := u.orderRepo.InsertOrder(order)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *orderUsecase) AddOrderDetails(orderDetails *[]entity.OrderDetails) error {
	return u.orderRepo.InsertOrderDetails(orderDetails)
}

func NewOrderUsecase(orderRepo repository.OrderRepository) OrderUsecase {
	return &orderUsecase{
		orderRepo: orderRepo,
	}
}