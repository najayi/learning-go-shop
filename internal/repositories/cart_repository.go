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
	//TODO implement me
	panic("implement me")
}

func (c *CartRepository) Create(cart *models.Cart) error {
	//TODO implement me
	panic("implement me")
}

func (c *CartRepository) Update(cart *models.Cart) error {
	//TODO implement me
	panic("implement me")
}

func (c *CartRepository) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}
