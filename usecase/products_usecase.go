package usecase

import (
	"github.com/ReygaFitra/gorm-console-apps/config/entity"
	"github.com/ReygaFitra/gorm-console-apps/repository"
)

type ProductUsecase interface {
	AddProduct(product *entity.Products) error
}

type productUsecase struct {
	productRepo repository.ProductRepository
}

func (u *productUsecase) AddProduct(product *entity.Products) error {
	return u.productRepo.InsertProduct(product)
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase{
	return &productUsecase{
		productRepo: productRepo,
	}
}