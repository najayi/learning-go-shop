package repositories

import (
	"github.com/najayi/learning-go-shop/internal/models"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (c *CartRepository) GetByUserID(userID uint) (*models.Cart, error) {
	var cart models.Cart
	if err := c.db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (c *CartRepository) Create(cart *models.Cart) error {
	return c.db.Create(cart).Error
}

func (c *CartRepository) Update(cart *models.Cart) error {
	return c.db.Save(cart).Error
}

func (c *CartRepository) Delete(id uint) error {
	return c.db.Delete(&models.Cart{}, id).Error
}
