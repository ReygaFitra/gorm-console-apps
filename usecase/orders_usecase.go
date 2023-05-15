package usecase

import (
	"log"

	"github.com/ReygaFitra/gorm-console-apps/config/entity"
	"github.com/ReygaFitra/gorm-console-apps/repository"
)

type OrderUsecase interface {
	AddOrder(order *entity.Orders) (*entity.Orders, error)
	AddOrderDetails(orderDetail *[]entity.OrderDetails) error
	ShowAllOrders() ([]entity.OrderDetails, error)
	EditOrder(id int, order *entity.Orders) error
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
	 err := u.orderRepo.InsertOrderDetails(orderDetails)
	 if err != nil {
		log.Fatal("Something whent wrong")
		return err
	 }
	 return nil
}

func (u *orderUsecase) ShowAllOrders() ([]entity.OrderDetails, error) {
	res, err := u.orderRepo.GetAllOrders()
	if err != nil {
		log.Fatal("Something whent wrong")
		return nil, err
	}
	return res, nil
}

func (u *orderUsecase) EditOrder(id int, order *entity.Orders) error {
	err := u.orderRepo.UpdateOrder(id, order)
	if err != nil {
		log.Fatal("Something whent wrong")
		return err
	}
	return nil
}

func NewOrderUsecase(orderRepo repository.OrderRepository) OrderUsecase {
	return &orderUsecase{
		orderRepo: orderRepo,
	}
}