package repository

import (
	"github.com/ReygaFitra/gorm-console-apps/config/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProduct(product *entity.Products) error
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

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}
