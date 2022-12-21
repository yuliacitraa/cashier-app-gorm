package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (p *ProductRepository) AddProduct(product model.Product) error {
	err := p.db.Create(&product).Error
	return err
}

func (p *ProductRepository) ReadProducts() ([]model.Product, error) {
	product := []model.Product{}
	err := p.db.Find(&product).Error
	return product, err
}

func (p *ProductRepository) DeleteProduct(id uint) error {
	product := model.Product{}
	err := p.db.Where("id = ?", id).Delete(&product).Error
	return err
}

func (p *ProductRepository) UpdateProduct(id uint, product model.Product) error {
	err := p.db.Model(&model.Product{}).Where("id=?", id).Updates(&product).Error
	return err
}
