package usecase

import (
	"log"

	"github.com/ReygaFitra/gorm-console-apps/config/entity"
	"github.com/ReygaFitra/gorm-console-apps/repository"
)

type ProductUsecase interface {
	AddProduct(product *entity.Products) error
	ShowAllProducts() ([]entity.Products, error)
	ShowProductByPcode(pCode string) (*entity.Products, error)
	EditProduct(pCode string, dataProduct *entity.Products) error
}

type productUsecase struct {
	productRepo repository.ProductRepository
}

func (u *productUsecase) AddProduct(product *entity.Products) error {
	 err := u.productRepo.InsertProduct(product)
	 if err != nil {
		log.Fatal("Failed insert product")
		return err
	 }
	 return nil
}

func (u *productUsecase) ShowAllProducts() ([]entity.Products, error) {
	 res, err := u.productRepo.GetAllProducts()
	 if err != nil {
		log.Fatal("Show Product Failed!")
		return nil, err
	 } 
	 return res, nil
}

func (u *productUsecase) ShowProductByPcode(pCode string) (*entity.Products, error) {
	 res, err := u.productRepo.GetProductByPcode(pCode)
	 if err != nil {
		log.Fatal("Show Product Failed!")
		return nil, err
	 }
	 return res, nil
}

func (u *productUsecase) EditProduct(pCode string, dataProduct *entity.Products) error {
	err := u.productRepo.UpdateProduct(pCode, dataProduct)
	if err != nil {
		log.Fatal("Update Product Failed!")
		return err
	}
	return nil
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase{
	return &productUsecase{
		productRepo: productRepo,
	}
}