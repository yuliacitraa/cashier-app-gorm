package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return CartRepository{db}
}

func (c *CartRepository) ReadCart() ([]model.JoinCart, error) {
	joined := []model.JoinCart{}
	err := c.db.Table("carts").Select("carts.id as id, carts.product_id as product_id, products.name as name, carts.quantity as quantity, carts.total_price as total_price").Joins("join products on carts.product_id = products.id").Scan(&joined).Error
	if err != nil {
		return []model.JoinCart{}, err
	}
	return joined, nil
}

func (c *CartRepository) AddCart(product model.Product) error {
    total := product.Price
	getDiscount := ((product.Discount / 100) * product.Price)
    if product.Discount != 0 {
        total = product.Price - getDiscount
    }

    cart := model.Cart{
        ProductID:  product.ID,
        Quantity:   1,
        TotalPrice: total,
    }

	err := c.db.Where("product_id = ?", product.ID).First(&cart).Error
    if err != nil {
		product.Stock = product.Stock - 1
		if err := c.db.Table("products").Where("id = ?", product.ID).Updates(product).Error; err != nil {
			return err
		}

		if err := c.db.Create(&cart).Error; err != nil {
			return err
		}

		return nil
    }

	product.Stock = product.Stock - 1
	if err := c.db.Table("products").Where("id = ?", product.ID).Updates(product).Error; err != nil {
		return err
	}

	cart.Quantity = cart.Quantity + 1
	cart.TotalPrice = cart.TotalPrice + total
	if err := c.db.Table("carts").Where("product_id = ?", product.ID).Updates(cart).Error; err != nil {
		return err
	}

    return nil
}

func (c *CartRepository) DeleteCart(id uint, productID uint) error {
	
	product := model.Product{}
	cart := model.Cart{}
	
	err := c.db.Model(cart).Where("product_id = ?", productID).First(&cart).Error
	if err != nil {
		return nil
	}

	err = c.db.Model(product).Where("id = ?", productID).First(&product).Error
	if err != nil {
		return nil
	}

	product.Stock = product.Stock + int(cart.Quantity)
	if err = c.db.Model(product).Where("id = ?", productID).Updates(product).Error; err != nil {
		return err
	}
	
	if err = c.db.Where("id = ?", id).Delete(&model.Cart{}).Error; err != nil {
		return err
	}

	return nil
}

func (c *CartRepository) UpdateCart(id uint, cart model.Cart) error {
	err := c.db.Model(&model.Cart{}).Where("id=?", id).Updates(&cart).Error
	return err
}
