package repository

import (
	"log"

	"github.com/ReygaFitra/gorm-console-apps/config/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProduct(product *entity.Products) error
	GetAllProducts() ([]entity.Products, error)
	GetProductByPcode(pCode string) (*entity.Products, error)
	UpdateProduct(pCode string, dataProduct *entity.Products) error
}

type productRepository struct {
	db *gorm.DB
}

func (r *productRepository) InsertProduct(product *entity.Products) error {
	res := r.db.Create(product)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *productRepository) GetAllProducts() ([]entity.Products, error) {
	var products []entity.Products
	res := r.db.Find(&products)
	if res.Error != nil {
		log.Println("Failed Show All Products")
		return nil, res.Error
	}
	return products, nil
}

func (r *productRepository) GetProductByPcode(pCode string) (*entity.Products, error) {
	var product *entity.Products
	res := r.db.Where("product_code = ?", pCode).First(&product)
	if res.Error != nil {
		log.Println("Failed Show Product")
		return nil, res.Error
	}
	return product, nil
}

func (r *productRepository) UpdateProduct(pCode string, dataProduct *entity.Products) error {
	res := r.db.Where("product_code = ?", pCode).Updates(&dataProduct)
	if res.Error != nil {
		log.Println("Failed Updating Product!")
		return res.Error
	}
	return nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}
